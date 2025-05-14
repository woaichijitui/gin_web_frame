package article_api

import (
	"fmt"
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/model/res"
	"gin_web_frame/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ArticleDeleteView 文章删除
// @Tags 文章管理
// @summary 文章删除
// @Description 文章删除
// @Param cr body models.RemoveRequest true "要删除的文章id列表"
// @Router /article/delete [delete]
// @Produce json
// @success 200 {object} res.Response{}
func (ArticleApi) ArticleDeleteView(c *gin.Context) {

	var cr models.RemoveRequest
	//	绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//	判断文件是否存在
	var articleList []models.Article
	count := global.DB.Find(&articleList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMassage("文章不存在", c)
		return
	}

	// 提取文章 ID 列表
	articleIDs := make([]uint, 0, len(articleList))
	for _, article := range articleList {
		articleIDs = append(articleIDs, article.ID)
	}
	//删除文章并删除关联记录
	err = service.Service.ArticleService.DeleteArticlesWithTag(articleIDs)
	if err != nil {
		zap.L().Error("删除文章失败", zap.Error(err))
		res.FailWithMassage(err.Error(), c)
		return
	}
	//	成功响应
	res.OkWithMassage(fmt.Sprintf("成功删除 %d 篇文章", count), c)
}
