package server

import (
	"bytes"
	"encoding/binary"
	"io"
	"strconv"
	"testing"
)

func TestFetch(t *testing.T) {

	t.Run("2 queue items|10 byte offset| limit 1 |pop 1", func(t *testing.T) {
		results := make([]string, 0)
		callback := func(bytes []byte) {
			results = append(results, string(bytes[:len(bytes)-1])) //exclude \n
		}
		var buffer bytes.Buffer
		for i := 0; i < 2; i++ {
			str := "test" + strconv.Itoa(i) + "\n"
			message := []byte(str)
			lengthMessage := len(message)
			b := make([]byte, headerMessageLength)
			binary.LittleEndian.PutUint32(b, uint32(lengthMessage))
			buffer.Write(b)
			buffer.Write(message)
		}
		b := buffer.Bytes()
		reader := bytes.NewReader(b)
		fetch(10, 1, reader, callback)

		if len(results) != 1 || results[0] != "test1" {
			t.Error("expected result: test1")
		}
	})
	t.Run("10 queue items|50 byte offset| limit 5| pop 5", func(t *testing.T) {
		results := make([]string, 0)
		callback := func(bytes []byte) {
			results = append(results, string(bytes[:len(bytes)-1]))
		}

		payloads := []string{"test0", "test1", "test2", "test3", "test4",
			"test5", "test6", "test7", "test8", "test9"}

		expectedOutput := []string{"test5",
			"test6", "test7", "test8", "test9"}

		var buffer bytes.Buffer
		for _, p := range payloads {
			message := []byte(p + "\n")
			lengthMessage := len(message)
			b := make([]byte, headerMessageLength)
			binary.LittleEndian.PutUint32(b, uint32(lengthMessage))
			buffer.Write(b)
			buffer.Write(message)
		}
		b := buffer.Bytes()
		reader := bytes.NewReader(b)

		wantNextOffset := int64(100)
		gotNextOffset, err := fetch(50, 5, reader, callback)

		if err != nil {
			t.Errorf("got %v want %v ", err, nil)
		}
		if gotNextOffset != wantNextOffset {
			t.Errorf("[nextOffset] got %v want %v ", gotNextOffset, wantNextOffset)
		}

		for idx, r := range results {
			if r != expectedOutput[idx] {
				t.Errorf("at[%v] got %s want %s ", idx, r, expectedOutput[idx])
			}
		}
	})
	t.Run("10 queue items|50 byte offset| limit 6| pop 5", func(t *testing.T) {
		results := make([]string, 0)
		callback := func(bytes []byte) {
			results = append(results, string(bytes[:len(bytes)-1]))
		}

		payloads := []string{"test0", "test1", "test2", "test3", "test4",
			"test5", "test6", "test7", "test8", "test9"}

		expectedOutput := []string{"test5",
			"test6", "test7", "test8", "test9"}

		var buffer bytes.Buffer
		for _, p := range payloads {
			message := []byte(p + "\n")
			lengthMessage := len(message)
			b := make([]byte, headerMessageLength)
			binary.LittleEndian.PutUint32(b, uint32(lengthMessage))
			buffer.Write(b)
			buffer.Write(message)
		}
		b := buffer.Bytes()
		reader := bytes.NewReader(b)

		wantNextOffset := int64(100)
		gotNextOffset, err := fetch(50, 6, reader, callback)

		if gotNextOffset != wantNextOffset {
			t.Errorf("[nextOffset] got %v want %v ", gotNextOffset, wantNextOffset)
		}

		if err != io.EOF {
			t.Errorf("got %v want %v ", err, io.EOF)
		}

		for idx, r := range results {
			if r != expectedOutput[idx] {
				t.Errorf("at[%v] got %s want %s ", idx, r, expectedOutput[idx])
			}
		}
	})
	t.Run("10 queue items|50 byte offset| limit 4| pop 4", func(t *testing.T) {
		results := make([]string, 0)
		callback := func(bytes []byte) {
			results = append(results, string(bytes[:len(bytes)-1]))
		}

		payloads := []string{"test0", "test1", "test2", "test3", "test4",
			"test5", "test6", "test7", "test8", "test9"}

		expectedOutput := []string{"test5",
			"test6", "test7", "test8"}

		var buffer bytes.Buffer
		for _, p := range payloads {
			message := []byte(p + "\n")
			lengthMessage := len(message)
			b := make([]byte, headerMessageLength)
			binary.LittleEndian.PutUint32(b, uint32(lengthMessage))
			buffer.Write(b)
			buffer.Write(message)
		}
		b := buffer.Bytes()
		reader := bytes.NewReader(b)

		wantNextOffset := int64(90)
		gotNextOffset, err := fetch(50, 4, reader, callback)

		if err != nil {
			t.Errorf("got %v want %v ", err, nil)
		}

		if gotNextOffset != wantNextOffset {
			t.Errorf("[nextOffset] got %v want %v ", gotNextOffset, wantNextOffset)
		}

		for idx, r := range results {
			if r != expectedOutput[idx] {
				t.Errorf("at[%v] got %s want %s ", idx, r, expectedOutput[idx])
			}
		}
	})

}
