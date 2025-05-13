package initialize

import (
	"gin_web_frame/global"
	"gin_web_frame/middleware"
	"gin_web_frame/routers"
	"github.com/swaggo/swag/example/basic/docs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type justFilesFile struct {
	fs http.FileSystem
}

func (fs justFilesFile) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if global.CONFIG.System.Mode == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	Router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group(global.CONFIG.System.RouterPrefix)

	// 链式中间件
	//ApiGroup.Use(middleware.JwtAuth())
	ApiGroup.Use(middleware.GinZapMiddleware(global.LOG))

	{
		// 健康监测
		ApiGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	// 注册用户路由
	routers.UserRouter(ApiGroup)
	routers.ArticleRouter(ApiGroup)

	global.ROUTERS = Router.Routes()

	global.LOG.Info("router register success")
	return Router
}
