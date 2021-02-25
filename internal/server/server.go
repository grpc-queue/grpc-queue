package server

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/grpc-queue/grpc-queue/pkg/grpc/v1/queue"
)

const (
	maxMessageBytes                = 4
	maxEntriesPerLogFile           = 10
	headPositionPaylod             = "{consume-group}|{logFile}|{byteoffset}"
	patitionInfoPayload            = "{lastLog}|{entryCount}"
	streamInfoPayload              = "{partitionCount}"
	positionPattern                = `(\w+)\|(\d+\.log)\|(\d+)`
	partitionInfoPattern           = `(\d+\.log)\|(\d+)`
	streamsLocation                = "/queuedata/streams"
	streamLocation                 = "/queuedata/streams/{streamName}"
	streamInfoLocation             = "/queuedata/streams/{streamName}/stream.info"
	streamPartitionLocation        = "/queuedata/streams/{streamName}/partition{partition}"
	streamPartitionInfoLocation    = "/queuedata/streams/{streamName}/partition{partition}/partition.info"
	streamParttionLogEntryLocation = "/queuedata/streams/{streamName}/partition{partition}/{logFile}"
	streamHeadPositionLocation     = "/queuedata/streams/{streamName}/partition{partition}/head.position"
)

type server struct {
	streamsMutex map[string]*sync.Mutex
	baseDataPath string
	*queue.UnimplementedQueueServiceServer
}

type headPosition struct {
	consumerGroup string
	logFile       string
	offset        int
}

type partitionInfo struct {
	lastLog    string
	EntryCount int
}

func NewServer(dataPath string) *server {
	location := buildPath(dataPath, streamsLocation, nil)
	os.MkdirAll(location, os.ModePerm)
	files, err := ioutil.ReadDir(location)
	if err != nil {
		log.Fatal(err)
	}

	streams := make(map[string]*sync.Mutex)
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".info") {
			streams[f.Name()] = &sync.Mutex{}
		}
	}
	return &server{streamsMutex: streams, baseDataPath: strings.TrimRight(dataPath, "/")}
}

