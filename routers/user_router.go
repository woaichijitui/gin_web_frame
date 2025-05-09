package routers

import (
	"gin_web_frame/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {

	loginApi := api.ApiGroupApp.LoginApi

	router.POST("/email_login", loginApi.EmailLoginView)
	/*router.GET("/users", middleware.JwtAuth(), loginApi.UserListView)
	router.PUT("/user_update_role", middleware.JwtAuth(), loginApi.UserUpdateRoleView)
	router.PUT("/user_update_pwd", middleware.JwtAuth(), loginApi.UserUpdatePwdView)
	router.GET("/user_logout", middleware.JwtAuth(), loginApi.UserLogoutView)
	router.DELETE("/user_delete", middleware.JwtAuth(), loginApi.UserRemoveView)
	router.POST("/user_bind_email", middleware.JwtAuth(), loginApi.UserBindMailView)
	router.POST("/user_register", loginApi.UserRegisterView)*/

}
