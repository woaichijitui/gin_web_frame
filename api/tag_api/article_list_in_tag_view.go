package tag_api

import (
	models "gin_web_frame/model"
	"gin_web_frame/model/res"
	"gin_web_frame/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ArticleListInTagView 分类的文章列表页
// @Tags 分类管理
// @summary 分类的文章列表页
// @Description 分类的文章列表页
// @Param id path string  true "URL 参数: id"
// @Param page query models.PageInfo false "页码"
// @Router /tag/{id}/articles [get]
// @Produce json
// @success 200 {object} res.Response
func (TagApi) ArticleListInTagView(c *gin.Context) {

	id := c.Param("id")
	tagId, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("查询参数错误", zap.Error(err))
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//	绑定参数
	var page models.PageInfo
	//绑定参数
	err = c.ShouldBindQuery(&page)
	if err != nil {
		zap.L().Error("查询参数错误", zap.Error(err))
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//	查询
	tag, articles, count, err := service.Service.ArticleService.GetArticlesByTagId(uint(tagId), page)
	
	if err != nil {
		//
		if tag.ID == -1 {
			res.FailWithMassage("没有该标签", c)
			return
		}
		zap.L().Error("查询错误", zap.Error(err))
		res.FailWithCode(res.DBError, c)
		return
	}

	//	响应
	res.OkWithData(gin.H{"tag": tag, "list": articles, "count": count}, c)
}