func buildPath(basePath, resource string, replacer *strings.Replacer) string {
	if replacer != nil {
		return basePath + replacer.Replace(resource)
	}
	return basePath + resource
}
func (s *server) savePartitionInfo(streamName string, partition int, p *partitionInfo) {
	replacer := strings.NewReplacer("{streamName}", streamName, "{partition}", strconv.Itoa(partition))
	file, err := os.OpenFile(buildPath(s.baseDataPath, streamPartitionInfoLocation, replacer), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
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
	replacer := strings.NewReplacer("{streamName}", streamName)
	file, err := os.Open(buildPath(s.baseDataPath, streamInfoLocation, replacer))
	if err != nil {
		log.Fatal(err)
		return 0, errors.New("no partition found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	i, _ := strconv.Atoi(scanner.Text())
	return i, nil

}
func (s *server) getPartitionInfo(streamName string, partition int) *partitionInfo {
	replacer := strings.NewReplacer("{streamName}", streamName, "{partition}", strconv.Itoa(partition))
	file, err := os.Open(buildPath(s.baseDataPath, streamPartitionInfoLocation, replacer))
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
	replacer := strings.NewReplacer("{streamName}", streamName,
		"{partition}", strconv.Itoa(partition),
		"{logFile}", p.lastLog)

	location := buildPath(s.baseDataPath, streamParttionLogEntryLocation, replacer)
	file, err := os.OpenFile(location, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	var buffer bytes.Buffer
	lengthMessage := len(message)
	b := make([]byte, maxMessageBytes)
	binary.LittleEndian.PutUint32(b, uint32(lengthMessage))
	buffer.Write(b)
	buffer.Write(message)
	buffer.WriteString("\n")
	file.Write(buffer.Bytes())

}
func (s *server) saveHeadPostion(streamName string, partition int, h *headPosition) {
	replacer := strings.NewReplacer("{streamName}", streamName, "{partition}", strconv.Itoa(partition))
	file, err := os.Open(buildPath(s.baseDataPath, streamHeadPositionLocation, replacer))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	payloadReplacer := strings.NewReplacer("{consume-group}", h.consumerGroup,
		"{logFile}", h.logFile,
		"{byteoffset}", strconv.Itoa(h.offset))
	data := payloadReplacer.Replace(headPositionPaylod)
	err = ioutil.WriteFile(buildPath(s.baseDataPath, streamHeadPositionLocation, replacer), []byte(data), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
func (s *server) ack(streamName string, paritionNumber int, currentHead *headPosition) error {
	futureHead := &headPosition{consumerGroup: currentHead.consumerGroup, logFile: currentHead.logFile, offset: currentHead.offset}
	p := s.getPartitionInfo(streamName, paritionNumber)
	if futureHead.offset+1 == maxEntriesPerLogFile {
		lastI, _ := strconv.Atoi(strings.Split(p.lastLog, ".")[0])

		currentI, _ := strconv.Atoi(strings.Split(futureHead.logFile, ".")[0])

		if currentI == lastI {
			return errors.New("queue is empty")
		}

		currentI++
		futureHead.logFile = strconv.Itoa(currentI) + ".log"
		futureHead.offset = 0
	} else {
		futureHead.offset++
	}
	s.saveHeadPostion(streamName, paritionNumber, futureHead)
	return nil
}
func (s *server) getHeadPostion(consumerGroup, streamName string, partition int) (*headPosition, error) {
	replacer := strings.NewReplacer("{streamName}", streamName, "{partition}", strconv.Itoa(partition))
	file, err := os.Open(buildPath(s.baseDataPath, streamHeadPositionLocation, replacer))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(positionPattern)
	for scanner.Scan() {
		groups := re.FindStringSubmatch(scanner.Text())
		if groups[1] == consumerGroup {
			offset, _ := strconv.Atoi(groups[3])
			return &headPosition{consumerGroup: groups[1], logFile: groups[2], offset: offset}, nil
		}

	}
	return nil, errors.New("Head position not found")
}

func (s *server) readEntry(streamName string, paritionNumber int, h *headPosition) (*string, error) {
	logLocationReplacer := strings.NewReplacer("{streamName}", streamName,
		"{partition}", strconv.Itoa(paritionNumber),
		"{logFile}", h.logFile)
	location := buildPath(s.baseDataPath, streamParttionLogEntryLocation, logLocationReplacer)
	file, err := os.Open(location)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if i == h.offset {
			bytes := scanner.Bytes()
			message := string(bytes[maxMessageBytes:])
			return &message, nil
		}
		i++
	}
	return nil, errors.New("partition empty")

}

func (s *server) CreateStream(ctx context.Context, request *queue.CreateStreamRequest) (*queue.CreateStreamResponse, error) {

	streamLocationReplacer := strings.NewReplacer("{streamName}", request.Name)
	location := buildPath(s.baseDataPath, streamLocation, streamLocationReplacer)
	if _, err := os.Stat(location); !os.IsNotExist(err) {
		return nil, errors.New("stream already exists")
	}
	os.Mkdir(location, os.ModePerm)

	err := ioutil.WriteFile(buildPath(s.baseDataPath, streamInfoLocation, streamLocationReplacer), []byte(strings.Replace(streamInfoPayload, "{partitionCount}", strconv.Itoa(int(request.PartitionCount)), 1)), 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < int(request.PartitionCount); i++ {
		replacer := strings.NewReplacer("{streamName}", request.Name, "{partition}", strconv.Itoa(i))
		os.Mkdir(buildPath(s.baseDataPath, streamPartitionLocation, replacer), os.ModePerm)
		payloadReplacer := strings.NewReplacer("{consume-group}", "main",
			"{logFile}", "0.log",
			"{entriesCounter}", "0",
			"{byteoffset}", "0")
		data := payloadReplacer.Replace(headPositionPaylod)

		headPositionLocationReplacer := strings.NewReplacer("{streamName}", request.Name, "{partition}", strconv.Itoa(i))

		err := ioutil.WriteFile(buildPath(s.baseDataPath, streamHeadPositionLocation, headPositionLocationReplacer), []byte(data), 0644)
		if err != nil {
			log.Fatal(err)
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
	currentHead, _ := s.getHeadPostion("main", request.Stream.Name, int(request.Stream.Partition)-1)
	message, err := s.readEntry(request.Stream.Name, int(request.Stream.Partition)-1, currentHead)
	if err != nil {
		return err
	}

	err = s.ack(request.Stream.Name, int(request.Stream.Partition)-1, currentHead)
	if err != nil {
		return err
	}
	service.Send(&queue.PopItemResponse{Message: &queue.PopItemResponse_Item{&queue.Item{Payload: []byte(*message)}}})

	return nil
}
