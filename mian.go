package main

import (
	"gin_web_frame/core"
	"gin_web_frame/global"
)

// @title 		    			zztag Swagger API接口文档
// @version 					1.0.1
// @description					这是一个自动生成关zztag项目api文档
// @securityDefinitions.apikey	ApiKeyAuth
// @in 							header
// @name 						x-token\
// @BasePath                    /
func main() {
	global.VP = core.Viper()    // 初始化Viper
	global.LOG = core.ZapInit() // 初始化zap日志库
	global.DB = core.Gorm()     // gorm连接数据库
	if global.DB != nil {
		core.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()

}
