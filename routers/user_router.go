package routers

import (
	"gin_web_frame/api"
	"gin_web_frame/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {

	userApi := api.ApiGroupApp.UserApi

	router.POST("/login", userApi.EmailLoginView)
	router.POST("/user_register", userApi.UserRegisterView)
	router.GET("/user_logout", middleware.JwtAuth(), userApi.UserLogoutView)
	//router.GET("/users", middleware.JwtAuth(), userApi.UserListView)
	//router.PUT("/user_update_role", middleware.JwtAuth(), userApi.UserUpdateRoleView)
	//router.PUT("/user_update_pwd", middleware.JwtAuth(), userApi.UserUpdatePwdView)
	//router.DELETE("/user_delete", middleware.JwtAuth(), userApi.UserRemoveView)
	//router.POST("/user_bind_email", middleware.JwtAuth(), userApi.UserBindMailView)

}
