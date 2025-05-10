package service

import (
	"gin_web_frame/service/redis_ser"
	"gin_web_frame/service/user_ser"
)

type _Service struct {
	RedisService redis_ser.RedisService
	UserService  user_ser.UserService
}

var Service = new(_Service)
