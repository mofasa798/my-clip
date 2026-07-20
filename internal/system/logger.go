package system

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// LogLevel represents the severity of a log message.
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger provides structured logging for the application.
type Logger struct {
	debug  bool
	logger *log.Logger
	file   io.Closer
}

// NewLogger creates a new Logger instance.
func NewLogger() *Logger {
	logDir := getLogDir()
	os.MkdirAll(logDir, 0755)

	logPath := fmt.Sprintf("%s/myclip_%s.log", logDir, time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Warning: could not open log file %s: %v", logPath, err)
		return &Logger{
			logger: log.New(os.Stdout, "", log.LstdFlags),
			debug:  false,
		}
	}

	multi := io.MultiWriter(os.Stdout, f)
	return &Logger{
		logger: log.New(multi, "", log.LstdFlags),
		file:   f,
		debug:  false,
	}
}

// EnableDebug turns on debug-level logging.
func (l *Logger) EnableDebug() {
	l.debug = true
}

func getLogDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		return os.TempDir()
	}
	return dir + "/my-clip/logs"
}

// Debug logs a debug message.
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.debug {
		l.log("DEBUG", format, args...)
	}
}

// Info logs an informational message.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log("INFO", format, args...)
}

// Warn logs a warning message.
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log("WARN", format, args...)
}

// Error logs an error message.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log("ERROR", format, args...)
}

func (l *Logger) log(level, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.logger.Printf("[%s] %s", level, msg)
}
