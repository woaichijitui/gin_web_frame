package routers

import (
	"gin_web_frame/api"
	"gin_web_frame/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {

	loginApi := api.ApiGroupApp.LoginApi

	router.POST("/login", loginApi.EmailLoginView)
	router.POST("/user_register", loginApi.UserRegisterView)
	router.GET("/user_logout", middleware.JwtAuth(), loginApi.UserLogoutView)
	//router.GET("/users", middleware.JwtAuth(), loginApi.UserListView)
	//router.PUT("/user_update_role", middleware.JwtAuth(), loginApi.UserUpdateRoleView)
	//router.PUT("/user_update_pwd", middleware.JwtAuth(), loginApi.UserUpdatePwdView)
	//router.DELETE("/user_delete", middleware.JwtAuth(), loginApi.UserRemoveView)
	//router.POST("/user_bind_email", middleware.JwtAuth(), loginApi.UserBindMailView)

}
