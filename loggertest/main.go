package main

import (
	"github.com/lingjiao0710/test/logger"
)

var mlog logger.Logger

func main() {

	//logger.GetInfo(1)
	mlog = logger.NewConsoleLogger("debug")
	//mlog = logger.NewFileLogger("debug", "./", "log.txt", 10*1024*1024)
	for {
		mlog.Debug("%s %v", "一条日志", 10)
		mlog.Trace("一条日志")
		mlog.Info("一条日志")
		mlog.Warning("一条日志")
		mlog.Error("一条日志")
		mlog.Fatal("trace log ....")

		//time.Sleep(time.Second)
	}

}
