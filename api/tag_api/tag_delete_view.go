package tag_api

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/model/res"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TagDeleteView 删除分类
// @Tags 分类管理
// @summary 删除分类
// @Description 删除分类
// @Param id path string  true "URL 参数: id"
// @Router /tag/delete/{id} [delete]
// @Produce json
// @success 200 {object} res.Response
func (a TagApi) TagDeleteView(ctx *gin.Context) {
	//	接收参数
	id := ctx.Param("id")

	var tagModel models.Tag
	//	判断分类是否存在
	row := global.DB.Find(&tagModel, id).RowsAffected
	if row == 0 {
		res.FailWithMassage("标签不存在", ctx)
		return
	}

	// 删除分类和其关联的文章
	err := global.DB.Select("Articles").Delete(&tagModel).Error
	if err != nil {
		zap.L().Error("删除tag失败", zap.Error(err))
		res.FailWithMassage("删除失败", ctx)
		return
	}
	res.OkWithData("删除成功", ctx)

}
