package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"gin_web_frame/config"
	"gin_web_frame/global"
	"gin_web_frame/model/ctype"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// MyCustomClaims
type MyCustomClaims struct {
	UserID     int    `json:"user_id"`
	Username   string `json:"name"`
	GrantScope string `json:"grant_scope"`
	Role       int    `json:"role"`
	jwt.RegisteredClaims
}

// pkcs1 PKCS #1（Public-Key Cryptography Standards #1）是由RSA实验室定义的一组公钥加密标准，用于描述RSA加密和签名算法的实现细节。PKCS #1标准包括了关于密钥格式、填充方案以及数字签名和加密流程的规范。Go语言中，常用的库如crypto/rsa和crypto/x509提供了对PKCS #1标准的支持。
func parsePriKeyBytes(buf []byte) (*rsa.PrivateKey, error) {
	// 解码 PEM 块
	p, _ := pem.Decode(buf)
	if p == nil {
		return nil, errors.New("parse key error: failed to decode PEM block")
	}

	// 尝试解析 PKCS#1 私钥
	if key, err := x509.ParsePKCS1PrivateKey(p.Bytes); err == nil {
		return key, nil
	}

	// 尝试解析 PKCS#8 私钥
	if key, err := x509.ParsePKCS8PrivateKey(p.Bytes); err == nil {
		switch k := key.(type) {
		case *rsa.PrivateKey:
			return k, nil
		default:
			return nil, errors.New("parse key error: not an RSA private key")
		}
	}

	return nil, errors.New("parse key error: failed to parse private key")
}

// GenerateTokenUsingRS256 生成token
func GenerateTokenUsingRS256(userID uint, username string, role ctype.Role) (string, error) {
	claims := MyCustomClaims{
		UserID:     int(userID),
		Username:   username,
		Role:       int(role),
		GrantScope: global.CONFIG.JWT.GrantScope,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.CONFIG.JWT.Issuer,                                                                 //iss(Issuer)	发行者，标识 JWT 的发行者。
			Subject:   global.CONFIG.JWT.Subject,                                                                //sub(Subject)	主题，标识 JWT 的主题，通常指用户的唯一标识
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP", "WEB_APP"},                                    //aud(Audience)	观众，标识 JWT 的接收者
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(global.CONFIG.JWT.Expires))), //exp(Expiration Time)	过期时间。标识 JWT 的过期时间，这个时间必须是将来的
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Millisecond)),                                     //nbf(Not Before)	不可用时间。在此时间之前，JWT 不应被接受处理
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                           //iat(Issued At)	发行时间，标识 JWT 的发行时间
			ID:        GenerateSalt(10),                                                                         //jti(JWT ID)	JWT 的唯一标识符，用于防止 JWT 被重放（即重复使用）
		},
	}

	//
	rsaPriKey, err := parsePriKeyBytes([]byte(config.PRI_KEY))
	if err != nil {
		return "", err
	}

	//	生成token
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPriKey)
	if err != nil {
		return "", err
	}

	return token, nil

}

// parsePubKeyBytes 解析 PEM 编码的 RSA 公钥
func ParsePubKeyBytes(pub_key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub_key)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the key")
	}

	var pub interface{}
	var err error

	// 尝试解析 PKIX 格式的公钥
	pub, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err == nil {
		if rsaPub, ok := pub.(*rsa.PublicKey); ok {
			return rsaPub, nil
		}
		return nil, errors.New("key is not of type RSA")
	}

	// 尝试解析 PKCS1 格式的公钥
	pub, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		if rsaPub, ok := pub.(*rsa.PublicKey); ok {
			return rsaPub, nil
		}
		return nil, errors.New("key is not of type RSA")
	}

	return nil, errors.New("failed to parse RSA public key")
}

// ParseTokenRs256 解析token
func ParseTokenRs256(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		pub, err := ParsePubKeyBytes([]byte(config.PUB_KEY))
		if err != nil {
			//global.Log.Errorf("parsePubKeyBytes error: %v", err)
			return nil, err
		}
		return pub, nil
	})
	if err != nil {
		return nil, err
	}
	//	判断是否失效
	if !token.Valid {
		return nil, errors.New("claim invalid")
	}
	//	断言判断是否成功
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}
	return claims, nil
}
