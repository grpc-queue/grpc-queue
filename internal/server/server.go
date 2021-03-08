package server

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/grpc-queue/grpc-queue/internal/location"
	"github.com/grpc-queue/grpc-queue/pkg/grpc/v1/queue"
)

const (
	headerMessageLength = 4
	dicardBufferSize    = 1024
	readBufferSize      = 1024
	logfileThreshold    = 1024
	headPositionPaylod  = "{consume-group}|{logFile}|{byteoffset}"
	headPositionPattern = `(\w+)\|(\d+\.log)\|(\d+)`
	streamInfoPayload   = "{partitionCount}"
	consumerGroup       = "main"
)

type server struct {
	streamsMutex map[string][]partition
	location     *location.Location
	*queue.UnimplementedQueueServiceServer
}

type logfile struct {
	size int64
}
type partition struct {
	logs      []*logfile
	pushMutex sync.Mutex
	popMutex  sync.Mutex
}

type headPosition struct {
	consumerGroup string
	logFile       string
	offset        int64
}

type partitionInfo struct {
	lastLog    string
	EntryCount int
}

func NewServer(dataPath string) *server {
	location := location.NewLocation(dataPath)
	os.MkdirAll(location.StreamsFolder(), os.ModePerm)
	files, err := ioutil.ReadDir(location.StreamsFolder())
	if err != nil {
		log.Fatal(err)
	}
	s := &server{location: location}
	streams := make(map[string][]partition)
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".info") {
			partitions := make([]partition, 0)
			partitionsCount, _ := s.getStreamPartitionSize(f.Name())
			for i := 0; i < partitionsCount; i++ {
				logsCount := s.LogFilesCount(f.Name(), i)
				p := partition{logs: make([]*logfile, 0)}
				for j := 0; j < logsCount; j++ {
					l, _ := s.LogFileSize(f.Name(), i, j)
					p.logs = append(p.logs, &logfile{size: l})
				}
				partitions = append(partitions, p)
			}
			streams[f.Name()] = partitions
		}
	}
	s.streamsMutex = streams
	return s

}
func (s *server) LogFileSize(streamName string, partition, log int) (int64, error) {
	fi, err := os.Stat(s.location.StreamParttionLogEntryFile(streamName, strconv.Itoa(log)+".log", partition))
	if err != nil {
		return 0, err
	}
	size := fi.Size()
	return size, nil
}
func (s *server) LogFilesCount(streamName string, partition int) int {
	files, _ := ioutil.ReadDir(s.location.StreamPartitionFolder(streamName, partition))
	return len(files) - 1 //1 to ignore head.position
}
func (s *server) candidateLogFile(streamName string, partition int, payloadSize int64) string {
	lastLog := s.streamsMutex[streamName][partition].logs[len(s.streamsMutex[streamName][partition].logs)-1]
	if payloadSize+lastLog.size > logfileThreshold {
		s.createLogFile(streamName, partition)
		newLogFile := s.streamsMutex[streamName][partition].logs[len(s.streamsMutex[streamName][partition].logs)-1]
		newLogFile.size += payloadSize
		return strconv.Itoa(len(s.streamsMutex[streamName][partition].logs)-1) + ".log"
	}
	lastLog.size += payloadSize
	return strconv.Itoa(len(s.streamsMutex[streamName][partition].logs)-1) + ".log"
}
func (s *server) createLogFile(streamName string, partition int) string {
	s.streamsMutex[streamName][partition].logs = append(s.streamsMutex[streamName][partition].logs, &logfile{})
	logfile := strconv.Itoa(len(s.streamsMutex[streamName][partition].logs)-1) + ".log"
	f, _ := os.Create(s.location.StreamParttionLogEntryFile(streamName, logfile, partition))
	defer f.Close()
	return logfile
}

