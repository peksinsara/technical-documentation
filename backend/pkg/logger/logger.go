package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	logger *log.Logger
)

// Init initializes the logger
func Init(logDir string) error {
	// Get absolute path
	absPath, err := filepath.Abs(logDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	fmt.Printf("Initializing logger in directory: %s\n", absPath)

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(absPath, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	logPath := filepath.Join(absPath, "technical-docu.log")
	fmt.Printf("Creating log file at: %s\n", logPath)

	// Create log file
	logFile, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	// Initialize logger
	logger = log.New(logFile, "", log.Ldate|log.Ltime)

	// Write initial log message
	logger.Printf("[INFO] Logger initialized")
	fmt.Println("Logger initialized successfully")

	return nil
}

// Info logs an info message
func Info(format string, v ...interface{}) {
	logger.Printf("[INFO] "+format, v...)
}

// Error logs an error message
func Error(format string, v ...interface{}) {
	logger.Printf("[ERROR] "+format, v...)
}

// Debug logs a debug message
func Debug(format string, v ...interface{}) {
	logger.Printf("[DEBUG] "+format, v...)
}

// Request logs an HTTP request
func Request(method, path, status string, duration time.Duration, userID uint) {
	logger.Printf("[REQUEST] %s %s | Status: %s | Duration: %v | UserID: %d",
		method, path, status, duration, userID)
}
