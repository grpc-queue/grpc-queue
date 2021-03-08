package location

import (
	"strconv"
	"strings"
)

const (
	streamsFolder              = "/queuedata/streams"
	streamFolder               = "/queuedata/streams/{streamName}"
	streamInfoFile             = "/queuedata/streams/{streamName}/stream.info"
	streamPartitionFolder      = "/queuedata/streams/{streamName}/partition{partition}"
	streamParttionLogEntryFile = "/queuedata/streams/{streamName}/partition{partition}/{logFile}"
	streamHeadPositionFile     = "/queuedata/streams/{streamName}/partition{partition}/head.position"
)

type Location struct {
	baseDataPath string
}

func NewLocation(basePath string) *Location {

	return &Location{baseDataPath: strings.TrimRight(basePath, "/")}
}

func buildPath(basePath, resource string, replacer *strings.Replacer) string {
	if replacer != nil {
		return basePath + replacer.Replace(resource)
	}
	return basePath + resource
}

func (l *Location) StreamsFolder() string {
	return buildPath(l.baseDataPath, streamsFolder, nil)
}
func (l *Location) StreamInfoFile(streamName string) string {
	replacer := strings.NewReplacer("{streamName}", streamName)
	return buildPath(l.baseDataPath, streamInfoFile, replacer)
}
func (l *Location) StreamParttionLogEntryFile(streamName, logFile string, partition int) string {
	replacer := strings.NewReplacer("{streamName}", streamName,
		"{partition}", strconv.Itoa(partition),
		"{logFile}", logFile)
	return buildPath(l.baseDataPath, streamParttionLogEntryFile, replacer)
}
func (l *Location) StreamHeadPositionFile(streamName string, partition int) string {
	replacer := strings.NewReplacer("{streamName}", streamName, "{partition}", strconv.Itoa(partition))
	return buildPath(l.baseDataPath, streamHeadPositionFile, replacer)
}

func (l *Location) StreamLocationFolder(streamName string) string {

	streamLocationReplacer := strings.NewReplacer("{streamName}", streamName)
	return buildPath(l.baseDataPath, streamFolder, streamLocationReplacer)
}

func (l *Location) StreamPartitionFolder(streamName string, parition int) string {
	replacer := strings.NewReplacer("{streamName}", streamName, "{partition}", strconv.Itoa(parition))
	return buildPath(l.baseDataPath, streamPartitionFolder, replacer)
}
