package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type LogLevel int16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	Level LogLevel
}

func parseLogLevel(levelString string) LogLevel {
	lowerString := strings.ToLower(levelString)
	switch lowerString {
	case "debug":
		return DEBUG
	case "trace":
		return TRACE
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return INFO
	}
}

func NewLog(levelString string) Logger {
	logLevel := parseLogLevel(levelString)
	return Logger{
		Level: logLevel,
	}
}

func (l Logger) enable(level LogLevel) bool {
	return level >= l.Level
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(os.Stdout, "[%s] [DEBUG] %s\n", timeFormat, msg)
	}

}

func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(os.Stdout, "[%s] [TRACE] %s\n", timeFormat, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(os.Stdout, "[%s] [INFO] %s\n", timeFormat, msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(os.Stdout, "[%s] [WARNING] %s\n", timeFormat, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(os.Stdout, "[%s] [ERROR] %s\n", timeFormat, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(os.Stdout, "[%s] [FATAL] %s\n", timeFormat, msg)
	}
}
