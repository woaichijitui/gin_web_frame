package user_api

import (
	"gin_web_frame/global"
	"gin_web_frame/model/res"
	"gin_web_frame/service"
	"github.com/gin-gonic/gin"
)

// UserLogoutView 用户注销
// @Tags 用户管理
// @summary 用户注销
// @Description 用户注销
// @Param Authorization header string  true "token"
// @Router /user_logout [get]
// @Produce json
// @success 200 {object} res.Response
func (UserApi) UserLogoutView(c *gin.Context) {

	token := c.GetHeader("Authorization")

	exp := service.Service.UserService.GetTokenExp(c)

	//	 将token和过期存入redis
	err := service.Service.RedisService.SetLogoutToken(token, exp)
	if err != nil {
		global.LOG.Error(err.Error())
		res.FailWithMassage("注销失败", c)
		return
	}

	//	成功响应
	res.OkWithMassage("注销成功", c)

}
