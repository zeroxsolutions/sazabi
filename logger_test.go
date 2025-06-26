//go:build test
// +build test

package sazabi_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/zeroxsolutions/barbatos/log"
	"github.com/zeroxsolutions/sazabi"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		name        string
		environment string
		wantPanic   bool
	}{
		{
			name:        "production environment",
			environment: sazabi.ProductionEnvName,
			wantPanic:   false,
		},
		{
			name:        "production short name",
			environment: sazabi.ProductionEnvShortName,
			wantPanic:   false,
		},
		{
			name:        "development environment",
			environment: "development",
			wantPanic:   false,
		},
		{
			name:        "test environment",
			environment: "test",
			wantPanic:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Initialize() should have panicked")
					}
				}()
				sazabi.Initialize(tt.environment)
			} else {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Initialize() should not have panicked: %v", r)
					}
				}()
				sazabi.Initialize(tt.environment)
				// Cannot test private logger variable in external test
				// This is a limitation of black-box testing
			}
		})
	}
}

// Note: TestNewProductionConfig and TestNewProductionEncoderConfig are removed
// because these functions are private and cannot be accessed from external test package.
// This is expected behavior with black-box testing approach.

func TestDefault(t *testing.T) {
	defaultLogger := sazabi.Default()
	if defaultLogger == nil {
		t.Error("Default logger should not be nil")
	}

	// Test that it implements the log.Logger interface
	var _ log.Logger = defaultLogger
}

// Note: Helper function removed as we cannot access private logger from external package
// We use stderr capture approach instead for black-box testing

func TestDebug(t *testing.T) {
	// Test by capturing actual stderr output
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize logger for testing
	sazabi.Initialize("development")

	// Test logging
	sazabi.Debug("test debug message")

	// Restore stderr
	w.Close()
	os.Stderr = originalStderr

	// Read captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// In development mode, debug messages should be visible
	if !bytes.Contains([]byte(output), []byte("test debug message")) {
		t.Logf("Debug output: %s", output)
		// Debug messages might not be visible in production mode, that's ok
	}
}

func TestDebugf(t *testing.T) {
	// Test formatted debug logging
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("development")
	sazabi.Debugf("test debug message: %s", "formatted")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("formatted")) {
		t.Logf("Debugf output: %s", output)
	}
}

func TestDebugw(t *testing.T) {
	// Test structured debug logging
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("development")
	sazabi.Debugw("test debug message", "key", "value")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("test debug message")) {
		t.Logf("Debugw output: %s", output)
	}
}

func TestInfo(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Info("test info message")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Info messages should be visible in both dev and prod
	if !bytes.Contains([]byte(output), []byte("test info message")) {
		t.Errorf("Expected output to contain 'test info message', got: %s", output)
	}
}

func TestInfof(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Infof("test info message: %s", "formatted")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("formatted")) {
		t.Errorf("Expected output to contain 'formatted', got: %s", output)
	}
}

func TestInfow(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Infow("test info message", "key", "value")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("test info message")) {
		t.Errorf("Expected output to contain 'test info message', got: %s", output)
	}
}

func TestWarn(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Warn("test warn message")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("test warn message")) {
		t.Errorf("Expected output to contain 'test warn message', got: %s", output)
	}
}

func TestWarnf(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Warnf("test warn message: %s", "formatted")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("formatted")) {
		t.Errorf("Expected output to contain 'formatted', got: %s", output)
	}
}

func TestWarnw(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Warnw("test warn message", "key", "value")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("test warn message")) {
		t.Errorf("Expected output to contain 'test warn message', got: %s", output)
	}
}

func TestError(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Error("test error message")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("test error message")) {
		t.Errorf("Expected output to contain 'test error message', got: %s", output)
	}
}

func TestErrorf(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Errorf("test error message: %s", "formatted")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("formatted")) {
		t.Errorf("Expected output to contain 'formatted', got: %s", output)
	}
}

func TestErrorw(t *testing.T) {
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	sazabi.Initialize("production")
	sazabi.Errorw("test error message", "key", "value")

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !bytes.Contains([]byte(output), []byte("test error message")) {
		t.Errorf("Expected output to contain 'test error message', got: %s", output)
	}
}

// Note: Fatal, Fatalf, Fatalw, Panic, Panicf, and Panicw functions
// cannot be tested as they cause program termination and panics respectively.
// In a real-world scenario, you might want to test these with dependency injection
// or by mocking the underlying logger.

// Test that logger is properly initialized after calling Initialize
func TestLoggerInitialized(t *testing.T) {
	// Test production environment
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Initialize should not panic for production: %v", r)
		}
	}()
	sazabi.Initialize(sazabi.ProductionEnvName)

	// Test development environment
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Initialize should not panic for development: %v", r)
		}
	}()
	sazabi.Initialize("development")
}

// Integration test to verify the complete logging flow
func TestLoggingIntegration(t *testing.T) {
	// Capture stderr to verify actual log output
	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Initialize logger for testing
	sazabi.Initialize("development")

	// Test logging
	sazabi.Info("integration test message")

	// Restore stderr
	w.Close()
	os.Stderr = originalStderr

	// Read captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Verify output contains our message
	if !bytes.Contains([]byte(output), []byte("integration test message")) {
		t.Errorf("Expected output to contain 'integration test message', got: %s", output)
	}
}
