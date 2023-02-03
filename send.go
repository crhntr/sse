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
	for _, s := range [...]string{
		"id: ", strconv.Itoa(msgNumber), "\n",
		"event: ", string(event), "\n",
		"data: ", data, "\n\n",
	} {
		if _, err := buf.WriteString(s); err != nil {
			return err
		}
	}
	return nil
}
