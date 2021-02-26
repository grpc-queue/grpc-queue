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
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/grpc-queue/grpc-queue/internal/location"
	"github.com/grpc-queue/grpc-queue/pkg/grpc/v1/queue"
)

const (
	maxMessageBytes      = 4
	maxEntriesPerLogFile = 2
	headPositionPaylod   = "{consume-group}|{logFile}|{byteoffset}"
	patitionInfoPayload  = "{lastLog}|{entryCount}"
	streamInfoPayload    = "{partitionCount}"
	positionPattern      = `(\w+)\|(\d+\.log)\|(\d+)`
	partitionInfoPattern = `(\d+\.log)\|(\d+)`
	consumerGroup        = "main"
)

type server struct {
	streamsMutex map[string]*sync.Mutex
	location     *location.Location
	*queue.UnimplementedQueueServiceServer
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

	streams := make(map[string]*sync.Mutex)
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".info") {
			streams[f.Name()] = &sync.Mutex{}
		}
	}
	return &server{streamsMutex: streams, location: location}
}

func (s *server) savePartitionInfo(streamName string, partition int, p *partitionInfo) {
	file, err := os.OpenFile(s.location.StreamPartitionInfoFile(streamName, partition), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	payloadReplacer := strings.NewReplacer("{lastLog}", p.lastLog,
		"{entryCount}", strconv.Itoa(p.EntryCount))
	data := payloadReplacer.Replace(patitionInfoPayload)

	file.Write([]byte(data))
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
func (s *server) getPartitionInfo(streamName string, partition int) *partitionInfo {
	file, err := os.Open(s.location.StreamPartitionInfoFile(streamName, partition))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(partitionInfoPattern)
	scanner.Scan()
	groups := re.FindStringSubmatch(scanner.Text())

	lastLog := groups[1]
	entryCount, _ := strconv.Atoi(groups[2])

	return &partitionInfo{lastLog: lastLog, EntryCount: entryCount}

}

func (s *server) updatePartitionInfo(partition, amount int, streamName string) *partitionInfo {
	p := s.getPartitionInfo(streamName, partition)
	if p.EntryCount == maxEntriesPerLogFile {
		i, _ := strconv.Atoi(strings.Split(p.lastLog, ".")[0])
		i++
		p.lastLog = strconv.Itoa(i) + ".log"
		p.EntryCount = 0
	} else {
		p.EntryCount++
	}
	s.savePartitionInfo(streamName, partition, p)
	return p
}
func (s *server) writeEntry(streamName string, message []byte, partition int, p *partitionInfo) {

	location := s.location.StreamParttionLogEntryFile(streamName, p.lastLog, partition)
	file, err := os.OpenFile(location, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	var buffer bytes.Buffer
	lengthMessage := len(message) + 1 //len of message plus \n
	b := make([]byte, maxMessageBytes)
	binary.LittleEndian.PutUint32(b, uint32(lengthMessage))
	buffer.Write(b)
	buffer.Write(message)
	buffer.WriteString("\n")
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
	re := regexp.MustCompile(positionPattern)
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
		log.Fatalf("%s: %w", "Error sabing streamInfo", err)
	}

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
			log.Fatalf("%s: %w", "Error saving StreamHeadPosition", err)
		}

		s.savePartitionInfo(request.Name, i, &partitionInfo{lastLog: "0.log", EntryCount: 0})
	}
	s.streamsMutex[request.Name] = &sync.Mutex{}

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
	partitionCount, err := s.getStreamPartitionSize(request.Stream.Name)
	if err != nil {
		return nil, err
	}
	if int(request.Stream.Partition)-1 > partitionCount {
		return nil, errors.New("invalid partition")
	}
	s.streamsMutex[request.Stream.Name].Lock()
	defer s.streamsMutex[request.Stream.Name].Unlock()

	p := s.updatePartitionInfo(partitionNumber, 1, request.Stream.Name)
	s.writeEntry(request.Stream.Name, request.Item.Payload, partitionNumber, p)
	return &queue.PushItemResponse{}, nil
}

func (s *server) Pop(request *queue.PopItemRequest, service queue.QueueService_PopServer) error {

	partition := int(request.Stream.Partition) - 1
	currentHead, err := s.getHeadPostion(consumerGroup, request.Stream.Name, partition)
	if err != nil {
		return errors.New("Combination of stream and partition not found")
	}

	totalRead := 0
	callBack := func(bytes []byte) {
		totalRead++
		service.Send(&queue.PopItemResponse{Message: &queue.PopItemResponse_Item{&queue.Item{Payload: bytes}}})
	}
	currentFile := currentHead.logFile
	currentOffset := currentHead.offset

	for {

		l := s.location.StreamParttionLogEntryFile(request.Stream.Name, currentFile, partition)
		file, err := os.Open(l)
		if err != nil {
			return err
		}

		nextOffset, err := fetch(currentOffset, int(request.Quantity), file, callBack)

		if totalRead == int(request.Quantity) || err != io.EOF {
			currentOffset = nextOffset
			break
		}

		if err == io.EOF {
			candidateFile := incrementFilePath(currentFile)
			candidateOffset := int64(0)

			candidateLocation := s.location.StreamParttionLogEntryFile(request.Stream.Name, candidateFile, partition)

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
	s.saveHeadPostion(request.Stream.Name, partition, currentHead)

	return nil
}

func incrementFilePath(path string) string {
	n, _ := strconv.Atoi(strings.Split(path, ".log")[0])
	n++
	return strconv.Itoa(n) + ".log"
}

func fetch(offset int64, limit int, reader io.ReadSeeker, callBack func([]byte)) (currentOffset int64, e error) {
	reader.Seek(offset, io.SeekStart)
	var currentPosition int64
	for i := 0; i < limit; i++ {
		headerContentBuffer := make([]byte, 4)
		_, err := reader.Read(headerContentBuffer)
		if err != nil {
			return currentPosition, err
		}
		sizePayloadBuf := int(binary.LittleEndian.Uint32(headerContentBuffer))
		payloadBuffer := make([]byte, sizePayloadBuf)
		reader.Read(payloadBuffer)
		callBack(payloadBuffer)
		currentPosition, _ = reader.Seek(0, io.SeekCurrent)
	}

	return currentPosition, nil
}