func (s *server) getStreamPartitionSize(streamName string) (int, error) {

	file, err := os.Open(s.location.StreamInfoFile(streamName))
	if err != nil {
		return 0, errors.New("combination of stream and partition not found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())
	return i, nil

}

func (s *server) writeEntry(streamName, logFile string, message []byte, partition int) {

	location := s.location.StreamParttionLogEntryFile(streamName, logFile, partition)
	file, err := os.OpenFile(location, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	var buffer bytes.Buffer
	lengthMessage := len(message)
	b := make([]byte, headerMessageLength)
	binary.LittleEndian.PutUint32(b, uint32(lengthMessage))
	buffer.Write(b)
	buffer.Write(message)
	file.Write(buffer.Bytes())

}
func (s *server) saveHeadPostion(streamName string, partition int, h *headPosition) {
	location := s.location.StreamHeadPositionFile(streamName, partition)
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	payloadReplacer := strings.NewReplacer("{consume-group}", h.consumerGroup,
		"{logFile}", h.logFile,
		"{byteoffset}", strconv.Itoa(int(h.offset)))
	data := payloadReplacer.Replace(headPositionPaylod)
	err = ioutil.WriteFile(location, []byte(data), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
func (s *server) getHeadPostion(consumerGroup, streamName string, partition int) (*headPosition, error) {
	location := s.location.StreamHeadPositionFile(streamName, partition)
	file, err := os.Open(location)
	if err != nil {
		return nil, errors.New("Head position not found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(headPositionPattern)
	for scanner.Scan() {
		groups := re.FindStringSubmatch(scanner.Text())
		if groups[1] == consumerGroup {
			offset, _ := strconv.Atoi(groups[3])
			return &headPosition{consumerGroup: groups[1], logFile: groups[2], offset: int64(offset)}, nil
		}

	}
	return nil, errors.New("Head position not found")
}

func (s *server) CreateStream(ctx context.Context, request *queue.CreateStreamRequest) (*queue.CreateStreamResponse, error) {
	location := s.location.StreamLocationFolder(request.Name)
	if _, err := os.Stat(location); !os.IsNotExist(err) {
		return nil, errors.New("stream already exists")
	}
	os.Mkdir(location, os.ModePerm)

	err := ioutil.WriteFile(s.location.StreamInfoFile(request.Name), []byte(strings.Replace(streamInfoPayload, "{partitionCount}", strconv.Itoa(int(request.PartitionCount)), 1)), 0644)
	if err != nil {
		log.Fatalf("%s: %s", "Error sabing streamInfo", err.Error())
	}

	partitions := make([]partition, 0)
	s.streamsMutex[request.Name] = partitions
	for i := 0; i < int(request.PartitionCount); i++ {
		l := s.location.StreamPartitionFolder(request.Name, i)
		os.Mkdir(l, os.ModePerm)
		payloadReplacer := strings.NewReplacer("{consume-group}", consumerGroup,
			"{logFile}", "0.log",
			"{entriesCounter}", "0",
			"{byteoffset}", "0")
		data := payloadReplacer.Replace(headPositionPaylod)

		err := ioutil.WriteFile(s.location.StreamHeadPositionFile(request.Name, i), []byte(data), 0644)
		if err != nil {
			log.Fatalf("%s: %s", "Error saving StreamHeadPosition", err.Error())
		}
		s.streamsMutex[request.Name] = append(s.streamsMutex[request.Name], partition{logs: make([]*logfile, 0)})
		s.createLogFile(request.Name, i)
	}

	return &queue.CreateStreamResponse{}, nil
}

func (s *server) GetStreams(ctx context.Context, request *queue.GetStreamsRequest) (*queue.GetStreamsResponse, error) {
	result := make([]*queue.Stream, 0, len(s.streamsMutex))

	for key := range s.streamsMutex {
		i, err := s.getStreamPartitionSize(request.Stream.Name)

		if err != nil {
			return nil, err
		}

		result = append(result, &queue.Stream{Name: key, Partition: int32(i)})
	}
	return &queue.GetStreamsResponse{Message: &queue.GetStreamsResponse_Streams{Streams: &queue.Streams{Streams: result}}}, nil
}

func (s *server) Push(ctx context.Context, request *queue.PushItemRequest) (*queue.PushItemResponse, error) {
	partitionNumber := int(request.Stream.Partition) - 1

	partition, err := s.retrievePartition(request.Stream.Name, partitionNumber)
	if err != nil {
		return nil, err
	}
	partition.pushMutex.Lock()
	defer partition.pushMutex.Unlock()

	logFile := s.candidateLogFile(request.Stream.Name, partitionNumber, int64(len(request.Item.Payload)+5))

	request.Item.Payload = append(request.Item.Payload, []byte("\n")...)
	s.writeEntry(request.Stream.Name, logFile, request.Item.Payload, partitionNumber)

	return &queue.PushItemResponse{}, nil
}

func (s *server) Pop(request *queue.PopItemRequest, service queue.QueueService_PopServer) error {

	partitionNumber := int(request.Stream.Partition) - 1

	partition, err := s.retrievePartition(request.Stream.Name, partitionNumber)
	if err != nil {
		return err
	}
	partition.popMutex.Lock()
	defer partition.popMutex.Unlock()

	currentHead, err := s.getHeadPostion(consumerGroup, request.Stream.Name, partitionNumber)
	if err != nil {
		return errors.New("Combination of stream and partition not found")
	}

	totalRead := 0
	callBack := func(bytes []byte) {
		totalRead++
		service.Send(&queue.PopItemResponse{Message: &queue.PopItemResponse_Item{&queue.Item{Payload: bytes[:len(bytes)-1]}}})
	}
	currentFile := currentHead.logFile
	currentOffset := currentHead.offset

	for {

		l := s.location.StreamParttionLogEntryFile(request.Stream.Name, currentFile, partitionNumber)
		file, err := os.Open(l)
		if err != nil {
			return err
		}
		bufferedReader := bufio.NewReaderSize(file, readBufferSize)

		nextOffset, err := fetch(currentOffset, int(request.Quantity), bufferedReader, callBack)

		if totalRead == int(request.Quantity) || err != io.EOF {
			currentOffset = nextOffset
			break
		}

		if err == io.EOF {
			candidateFile := incrementFilePath(currentFile)
			candidateOffset := int64(0)

			candidateLocation := s.location.StreamParttionLogEntryFile(request.Stream.Name, candidateFile, partitionNumber)

			if _, err := os.Stat(candidateLocation); os.IsNotExist(err) {
				file.Close()
				return io.EOF
			} else {
				currentFile = candidateFile
				currentOffset = candidateOffset
			}
		}

		file.Close()
	}
	currentHead.logFile = currentFile
	currentHead.offset = currentOffset
	s.saveHeadPostion(request.Stream.Name, partitionNumber, currentHead)

	return nil
}

func incrementFilePath(path string) string {
	n, _ := strconv.Atoi(strings.Split(path, ".log")[0])
	n++
	return strconv.Itoa(n) + ".log"
}

func fetch(offset int64, limit int, reader io.Reader, callBack func([]byte)) (currentOffset int64, e error) {
	err := alternativeSeek(reader, offset)
	if err != nil {
		return 0, err
	}
	currentPosition := offset
	for i := 0; i < limit; i++ {
		headerContentBuffer := make([]byte, headerMessageLength)
		var currentReaded int
		currentReaded, err := reader.Read(headerContentBuffer)
		if err != nil {
			return currentPosition, err
		}
		sizePayloadBuf := int(binary.LittleEndian.Uint32(headerContentBuffer))
		payloadBuffer := make([]byte, sizePayloadBuf)
		n, err := reader.Read(payloadBuffer)
		if err != nil {
			return currentPosition, err
		}
		currentReaded += n
		callBack(payloadBuffer)
		currentPosition += int64(currentReaded)
	}

	return currentPosition, nil
}

func (s *server) retrievePartition(streamName string, partition int) (*partition, error) {
	if stream, ok := s.streamsMutex[streamName]; ok {
		if len(stream)-1 >= partition {
			return &stream[partition], nil
		}
		return nil, errors.New("partition number not found")

	}
	return nil, errors.New("stream name not found")
}

//alternativeSeek because bufio does not implement seek
func alternativeSeek(reader io.Reader, discard int64) error {
	var counter int
	for {
		var bytes []byte
		if (counter + dicardBufferSize) > int(discard) {
			bytes = make([]byte, int(math.Abs(float64(counter)-float64(discard))))
		} else if (counter + dicardBufferSize) == int(discard) {
			bytes = make([]byte, dicardBufferSize)
		} else {
			bytes = make([]byte, discard-int64(counter))
		}

		n, err := reader.Read(bytes)
		if err != nil {
			return err
		}

		counter += n
		if counter > int(discard) {
			return errors.New("should not be higher ")
		}
		if counter == int(discard) {
			break
		}
	}
	return nil
}
