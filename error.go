package sse

type ErrorResponseWriterDoesNotImplementFlusher struct{}

func (ErrorResponseWriterDoesNotImplementFlusher) Error() string {
	return "http.ResponseWriter does not implement sse.WriteFlusher"
}
