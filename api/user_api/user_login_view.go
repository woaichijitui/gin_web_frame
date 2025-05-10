package user_api

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"

	"gin_web_frame/model/res"
	"gin_web_frame/utils"
	"gin_web_frame/utils/token"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required" msg:"用户名不存在"`
	Password string `json:"password" binding:"required" msg:"密码不正确"`
}

// EmailLoginView 用户管理
// @Tags 用户管理
// @summary 用户管理
// @Description 用户管理
// @Param cr body LoginRequest true "用户 密码 "
// @Router /login [post]
// @Produce json
// @success 200 {object} res.Response
func (l UserApi) EmailLoginView(c *gin.Context) {
	//	接收参数
	var cr LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//	判断是否有该用户
	var userModel models.UserModel
	row := global.DB.Find(&userModel, "username = ? or email = ?", cr.Username, cr.Username).RowsAffected
	if row == 0 {
		res.FailWithMassage("用户名或密码错误", c)
		return
	}

	//	密码验证
	ok := utils.PasswordVerify(cr.Password, userModel.Password)
	if !ok {
		res.FailWithMassage("密码输入错误", c)
		return
	}

	//	生成token
	token, err := token.GenerateTokenUsingRS256(userModel.ID, userModel.Username, userModel.Role)
	if err != nil {
		global.LOG.Error(err.Error())
		res.FailWithMassage("生成token错误", c)
		return
	}
	//	响应
	res.OkWithData(token, c)
}
