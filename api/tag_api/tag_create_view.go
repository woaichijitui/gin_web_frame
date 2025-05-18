package tag_api

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/model/req"
	"gin_web_frame/model/res"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// TagCreateView 创建分类
// @Tags 分类管理
// @summary 创建分类
// @Description 创建分类
// @Param cr body req.TagReq true "创建标签body"
// @Router /tag/create [post]
// @Produce json
// @success 200 {object} res.Response
func (a TagApi) TagCreateView(ctx *gin.Context) {
	//	接收参数
	var cr req.TagReq
	err := ctx.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	/*var tagModel *models.Tag

	// 检查标签是否已存在
	if row := global.DB.Where("tag_name = ?", cr.TagName).First(tagModel).RowsAffected; row != 0 {
		res.FailWithMassage("标签已存在", ctx)
		return
	}*/
	//创建标签模型
	tagModel := &models.Tag{
		TagName: cr.TagName,
		TagDesc: cr.TagDesc,
	}
	err = global.DB.Create(&tagModel).Error
	if err != nil {
		// 判断是否为唯一约束冲突（错误码 1062）
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			res.FailWithMassage("标签重复", ctx)
			return
		}
		zap.L().Error("分类插入错误: ", zap.Error(err))
		res.FailWithCode(res.DBError, ctx)
		return
	}

	res.OkWithData("创建成功", ctx)

}
