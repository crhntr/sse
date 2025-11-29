# Contributing to SSE

Thank you for your interest in contributing to the Server-Sent Events (SSE) utility library!

## Getting Started

This project is a Go module that provides utilities for working with Server-Sent Events in Go.

### Project Overview

This is a small, focused library with minimal dependencies (stdlib only). The codebase consists of:

- `send.go` - Core SSE message formatting and sending logic
- `source.go` - EventSource type for managing SSE connections with thread-safe operations
- `error.go` - Custom error types
- `flusher.go` - WriteFlusher interface implementations

Key concepts:
- SSE messages are formatted according to the [W3C Server-Sent Events specification](https://html.spec.whatwg.org/multipage/server-sent-events.html)
- Each event has an ID, event type, and data field
- Multi-line messages are supported (each line prefixed with "data: ")
- The library requires http.ResponseWriter to implement http.Flusher for streaming

### Prerequisites

- Go (see go.mod for the version)
- Git

### Setting Up Your Development Environment

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/sse.git
   cd sse
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/crhntr/sse.git
   ```

## Development Workflow

### Making Changes

1. Create a new branch for your work:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes to the codebase

3. Run tests to ensure everything works:
   ```bash
   go test ./...
   ```

4. Format your code:
   ```bash
   go fmt ./...
   ```

5. Run static analysis:
   ```bash
   go vet ./...
   ```

### Commit Guidelines

- Write clear, concise commit messages (https://www.conventionalcommits.org)
- Use present tense ("feat: do thing" not "Did thing")
- Reference issues and pull requests where appropriate

### Submitting Changes

1. Push your changes to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

2. Open a Pull Request against the `main` branch

3. Ensure your PR description clearly describes:
   - What problem it solves
   - How it solves it
   - Any breaking changes

## Code Standards

### General Guidelines

- Follow standard Go conventions and idioms
- Add comments for exported functions, types, and methods
- Keep functions focused and concise
- Write tests for new functionality
- No external dependencies - stdlib only
- Maintain thread-safety where appropriate (see EventSource.Send)

### Architecture Principles

- **Simplicity**: This is a focused utility library. Keep it simple and avoid feature bloat
- **Performance**: Reuse buffers where possible (see `bytes.Buffer` usage in `Send` and `EventSource`)
- **Correctness**: Follow the SSE specification strictly
- **Compatibility**: Maintain backward compatibility; any breaking changes must be clearly documented

### Common Patterns

When working on this codebase, note these patterns:

1. **Buffer reuse**: Functions accept `*bytes.Buffer` to allow caller-controlled buffer reuse
2. **Interface composition**: `WriteFlusher` combines `io.Writer` and `http.Flusher`
3. **Error handling**: Return errors immediately; no panics in library code
4. **Thread safety**: `EventSource` uses `sync.Mutex` to protect concurrent access to internal state

## Testing

- Add tests for any new features or bug fixes
- Ensure all tests pass before submitting a PR
- Aim for clear, readable test names that describe what they test

### Testing Notes

Currently, the test suite is minimal. When adding tests:
- Test edge cases (empty strings, multi-line data, concurrent access)
- Verify SSE message format correctness
- Test error conditions (nil pointers, failed writes, type assertions)
- Consider using `httptest.ResponseRecorder` for testing HTTP handlers

## For AI Coding Agents

If you're an AI coding agent working on this repository, here are some specific guidelines:

### Understanding the Codebase

1. **Start with the interfaces**: The core abstraction is `WriteFlusher` (send.go:26-29)
2. **Key entry points**:
   - `NewEventSource()` - Creates a new SSE connection (source.go:19)
   - `Send()` - Low-level function to send a single event (send.go:31)
   - `EventSource.Send()` - Thread-safe method for sending events (source.go:32)
   - `EventSource.SendJSON()` - Convenience method for JSON serialization (source.go:44)

3. **Message format**: Study `WriteEventString()` (send.go:52-76) to understand how SSE messages are constructed

### Common Tasks

**Adding a new feature:**
- Ensure it aligns with the SSE specification
- Keep the API minimal and focused
- Add appropriate documentation
- Consider thread-safety implications
- Avoid adding external dependencies

**Fixing bugs:**
- Reproduce the issue first. Create a replication using a main.go file in cmd/try/some-bug/main.go.
- Check if it's a specification compliance issue
- Ensure the fix doesn't break existing behavior
- Add a test case that would catch this bug

**Adding tests:**
- The project currently has no `*_test.go` files. Add them as changes are made. Make sure to cover existing behavior.
- When creating the test suite, cover:
  - Message formatting (single-line, multi-line, special characters)
  - Thread safety of EventSource
  - Error conditions
  - Buffer reuse patterns
  - HTTP header setting

### CI/CD Context

- GitHub Actions workflow runs on push/PR to main
- Tests are currently commented out in `.github/workflows/go.yml`
- Build must succeed for all changes

### Current Limitations to Be Aware Of

- No test coverage yet
- EventSource doesn't handle connection drops or client disconnects
- No built-in retry logic (clients should implement this)

## Questions?

Feel free to open an issue for:
- Bug reports
- Feature requests
- Questions about the codebase

## License

By contributing, you agree that your contributions will be licensed under the MIT License.