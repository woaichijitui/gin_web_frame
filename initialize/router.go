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

	Router.Use(middleware.GinZapMiddleware(global.LOG))

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	//Router.StaticFS(global.CONFIG.Local.StorePath, justFilesFile{http.Dir(global.CONFIG.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	Router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group(global.CONFIG.System.RouterPrefix)

	//ApiGroup.Use(middleware.JwtAuth())

	{
		// 健康监测
		ApiGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	/*{
		Router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		Router.InitInitRouter(PublicGroup) // 自动初始化相关
	}

	{
		Router.InitApiRouter(PrivateGroup, PublicGroup)                // 注册功能api路由
		Router.InitJwtRouter(PrivateGroup)                             // jwt相关路由
		Router.InitUserRouter(PrivateGroup)                            // 注册用户路由
		Router.InitMenuRouter(PrivateGroup)                            // 注册menu路由
		Router.InitRouter(PrivateGroup)                                // 相关路由
		Router.InitCasbinRouter(PrivateGroup)                          // 权限相关路由
		Router.InitAutoCodeRouter(PrivateGroup, PublicGroup)           // 创建自动化代码
		Router.InitAuthorityRouter(PrivateGroup)                       // 注册角色路由
		Router.InitSysDictionaryRouter(PrivateGroup)                   // 字典管理
		Router.InitAutoCodeHistoryRouter(PrivateGroup)                 // 自动化代码历史
		Router.InitSysOperationRecordRouter(PrivateGroup)              // 操作记录
		Router.InitSysDictionaryDetailRouter(PrivateGroup)             // 字典详情管理
		Router.InitAuthorityBtnRouterRouter(PrivateGroup)              // 按钮权限管理
		Router.InitSysExportTemplateRouter(PrivateGroup, PublicGroup)  // 导出模板
		Router.InitSysParamsRouter(PrivateGroup, PublicGroup)          // 参数管理
		exampleRouter.InitCustomerRouter(PrivateGroup)                 // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)    // 文件上传下载功能路由
		exampleRouter.InitAttachmentCategoryRouterRouter(PrivateGroup) // 文件上传下载分类

	}*/

	// 注册用户路由
	routers.UserRouter(ApiGroup)
	routers.ArticleRouter(ApiGroup)

	global.ROUTERS = Router.Routes()

	global.LOG.Info("router register success")
	return Router
}
