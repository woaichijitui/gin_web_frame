package tag_api

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/model/req"
	"gin_web_frame/model/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// TagUpdateView 更新分类
// @Tags 分类管理
// @summary 更新分类
// @Description 更新分类
// @Param id path string  true "URL 参数 ：id"
// @Param cr body req.TagReq true "更新标签body"
// @Router /tag/update/{id} [put]
// @Produce json
// @success 200 {object} res.Response
func (a TagApi) TagUpdateView(ctx *gin.Context) {
	//	接收参数
	id := ctx.Param("id")
	var cr req.TagReq
	err := ctx.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	//
	var tagModel models.Tag
	mp := structs.Map(&cr)
	if err = global.DB.Model(&tagModel).Where("id = ?", id).Updates(mp).Error; err != nil {
		// 判断是否为唯一键冲突错误
		//MySQL错误判断（根据驱动类型选择）
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				res.FailWithMassage("标题重复，更新失败", ctx)
				return
			}
		}
		zap.L().Error(err.Error())
		res.FailWithMassage("更新失败", ctx)
		return

	}

	res.OkWithData("更新成功", ctx)

}
