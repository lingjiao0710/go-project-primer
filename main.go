package main

import (
	"github.com/lingjiao0710/test/logger"
	"time"
)

func main() {
	mlog := logger.NewLog("error")
	for {
		mlog.Debug("一条日志")
		mlog.Trace("一条日志")
		mlog.Info("一条日志")
		mlog.Warning("一条日志")
		mlog.Error("一条日志")
		mlog.Fatal("trace log ....")

		time.Sleep(time.Second)
	}

}
