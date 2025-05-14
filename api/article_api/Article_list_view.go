package article_api

import (
	models "gin_web_frame/model"
	"gin_web_frame/model/res"
	"gin_web_frame/service/service_com"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ArticleListView 文章列表
// @Tags 文章管理
// @summary 文章列表
// @Description 文章列表
// @Param page query models.PageInfo true "表示单个参数"
// @Router /article/list [get]
// @Produce json
// @success 200 {object} res.Response
func (ArticleApi) ArticleListView(c *gin.Context) {

	//	绑定参数
	var page models.PageInfo
	//绑定参数
	err := c.ShouldBindQuery(&page)
	if err != nil {
		zap.L().Error("models.PageInfo with query err", zap.Error(err))
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var articleModel models.Article
	//	用户列表
	list, count, err := service_com.ComList(articleModel, service_com.Option{page, true})
	if err != nil {
		zap.L().Error("service_com.ComList err", zap.Error(err))
		res.FailWithMassage(err.Error(), c)
		return
	}

	//	响应
	res.OkWithList(list, count, c)
}
