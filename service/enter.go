package service

import (
	"gin_web_frame/service/article_tag_ser"
	"gin_web_frame/service/casbin_ser"
	"gin_web_frame/service/redis_ser"
	"gin_web_frame/service/user_ser"
)

type _Service struct {
	RedisService   redis_ser.RedisService
	UserService    user_ser.UserService
	CasbinService  casbin_ser.CasbinService
	ArticleService article_tag_ser.ArticleTagService
}

var Service = new(_Service)
