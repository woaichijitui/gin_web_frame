package user_api

import (
	"errors"
	"fmt"
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/model/ctype"
	"gin_web_frame/model/res"
	"gin_web_frame/utils"
	"gin_web_frame/utils/email"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
)

type RegisterRequest struct {
	Username   string  `json:"username,omitempty"  binding:"required" msg:"请正确输入用户名"`
	Nickname   string  `json:"nickname,omitempty" binding:"required" msg:"请正确输入昵称"`
	Password   string  `json:"password,omitempty" binding:"required" msg:"请正确输入密码"`
	RePassword string  `json:"re_password,omitempty" binding:"required" msg:"请正确输入密码"`
	Email      string  `json:"email,omitempty" binding:"required,email" msg:"请正确输入邮箱"`
	Code       *string `json:"code"`
}

var codeCache = cache.New(5*time.Minute, 10*time.Minute)

// UserRegisterView 邮箱或者用户名注册
// @Tags 用户管理
// @summary 邮箱或者用户名注册
// @Description 邮箱或者用户名注册
// @Param cr body RegisterRequest true "注册信息 "
// @Router /user_register [post]
// @Produce json
// @success 200 {object} res.Response
func (l UserApi) UserRegisterView(c *gin.Context) {
	//	接收参数
	var cr RegisterRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//判断用户名是否存在
	var userModel models.UserModel
	row := global.DB.Find(&userModel, "username = ?", cr.Username).RowsAffected
	if row > 0 {
		//	用户存在
		err := errors.New("用户已经存在")
		global.LOG.Error(err.Error())
		res.FailWithMassage(err.Error(), c)
		return
	}
	//校验两次密码
	if cr.Password != cr.RePassword {
		//	密码不一致
		err := errors.New("两次密码不一致")
		global.LOG.Error(err.Error())
		res.FailWithMassage(err.Error(), c)
		return
	}
	//hash加密密码
	hashPwd, err := utils.PasswordHash(cr.Password)
	if err != nil {
		err := errors.New(fmt.Sprintf("加密密码失败：%v\n", err))
		global.LOG.Error(err.Error())
		res.FailWithMassage(err.Error(), c)
		return
	}

	//验证邮箱

	//	第一次发送邮箱验证码
	if cr.Code == nil {

		//随机验证码
		code := utils.Code()
		//	发送验证码
		err = email.NewCode().SendEmail(cr.Email, fmt.Sprintf("你的验证码 %s", code))
		if err != nil {
			global.LOG.Error(err.Error())
			res.FailWithMassage("发送邮箱失败", c)
			return
		}
		//	将code存入本地内存
		codeCache.Set("mail_code:"+cr.Email, code, cache.DefaultExpiration)
		if err != nil {
			global.LOG.Error(err.Error())
			res.FailWithMassage("session设置失败", c)
			return
		}
		res.OkWithMassage("验证码邮件已发送", c)
		return
	}
	//	第二次 验证邮箱验证码

	redisCode, found := codeCache.Get("mail_code:" + cr.Email)
	if !found {
		res.FailWithMassage("验证码已过期或不存在", c)
		return
	}
	if *cr.Code != redisCode.(string) {
		res.OkWithMassage("验证码错误", c)
		return
	}

	//第一次邮箱 和 第二次邮箱一致性也要验证
	//这个放在前端验证了，这里就不验证了

	//头像
	//默认头像（地址
	avatar := "uploads/images/default.jpg"

	//	入库
	err = global.DB.Create(&models.UserModel{
		Nickname:   cr.Nickname,
		Username:   cr.Username,
		Password:   hashPwd,
		Avatar:     avatar,
		IP:         c.ClientIP(),
		Addr:       c.Request.URL.Path,
		Role:       ctype.PermissionUser,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		err := errors.New(fmt.Sprintf("创建用户失败：%v\n", err))
		global.LOG.Error(err.Error())
		res.FailWithMassage(err.Error(), c)
		return
	}
	global.LOG.Info(fmt.Sprintf("用户%s创建成功/n", cr.Username))
	//	响应
	res.OkWithData("注册成功", c)
}
