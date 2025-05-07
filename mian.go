package main

import (
	"gin_web_gin/core"
	"gin_web_gin/global"
	"time"
)

func main() {
	global.VP = core.Viper()
	global.LOG = core.ZapInit()
	global.LOG.Debug("test1")
	global.LOG.Info("test1")
	global.LOG.Warn("test1")
	global.LOG.Error("test1")

	time.Sleep(time.Hour)
}
