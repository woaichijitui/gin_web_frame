package token

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetClaimsForCtx(ctx *gin.Context) (*MyCustomClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists || claims == nil {
		return nil, errors.New("claims not found in context")
	}
	customClaims, ok := claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}
	return customClaims, nil

}

func GetClaimsRole(ctx *gin.Context) (role int, err error) {
	claims, err := GetClaimsForCtx(ctx)
	if err != nil {
		return 0, err
	}
	return claims.Role, nil
}

func GetClaimsId(ctx *gin.Context) (int, error) {
	claims, err := GetClaimsForCtx(ctx)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil

}
