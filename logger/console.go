package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	Level LogLevel
}

//NewLog 初始化日志结构
func NewLog(levelString string) Logger {
	logLevel := parseLogLevel(levelString)
	return Logger{
		Level: logLevel,
	}
}

//enable 日志开关
func (l Logger) enable(level LogLevel) bool {
	return level >= l.Level
}

func log(lv LogLevel, msg string) {
	timeFormat := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, line := getInfo(3)

	fmt.Fprintf(os.Stdout, "[%s] [%s] [%s:%s:%d] %s\n", timeFormat, getLogString(lv), fileName, funcName, line, msg)
}

//Debug 调试级别
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}

}

//Trace 跟踪级别
func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		log(TRACE, msg)
	}
}

//Info 信息级别
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

//Warning 警告级别
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}

//Error 错误级别
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}

//Fatal 致命错误
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
