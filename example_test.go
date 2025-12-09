package sse_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/crhntr/sse"
)

func ExampleEventSource() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Create a new EventSource
		source, err := sse.NewEventSource(w, http.StatusOK)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send events
		_, _ = source.Send(sse.MessageEvent, "Hello, World!")
		_, _ = source.Send(sse.MessageEvent, "This is a streaming update")
	}

	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	w := httptest.NewRecorder()
	handler(w, req)
}

func ExampleEventSource_SendJSON() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		source, err := sse.NewEventSource(w, http.StatusOK)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		type Update struct {
			Message string `json:"message"`
			Count   int    `json:"count"`
		}

		// Send JSON events
		_, _ = source.SendJSON(sse.MessageEvent, Update{
			Message: "Processing",
			Count:   1,
		})
	}

	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	w := httptest.NewRecorder()
	handler(w, req)
}

func ExampleSend() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Set SSE headers
		sse.SetHeaders(w)
		w.WriteHeader(http.StatusOK)

		// Create a buffer for reuse
		buf := bytes.NewBuffer(make([]byte, 0, 1024))

		// Send individual events
		_, _ = sse.Send(w.(sse.WriteFlusher), buf, 1, sse.MessageEvent, "First message")
		_, _ = sse.Send(w.(sse.WriteFlusher), buf, 2, sse.MessageEvent, "Second message")
	}

	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	w := httptest.NewRecorder()
	handler(w, req)
}
