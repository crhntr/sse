package sse

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func NewEventSource(res http.ResponseWriter) (*EventSource, error) {
	wf, ok := res.(WriteFlusher)
	if !ok {
		return nil, fmt.Errorf("http.ResponseWriter does not implement sse.WriteFlusher")
	}
	return &EventSource{
		buf: bytes.NewBuffer(make([]byte, 0, 1024)),
		res: wf,
	}, nil
}

func (src *EventSource) Send(event EventName, message string) error {
	src.mut.Lock()
	defer src.mut.Unlock()
	w := src.res
	if w == nil {
		w = NewNoopWriteFlusher(io.Discard)
	}
	src.id++
	_, err := Send(src.res, src.buf, src.NextEventID(), event, message)
	return err
}

func (src *EventSource) SendJSON(event EventName, data any) error {
	if event == "" {
		event = EventName(reflect.TypeOf(data).Name())
	}
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return src.Send(event, string(buf))
}

func (src *EventSource) NextEventID() int { return src.id }
