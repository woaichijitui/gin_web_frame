package middleware

import (
	"gin_web_frame/global"
	"gin_web_frame/model/ctype"
	"gin_web_frame/model/res"
	"gin_web_frame/service"

	utils "gin_web_frame/utils/token"

	"github.com/gin-gonic/gin"
)

// JwtAuth 普通用户授权中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否是管理者
		token := c.GetHeader("Authorization")
		//有无token
		if token == "" {
			res.FailWithMassage("未登录或非法访问", c)
			c.Abort()
			return
		}
		//注销的用户
		logout := service.Service.RedisService.CheckLogout(token)

		if logout {
			res.FailWithMassage("用户已注销", c)
			c.Abort()
			return
		}

		//解析token
		claims, err := utils.ParseTokenRs256(token)
		if err != nil {
			global.LOG.Error(err.Error())
			res.FailWithMassage("token解析失败", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)

	}
}

// JwtAdmin 管理用户授权中间件
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否是管理者
		token := c.GetHeader("Authorization")
		//有无token
		if token == "" {
			res.FailWithMassage("未携带token", c)
			c.Abort()
			return
		}
		//解析token
		claims, err := utils.ParseTokenRs256(token)
		if err != nil {
			global.LOG.Error(err.Error())
			res.FailWithMassage("token解析失败", c)
			c.Abort()
			return
		}

		//注销的用户
		logout := service.Service.RedisService.CheckLogout(token)
		if logout {
			res.FailWithMassage("用户已注销", c)
			c.Abort()
			return
		}

		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//	若不是管理员
			res.FailWithMassage("非管理员用户", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}

}
