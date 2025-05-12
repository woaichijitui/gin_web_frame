package token

import "github.com/gin-gonic/gin"

func GetClaimsForCtx(ctx *gin.Context) *MyCustomClaims {
	claims, exists := ctx.Get("claims")
	if !exists || claims == nil {
		return nil
	}
	return claims.(*MyCustomClaims)

}

func GetClaimsRole(ctx *gin.Context) (role int, err error) {
	claims := GetClaimsForCtx(ctx)
	if claims == nil {
		return 0, err
	}
	return claims.Role, nil
}
