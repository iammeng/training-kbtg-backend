# Training KBTG Backend

A simple REST API built with Go and Fiber framework.

## Features

- RESTful API server built with Go
- Uses Fiber web framework
- CORS and Logger middleware included
- Returns JSON response for GET / endpoint

## Quick Start

1. Install dependencies:
```bash
go mod tidy
```

2. Run the server:
```bash
go run main.go
```

3. Test the API:
```bash
curl http://localhost:3000/
```

Expected response:
```json
{"message": "hello world"}
```

## API Endpoints

- `GET /` - Returns a hello world message

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- Go 1.21+

## Port

The server runs on port 3000 by default.
