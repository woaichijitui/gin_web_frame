package article_api

import (
	"errors"
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/model/req"
	"gin_web_frame/model/res"
	"gin_web_frame/service"
	"gin_web_frame/utils/common"
	"go.uber.org/zap"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// ArticleUpdateView 文章更新
// @Tags 文章管理
// @summary 文章更新
// @Description 文章更新
// @Param id path string  true "URL 参数 ：id"
// @Param cr body req.ArticleCreateReq  false "更新文章"
// @Router /article/update/{id} [put]
// @Produce json
// @success 200 {object} res.Response{}
func (ArticleApi) ArticleUpdateView(c *gin.Context) {

	id := c.Param("id")

	var cr req.ArticleCreateReq
	//		绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//绑定错误 返回msg tag 中错误信息
		res.FailWithError(err, &cr, c)
		return
	}

	var article models.Article
	//根据id查询是否有此广告
	row := global.DB.Take(&article, "id = ?", id).RowsAffected
	if row == 0 {
		res.OkWithMassage("没有此广告", c)
		return
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

	//更新
	mp := structs.Map(&cr)
	err = service.Service.ArticleService.ArticleUpdateWithTags(&article, &mp, tags)
	if err != nil {
		err = errors.New("文章更新失败")
		zap.L().Error(err.Error())
		res.FailWithMassage(err.Error(), c)
		return
	}
	//	成功响应
	res.OkWithMassage("文章更新成功", c)
}
