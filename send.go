package sse

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type EventName string

const (
	MessageEvent EventName = "message"
	ErrorEvent   EventName = "error"
)

func SetHeaders(res http.ResponseWriter) {
	headers := res.Header()
	headers.Set("content-type", "text/event-stream; charset=utf-8")
	headers.Set("connection", "keep-alive")
}

type WriteFlusher interface {
	io.Writer
	http.Flusher
}

func Send(res WriteFlusher, buf *bytes.Buffer, msgNumber int, event EventName, data string) (int64, error) {
	if buf == nil {
		buf = bytes.NewBuffer(make([]byte, 0, 1024))
	} else {
		buf.Reset()
	}
	if msgNumber <= 0 {
		msgNumber = 1
	}
	err := WriteEventString(buf, msgNumber, event, data)
	if err != nil {
		return 0, err
	}
	n, err := io.Copy(res, buf)
	if err != nil {
		return n, err
	}
	res.Flush()
	return n, nil
}

func WriteEventString(buf io.StringWriter, msgNumber int, event EventName, data string) error {
	if err := writeEachString(buf, []string{
		"id: ", strconv.Itoa(msgNumber), "\n",
		"event: ", string(event), "\n",
	}); err != nil {
		return err
	}
	r := bufio.NewReader(strings.NewReader(data))
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if err := writeEachString(buf, []string{
			"data: ", line, "\n",
		}); err != nil {
			return err
		}
	}
	_, err := buf.WriteString("\n")
	return err
}

func writeEachString(buf io.StringWriter, strs []string) error {
	for _, s := range strs {
		if _, err := buf.WriteString(s); err != nil {
			return err
		}
	}
	return nil
}
