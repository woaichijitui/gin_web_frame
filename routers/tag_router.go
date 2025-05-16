package routers

import (
	"gin_web_frame/api"
	"gin_web_frame/middleware"
	"github.com/gin-gonic/gin"
)

func TagRouter(router *gin.RouterGroup) {

	tagApi := api.ApiGroupApp.TagApi

	router.GET("/tag/list", middleware.JwtAuth(), tagApi.TagListView)
	router.POST("/tag/create", middleware.JwtAuth(), tagApi.TagCreateView)
	router.PUT("/tag/update/:id", middleware.JwtAuth(), tagApi.TagUpdateView)
	router.DELETE("/tag/delete/:id", middleware.JwtAuth(), tagApi.TagDeleteView)
	router.GET("/tag/:id/articles", middleware.JwtAuth(), tagApi.ArticleListInTagView)
	//router.POST("/login", userApi.EmailLoginView)
	//router.POST("/user_register", userApi.UserRegisterView)
	//router.GET("/users", middleware.JwtAuth(), userApi.UserListView)
	//router.POST("/user_bind_email", middleware.JwtAuth(), userApi.UserBindMailView)

}
