package main

import (
	"gin_web_frame/core"
	"gin_web_frame/global"
	"time"
)

func main() {
	global.VP = core.Viper()    // 初始化Viper
	global.LOG = core.ZapInit() // 初始化zap日志库
	global.DB = core.Gorm()     // gorm连接数据库

	global.LOG.Debug("test1")
	global.LOG.Info("test1")
	global.LOG.Warn("test1")
	global.LOG.Error("test1")

	time.Sleep(time.Hour)
}
