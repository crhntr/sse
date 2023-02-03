package sse

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
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

func Send(res WriteFlusher, buf *bytes.Buffer, msgNumber int, event EventName, data string) (int, error) {
	if buf == nil {
		buf = bytes.NewBuffer(make([]byte, 0, 1024))
	} else {
		buf.Reset()
	}
	if msgNumber <= 0 {
		msgNumber = 1
	}
	writeMessageString(buf, msgNumber, event, data)
	n, err := io.Copy(res, buf)
	if err != nil {
		return int(n), err
	}
	res.Flush()
	return int(n), nil
}

func writeMessageString(buf *bytes.Buffer, msgNumber int, event EventName, data string) {
	_, _ = buf.WriteString("id: ")
	_, _ = buf.WriteString(strconv.Itoa(msgNumber))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString("event: ")
	_, _ = buf.WriteString(string(event))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString("data: ")
	_, _ = buf.WriteString(data)
	_, _ = buf.WriteString("\n\n")
}
