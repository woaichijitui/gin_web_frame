package routers

import (
	"gin_web_frame/api"
	"gin_web_frame/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.RouterGroup) {

	ArticleApi := api.ApiGroupApp.ArticleApi

	router.POST("/article/create", middleware.JwtAuth(), ArticleApi.ArticleCreateView)
	router.GET("/article/list", middleware.JwtAuth(), ArticleApi.ArticleListView)
	router.PUT("/article/update/:id", middleware.JwtAuth(), ArticleApi.ArticleUpdateView)

}
