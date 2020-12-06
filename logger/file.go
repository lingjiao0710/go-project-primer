package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel
	filePath    string //日志保存路径
	fileName    string //日志文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelString, filePath, fileName string, maxFileSize int64) *FileLogger {
	f1 := &FileLogger{
		Level:       parseLogLevel(levelString),
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}

	err := f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

//initFile 指定路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file %s failed, err :%v", fullFileName, err)
		return err
	}

	errfileObj, err := os.OpenFile(fullFileName+"err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file %s failed, err :%v", fullFileName, err)
		return err
	}

	f.fileObj = fileObj
	f.errFileObj = errfileObj

	return nil
}

//enable 日志开关
func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f *FileLogger) log(lv LogLevel, format string, args ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, args...)
		timeFormat := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, line := getInfo(3)

		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", timeFormat, getLogString(lv), fileName, funcName, line, msg)

		if lv >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", timeFormat, getLogString(lv), fileName, funcName, line, msg)
		}
	}

}

//Debug 调试级别
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DEBUG, format, args...)
}

//Trace 跟踪级别
func (f *FileLogger) Trace(format string, args ...interface{}) {
	f.log(DEBUG, format, args...)
}

//Info 信息级别
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(INFO, format, args...)
}

//Warning 警告级别
func (f *FileLogger) Warning(format string, args ...interface{}) {
	f.log(WARNING, format, args...)
}

//Error 错误级别
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ERROR, format, args...)
}

//Fatal 致命错误
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FATAL, format, args...)
}
