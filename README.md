# Sazabi

A Go logging library that provides a simplified interface around [Uber's Zap logger](https://github.com/uber-go/zap), with built-in production and development configurations.

## Features

- 🚀 **Production-ready**: Optimized configuration for production environments
- 🛠️ **Development-friendly**: Enhanced logging for development with readable output
- 📊 **Structured logging**: Support for key-value pairs and formatted messages
- 🎯 **Multiple log levels**: Debug, Info, Warn, Error, Fatal, and Panic
- 🔧 **Simple API**: Easy-to-use interface with global logger functions
- ⚡ **High performance**: Built on top of Zap's high-performance logging

## Installation

```shell
go get github.com/zeroxsolutions/sazabi
```

## Quick Start

```go
package main

import "github.com/zeroxsolutions/sazabi"

func main() {
    // Initialize logger for production environment
    sazabi.Initialize("production")
    
    // Or initialize for development
    // sazabi.Initialize("development")
    
    // Start logging
    sazabi.Info("Application started")
    sazabi.Debugf("Debug message with value: %d", 42)
    sazabi.Errorw("An error occurred", "error", "file not found", "path", "/tmp/test.txt")
}
```

## Configuration

### Environment-based Configuration

The logger automatically configures itself based on the environment string passed to `Initialize()`:

- **Production** (`"production"` or `"prod"`): 
  - Log level: Info and above
  - Format: Console encoding with structured output
  - Sampling enabled for performance
  - Output: stderr

- **Development** (any other value):
  - Log level: Debug and above  
  - Format: Human-readable development format
  - Full logging without sampling
  - Stack traces disabled

```go
// Production configuration
sazabi.Initialize("production")  // or "prod"

// Development configuration  
sazabi.Initialize("development") // or any non-production value
```

## API Reference

### Initialization

```go
// Initialize the global logger
sazabi.Initialize(environment string)

// Create a default development logger (without setting global logger)
logger := sazabi.Default()
```

### Logging Functions

All logging functions are available in three variants:

#### Debug Level
```go
sazabi.Debug(args ...interface{})                    // Simple message
sazabi.Debugf(template string, args ...interface{}) // Formatted message  
sazabi.Debugw(msg string, keysValues ...interface{}) // Structured logging
```

#### Info Level
```go
sazabi.Info(args ...interface{})
sazabi.Infof(template string, args ...interface{})
sazabi.Infow(msg string, keysValues ...interface{})
```

#### Warn Level
```go
sazabi.Warn(args ...interface{})
sazabi.Warnf(template string, args ...interface{})
sazabi.Warnw(msg string, keysValues ...interface{})
```

#### Error Level
```go
sazabi.Error(args ...interface{})
sazabi.Errorf(template string, args ...interface{})
sazabi.Errorw(msg string, keysValues ...interface{})
```

#### Fatal Level (exits application)
```go
sazabi.Fatal(args ...interface{})
sazabi.Fatalf(template string, args ...interface{})
sazabi.Fatalw(msg string, keysValues ...interface{})
```

#### Panic Level (panics)
```go
sazabi.Panic(args ...interface{})
sazabi.Panicf(template string, args ...interface{})
sazabi.Panicw(msg string, keysValues ...interface{})
```

## Usage Examples

### Basic Logging
```go
sazabi.Initialize("development")

sazabi.Debug("This is a debug message")
sazabi.Info("Application is running")
sazabi.Warn("This is a warning")
sazabi.Error("An error occurred")
```

### Formatted Logging
```go
name := "John"
age := 30

sazabi.Infof("User %s is %d years old", name, age)
sazabi.Errorf("Failed to process user %s: %v", name, err)
```

### Structured Logging
```go
sazabi.Infow("User login",
    "username", "john_doe",
    "ip", "192.168.1.1",
    "timestamp", time.Now(),
)

sazabi.Errorw("Database connection failed",
    "database", "users",
    "host", "localhost:5432",
    "error", err.Error(),
)
```

### Production vs Development

```go
// Production logging (only Info level and above)
sazabi.Initialize("production")
sazabi.Debug("This won't appear in production")  // Hidden
sazabi.Info("This will appear")                  // Visible

// Development logging (all levels)
sazabi.Initialize("development") 
sazabi.Debug("This will appear in development")  // Visible
sazabi.Info("This will also appear")             // Visible
```

## Testing

Run the test suite:

```shell
# Run tests with the test script
./bin/test.sh

# Or run directly with Go
go test -tags=test -v ./...
```

## Dependencies

- [go.uber.org/zap](https://github.com/uber-go/zap) - High-performance logging library
- [github.com/zeroxsolutions/barbatos](https://github.com/zeroxsolutions/barbatos) - Internal logging interface

## Requirements

- Go 1.18 or higher

## License

This project is part of the ZeroX Solutions ecosystem.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Support

For support and questions, please open an issue in the repository.
