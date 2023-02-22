package sse

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"sync"
)

type EventSource struct {
	mut sync.Mutex
	id  int
	buf *bytes.Buffer
	res WriteFlusher
}

func NewEventSource(res http.ResponseWriter, statusCode int) (*EventSource, error) {
	wf, ok := res.(WriteFlusher)
	if !ok {
		return nil, ErrorResponseWriterDoesNotImplementFlusher{}
	}
	SetHeaders(res)
	res.WriteHeader(statusCode)
	return &EventSource{
		buf: bytes.NewBuffer(make([]byte, 0, 1024)),
		res: wf,
	}, nil
}

func (src *EventSource) Send(event EventName, message string) (int, error) {
	src.mut.Lock()
	defer src.mut.Unlock()
	w := src.res
	if w == nil {
		w = NewNoopWriteFlusher(io.Discard)
	}
	src.id++
	_, err := Send(src.res, src.buf, src.id, event, message)
	return src.id, err
}

func (src *EventSource) SendJSON(event EventName, data any) (int, error) {
	if event == "" {
		event = EventName(reflect.TypeOf(data).Name())
	}
	buf, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	return src.Send(event, string(buf))
}

func (src *EventSource) LastEventID() int {
	src.mut.Lock()
	defer src.mut.Unlock()
	return src.id
}
