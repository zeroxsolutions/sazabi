package sazabi

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/zeroxsolutions/barbatos/log"
)

// Environment constants for logging configuration.
const (
	ProductionEnvName      = "production" // Full name for production environment
	ProductionEnvShortName = "prod"       // Short name for production environment
)

// logger represents the global logger instance used throughout the application.
var (
	logger log.Logger
)

// Initialize sets up the logger based on the specified environment.
// It configures the logger for production or development mode.
// In production, it uses a specific configuration to manage log levels and formats.
// If an error occurs during logger initialization, the application panics.
func Initialize(environment string) {
	var conf zap.Config
	conf = newProductionConfig()

	if environment != ProductionEnvName && environment != ProductionEnvShortName {
		conf = zap.NewDevelopmentConfig()
	}

	conf.DisableStacktrace = true
	log, err := conf.Build()
	if err != nil {
		panic(err) // Panic if logger configuration fails
	}

	logger = log.WithOptions(zap.AddCallerSkip(1)).Sugar() // Set the global logger
}

// newProductionConfig returns a zap.Config configured for production environment.
// It sets the log level to "info", disables development mode, and configures
// sampling and output formatting. Outputs are directed to "stderr".
func newProductionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // Set log level to Info
		Development: false,                               // Disable development mode
		Sampling: &zap.SamplingConfig{
			Initial:    100, // Initial number of logs to sample
			Thereafter: 100, // Subsequent logs to sample
		},
		Encoding:         "console",                    // Use console encoding for output
		EncoderConfig:    newProductionEncoderConfig(), // Configure the encoder
		OutputPaths:      []string{"stderr"},           // Log output to stderr
		ErrorOutputPaths: []string{"stderr"},           // Error output to stderr
	}
}

// newProductionEncoderConfig returns a zapcore.EncoderConfig configured for production environment.
// It defines key names and formats for logging output, including timestamps, levels, and messages.
func newProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",                           // Key for timestamp
		LevelKey:       "level",                        // Key for log level
		NameKey:        "logger",                       // Key for logger name
		CallerKey:      "caller",                       // Key for caller information
		MessageKey:     "msg",                          // Key for log message
		StacktraceKey:  "stacktrace",                   // Key for stack trace
		LineEnding:     zapcore.DefaultLineEnding,      // Line ending configuration
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Encode level as capital letters
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // Encode time in ISO8601 format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Encode duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Shorten caller information
	}
}

// Debug logs debug messages using the global logger.
func Debug(args ...interface{}) {
	logger.Debug(args...) // Log debug message
}

// Debugf logs formatted debug messages using the global logger.
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...) // Log formatted debug message
}

// Debugw logs debug messages with additional key-value pairs for structured logging using the global logger.
func Debugw(msg string, keysValues ...interface{}) {
	logger.Debugw(msg, keysValues...) // Log debug message with structured key-value pairs
}

// Info logs info messages using the global logger.
func Info(args ...interface{}) {
	logger.Info(args...) // Log info message
}

// Infof logs formatted info messages using the global logger.
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...) // Log formatted info message
}

// Infow logs info messages with additional key-value pairs for structured logging using the global logger.
func Infow(msg string, keysValues ...interface{}) {
	logger.Infow(msg, keysValues...) // Log info message with structured key-value pairs
}

// Warn logs warning messages using the global logger.
func Warn(args ...interface{}) {
	logger.Warn(args...) // Log warning message
}

// Warnf logs formatted warning messages using the global logger.
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...) // Log formatted warning message
}

// Warnw logs warning messages with additional key-value pairs for structured logging using the global logger.
func Warnw(msg string, keysValues ...interface{}) {
	logger.Warnw(msg, keysValues...) // Log warning message with structured key-value pairs
}

// Error logs error messages using the global logger.
func Error(args ...interface{}) {
	logger.Error(args...) // Log error message
}

// Errorf logs formatted error messages using the global logger.
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...) // Log formatted error message
}

// Errorw logs error messages with additional key-value pairs for structured logging using the global logger.
func Errorw(msg string, keysValues ...interface{}) {
	logger.Errorw(msg, keysValues...) // Log error message with structured key-value pairs
}

// Fatal logs fatal messages using the global logger.
func Fatal(args ...interface{}) {
	logger.Fatal(args...) // Log fatal message
}

// Fatalf logs formatted fatal messages using the global logger.
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...) // Log formatted fatal message
}

// Fatalw logs fatal messages with additional key-value pairs for structured logging using the global logger.
func Fatalw(msg string, keysValues ...interface{}) {
	logger.Fatalw(msg, keysValues...) // Log fatal message with structured key-value pairs
}

// Panic logs panic messages using the global logger.
func Panic(args ...interface{}) {
	logger.Panic(args...) // Log panic message
}

// Panicf logs formatted panic messages using the global logger.
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...) // Log formatted panic message
}

// Panicw logs panic messages with additional key-value pairs for structured logging using the global logger.
func Panicw(msg string, keysValues ...interface{}) {
	logger.Panicw(msg, keysValues...) // Log panic message with structured key-value pairs
}

// Default creates and returns a default logger configured for development environment.
// It disables stack traces and panics if there's an error while building the logger.
func Default() log.Logger {
	var conf zap.Config = zap.NewDevelopmentConfig() // Set up development logger configuration

	conf.DisableStacktrace = true // Disable stack trace for development logger
	log, err := conf.Build()
	if err != nil {
		panic(err) // Panic if logger configuration fails
	}

	return log.WithOptions(zap.AddCallerSkip(1)).Sugar() // Return a sugar logger with caller information
}
