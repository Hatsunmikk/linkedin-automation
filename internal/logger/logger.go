package logger

import (
	"log"
	"os"
)

//Logger wraps the standard Go logger to provide
//levelled logging (INFO, WARN, ERROR, DEBUG).
//This abstraction allows us to control log output
//centrally across the application.

type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// New creates and configures a new Logger instance
// Log output is written to stdout or stderr depending on severity
func New(debug bool) *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime),
	}
}

// Info logs general application events
func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
}

// Warn logs unexpected but recoverable situations
func (l *Logger) Warn(msg string) {
	l.warnLogger.Println(msg)
}

// Error logs critical failures that require attention
func (l *Logger) Error(msg string) {
	l.errorLogger.Println(msg)
}

// Debug logs verbose output useful during development
// and troubleshooting automation behaviour
func (l *Logger) Debug(msg string) {
	l.debugLogger.Println(msg)
}
