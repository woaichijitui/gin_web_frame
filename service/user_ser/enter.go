package user_ser

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gin_web_frame/utils"
	"gin_web_frame/utils/token"

	"github.com/gin-gonic/gin"
	"time"
)

type UserService struct {
}

// GetTokenExp 获取token的过期时间
func (u UserService) GetTokenExp(c *gin.Context) time.Duration {
	//获取token
	_claims, _ := c.Get("claims")
	claims := _claims.(*token.MyCustomClaims)

	//获取token过期时间
	expirationTime, _ := claims.GetExpirationTime()

	exp := expirationTime.Sub(time.Now())

	return exp
}

// CheckPwd 确认密码是否正确
func (u UserService) CheckPwd(userID uint, pwd string) bool {
	var userModel models.UserModel
	err := global.DB.Find(&userModel, "id = ? ", userID).Error
	if err != nil {
		return false
	}
	if ok := utils.PasswordVerify(pwd, userModel.Password); !ok {
		return false
	}

	return true

}
