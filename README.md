# sse

A lightweight Go library for Server-Sent Events (SSE) that provides a simple, idiomatic API for streaming real-time updates to clients over HTTP.

## Features

- Simple, idiomatic Go API
- Thread-safe event streaming with `EventSource`
- Automatic event ID tracking
- JSON serialization support
- Multi-line data handling
- Proper SSE header management
- No external dependencies

## Installation

```bash
go get github.com/crhntr/sse
```

See [`./go.mod`](./go.mod) for toolchain version requirement.

## Usage

See the [runnable examples](example_test.go) for complete usage demonstrations:

- **[ExampleEventSource](example_test.go)** - Basic usage with EventSource to send text messages
- **[ExampleEventSource_SendJSON](example_test.go)** - Streaming JSON-encoded data
- **[ExampleSend](example_test.go)** - Manual event sending with lower-level API for more control

You can also view these examples on [pkg.go.dev](https://pkg.go.dev/github.com/crhntr/sse#pkg-examples).

## API Reference

For complete API documentation, see [pkg.go.dev](https://pkg.go.dev/github.com/crhntr/sse).

## Client-Side Usage

```javascript
const eventSource = new EventSource('/events');

eventSource.addEventListener('message', (event) => {
    console.log('Received:', event.data);
    console.log('Event ID:', event.lastEventId);
});

eventSource.addEventListener('error', (event) => {
    console.error('Error:', event.data);
});
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](LICENSE) for details.

Copyright (c) 2023 Christopher Hunter