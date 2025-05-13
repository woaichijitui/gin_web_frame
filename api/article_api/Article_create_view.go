package article_api

import (
	models "gin_web_frame/model"
	"gin_web_frame/model/req"
	"gin_web_frame/model/res"
	"gin_web_frame/service"
	"gin_web_frame/utils/common"
	"gin_web_frame/utils/token"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ArticleCreateView  文章创建
// @Tags 文章管理
// @summary 文章创建
// @Description 文章创建
// @Param cr body req.ArticleCreateReq true "创建文章form"
// @Router /article/create [post]
// @Produce json
// @success 200 {object} res.Response
func (a ArticleApi) ArticleCreateView(ctx *gin.Context) {
	//	接收参数
	var cr req.ArticleCreateReq
	err := ctx.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}
	//	标题重复？

	// 获取当前用户id
	id, err := token.GetClaimsId(ctx)
	if err != nil {
		zap.L().Warn("获取token失败", zap.Error(err))
		res.FailWithMassage("获取token失败", ctx)
		return
	}

	// 默认封面
	if cr.Cover == "" {
		cr.Cover = "默认封面url"
	}
	// 默认分类
	if cr.Category == "" {
		cr.Category = "默认分类"
	}

	article := &models.Article{
		//	创建文章
		Category:    cr.Category,
		Title:       cr.Title,
		Content:     cr.Content,
		Cover:       cr.Cover,
		Description: cr.Description,
		AuthorId:    uint(id),
	}
	// 创建标签模型
	cr.Tags = common.ListUnique(cr.Tags) //切片去重
	var tags []models.Tag
	for _, tagName := range cr.Tags {
		tag := &models.Tag{
			TagName: tagName,
		}
		tags = append(tags, *tag)
	}

	//文章入库
	err = service.Service.ArticleService.ArticleCreateWithTags(article, tags)
	if err != nil {
		zap.L().Error("文章创建失败", zap.Error(err))
		res.FailWithMassage("文章创建失败", ctx)
		return
	}

	//	响应
	res.OkWithData("创建成功", ctx)

}
