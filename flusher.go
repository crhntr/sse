package sse

import "io"

type noopWriteFlusher struct {
	io.Writer
}

func (f noopWriteFlusher) Flush() {}

func NewNoopWriteFlusher(w io.Writer) WriteFlusher {
	return noopWriteFlusher{
		Writer: w,
	}
}
