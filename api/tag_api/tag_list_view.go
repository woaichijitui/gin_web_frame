package tag_api

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/model/res"
	"gin_web_frame/service/service_com"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TagListView 分类管理列表页
// @Tags 分类管理
// @summary 分类管理列表页
// @Description 分类管理列表页
// @Param page query models.PageInfo true "表示单个参数"
// @Router /tag/list [get]
// @Produce json
// @success 200 {object} res.Response
func (TagApi) TagListView(c *gin.Context) {

	//	绑定参数
	var page models.PageInfo
	//绑定参数
	err := c.ShouldBindQuery(&page)
	if err != nil {
		zap.L().Error("查询参数错误", zap.Error(err))
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 获取总数
	var total int64
	if err = global.DB.Model(&models.Tag{}).Count(&total).Error; err != nil {
		zap.L().Error("查询所有tag失败", zap.Error(err))
		res.FailWithCode(res.DBError, c)
		return
	}

	var tags []models.Tag
	//	tag列表
	// 查询分页数据并预加载标签
	err = global.DB.
		Preload("Articles").
		Scopes(service_com.Paginate(page)).
		Find(&tags).Error
	if err != nil {
		zap.L().Error("分页查询tag失败", zap.Error(err))
		res.FailWithCode(res.DBError, c)
		return
	}
	// 内存中计算标签数量
	for i := range tags {
		tags[i].ArticleCount = len(tags[i].Articles)
	}

	//	响应
	res.OkWithList(tags, total, c)
}
