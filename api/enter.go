package api

import (
	"gin_web_frame/api/article_api"
	"gin_web_frame/api/tag_api"
	"gin_web_frame/api/user_api"
)

type ApiGroup struct {
	//SettingsApi setting_api.SettingsApi
	//ImagesApi   images_api.ImagesApi
	//MenuAPi     menu_api.MenuAPi
	ArticleApi article_api.ArticleApi
	UserApi    user_api.UserApi
	TagApi     tag_api.TagApi
}

var ApiGroupApp = new(ApiGroup)
