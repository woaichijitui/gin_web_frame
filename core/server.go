package core

import (
	"fmt"
	"gin_web_frame/global"
	"gin_web_frame/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		if global.CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}
	//
	//if global.CONFIG..UseMongo {
	//	err := initialize.Mongo.Initialization()
	//	if err != nil {
	//		zap.L().Error(fmt.Sprintf("%+v", err))
	//	}
	//}

	//// 从db加载jwt数据 黑名单
	//if global.DB != nil {
	//	.LoadAll()
	//}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)

	global.LOG.Info("server run success on ", zap.String("address", address))

	global.LOG.Error(s.ListenAndServe().Error())
}
