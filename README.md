# Go Logging Library

This logging library is a wrapper around the popular `zap` logging package, providing a simple and extensible logging solution for Go applications. It offers structured logging with context support, making it easy to include additional metadata in your log entries.

## Features

- Structured logging using zap
- ECS (Elastic Common Schema) compatible log output
- Context-aware logging
- Multiple log levels (Debug, Info, Error, Fatal)
- Support for adding custom fields to log entries
- Middleware for extracting and including request information in logs

## Installation

To use this logging library in your Go project, run:

```bash
go get github.com/ahghazey/logging
```

## Usage
### Initializing the Logger
To initialize the logger, use the InitLogger function:

```go
package main 

import "github.com/ahghazey/logging"

err := logging.InitLogger("info", "your-service-name", "production")
if err != nil {
    // Handle error
}
```

The InitLogger function takes three parameters:

* Log level (debug, info, warn, error)
* Service name
* Environment (e.g., production, staging, development)

### Basic Logging
After initialization, you can use the global LogHandle to log messages:

```go
logging.LogHandle.Info("This is an info message")
logging.LogHandle.Debug("This is a debug message")
logging.LogHandle.Error("This is an error message")
```

### Logging with Formatting
The library supports formatted logging:
```go
logging.LogHandle.Infof("User %s logged in", username)
logging.LogHandle.Errorf("Failed to process request: %v", err)
```

### Logging with Fields
You can add custom fields to your log entries:
```go
fields := map[string]string{
    "user_id": "12345",
    "action": "login",
}
logger := logging.LogHandle.WithFields(fields)
logger.Info("User action logged")
```

### Context-Aware Logging
The library supports context-aware logging, which can automatically include request-specific information:
```go
ctx := context.Background()
ctx = context.WithValue(ctx, constant.RequestIdHeader, "req-123")
ctx = context.WithValue(ctx, constant.UserIdHeader, "user-456")

logger := logging.LogHandle.WithContext(ctx)
logger.Info("Processing request")
```

### Middleware
The library includes a middleware for Chi router (v5) that can extract information from the request and store it in the context:

```go
import "middleware github.com/ahghazey/logging/middleware/chi/autoconfig"

r := chi.NewRouter()
r.Use(middleware.ContextHolderMiddlewareChiV5)
```

### Configuration
The library uses the following constants for context keys:

```go
const (
    RequestIdHeader ContextKey = "requestId"
    TraceIdHeader   ContextKey = "traceId"
    UserIdHeader    ContextKey = "userId"
    SpanIdHeader    ContextKey = "spanId"
)
```

### TODO

* Complete the TokenRefactor function to parse the second part of the token.
* Implement the ContextHolderMiddlewareChiV5 middleware to extract information from the request auth header and store it in the context.
