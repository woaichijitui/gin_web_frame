package api

import (
	"gin_web_frame/api/user_api"
)

type ApiGroup struct {
	//SettingsApi setting_api.SettingsApi
	//ImagesApi   images_api.ImagesApi
	//AdvertApi   advert_api.AdvertApi
	//MenuAPi     menu_api.MenuAPi
	LoginApi user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
