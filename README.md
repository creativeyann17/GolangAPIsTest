# GolangAPIsTest

A performance comparison project for popular Go web frameworks, all implementing the same simple "Hello World" API endpoint.

## Frameworks Tested

1. **Fiber** - Express-inspired web framework built on top of FastHTTP
2. **Gin** - High-performance HTTP web framework with a Martini-like API
3. **Echo** - High performance, minimalist Go web framework
4. **Chi** - Lightweight, idiomatic router for building Go HTTP services
5. **HttpRouter** - High-performance HTTP request router with minimal overhead
6. **Hertz** - CloudWeGo's high-performance HTTP framework from ByteDance
7. **FastHTTP** - Fast HTTP implementation for Go (lower-level than other frameworks)

## Project Structure

```
GolangAPIsTest/
├── cmd/                    # Server implementations
│   ├── fiber/
│   ├── gin/
│   ├── echo/
│   ├── chi/
│   ├── httprouter/
│   ├── hertz/
│   └── fasthttp/
├── benchmark/              # Benchmark tests
│   └── benchmark_test.go
├── Makefile               # Build and run commands
├── go.mod
└── README.md
```

## Requirements

- Go 1.25+ (tested with Go 1.25.1)
- Make (optional, but recommended)

## Installation

Clone the repository and install dependencies:

```bash
cd GolangAPIsTest
make deps
```

Or manually:

```bash
go mod download
go mod tidy
```

## Running Individual Servers

Each framework can be run independently on port 8080:

```bash
make fiber       # Run Fiber server
make gin         # Run Gin server
make echo        # Run Echo server
make chi         # Run Chi server
make httprouter  # Run HttpRouter server
make hertz       # Run Hertz server
make fasthttp    # Run FastHTTP server
```

All servers expose the same endpoint:

```bash
curl http://localhost:8080/hello
```

Response:
```json
{
  "message": "Hello, World!",
  "framework": "FrameworkName"
}
```

## Running Benchmarks

To benchmark all frameworks:

```bash
make bench
```

The benchmark will:
1. Start each server sequentially
2. Perform 100 warmup requests
3. Execute 1,000 requests to the `/hello` endpoint
4. Measure mean, min, and max response times
5. Calculate success rate
6. Display a summary table

Example output:
```
================================================================================
BENCHMARK SUMMARY
================================================================================
Framework       | Mean Time    | Min Time     | Max Time     | Success Rate
--------------------------------------------------------------------------------
Fiber           | 234µs        | 156µs        | 1.2ms        | 100.00%
Gin             | 245µs        | 167µs        | 1.3ms        | 100.00%
Echo            | 256µs        | 178µs        | 1.4ms        | 100.00%
Chi             | 267µs        | 189µs        | 1.5ms        | 100.00%
HttpRouter      | 223µs        | 145µs        | 1.1ms        | 100.00%
Hertz           | 212µs        | 134µs        | 1.0ms        | 100.00%
FastHTTP        | 198µs        | 123µs        | 0.9ms        | 100.00%
================================================================================
```

## Benchmark Configuration

Edit `benchmark/benchmark_test.go` to adjust:
- `numRequests` - Number of requests per framework (default: 10000)
- `warmupReqs` - Number of warmup requests (default: 100)
- `port` - Server port (default: :8080)

## API Endpoint

All frameworks implement the same endpoint:

- **GET** `/hello`
  - Returns: `{"message": "Hello, World!", "framework": "FrameworkName"}`
  - Status: 200 OK
  - Content-Type: application/json

## Development

Run all tests (including benchmarks):

```bash
make test
```

Clean build artifacts:

```bash
make clean
```

View all available commands:

```bash
make help
```

## Notes

- All servers run on the same port (8080), so only one can run at a time
- The benchmark automatically starts and stops each server
- Results may vary based on system load and hardware
- Warmup requests are performed before benchmarking to ensure fair comparison
- All frameworks use their production/release modes where applicable

## Framework Documentation

- [Fiber](https://docs.gofiber.io/)
- [Gin](https://gin-gonic.com/)
- [Echo](https://echo.labstack.com/)
- [Chi](https://go-chi.io/)
- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [Hertz](https://www.cloudwego.io/docs/hertz/)
- [FastHTTP](https://github.com/valyala/fasthttp)
