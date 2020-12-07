package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
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

type Logger interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

//parseLogLevel 将字符串日志级别解析为常量
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

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "debug"
	case TRACE:
		return "trace"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		return "error"
	}

	return "error"
}

func getInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	//pc 4877540 file D:/code/go/src/github.com/lingjiao0710/test/main.go line 10
	//fmt.Printf("pc %v file %v line %v\n", pc, file, line)

	//main.main()
	funcName = runtime.FuncForPC(pc).Name()
	//main.go
	fileName = path.Base(file)

	//fmt.Printf("funcname :%v, fileName :%v\n", funcName, fileName)
	return
}
