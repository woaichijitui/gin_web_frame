package middleware

import (
	"gin_web_frame/global"
	"gin_web_frame/model/res"
	"gin_web_frame/service"
	"gin_web_frame/utils/token"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

var casbinService = service.Service.CasbinService

func casbin_rbac_mid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取当前用户的角色
		role, err := token.GetClaimsRole(ctx)
		if err != nil {
			zap.L().Error("token.GetClaimsRole error: %v", zap.Error(err))
			ctx.Abort()
			return
		}
		roleStr := strconv.Itoa(role)
		//获取请求的PATHi
		path := ctx.Request.URL.Path
		obj := strings.TrimPrefix(path, global.CONFIG.System.RouterPrefix)
		// 获取请求方法
		act := ctx.Request.Method

		casbin := casbinService.CasbinInit()
		//	验证该角色有无权限
		ok, err := casbin.CanAccess(roleStr, obj, act)
		if err != nil {
			zap.L().Error("casbin  error: %v", zap.Error(err))
			ctx.Abort()
			return
		}
		if !ok {
			res.FailWithMassage("权限不足", ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
