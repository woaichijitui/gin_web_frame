package models

import (
	"gin_web_frame/model/ctype"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model `json:"gorm.Model"`
	Nickname   string           `gorm:"size:36"   json:"nickname,omitempty"`      //昵称
	Username   string           `gorm:"size36"   json:"username,omitempty"`       //用户名
	Password   string           `gorm:"size:128" json:"password,omitempty"`       //密码
	Avatar     string           `gorm:"size:246"   json:"avatar,omitempty"`       //头像id
	Email      string           `gorm:"size:128"   json:"email,omitempty"`        //邮箱
	Tel        string           `gorm:"size:18"   json:"tel,omitempty"`           //手机
	Addr       string           `gorm:"size:64"   json:"addr,omitempty"`          //地址
	Token      string           `gorm:"size:64"   json:"token,omitempty"`         //其他平台唯一id
	IP         string           `gorm:"size:20"   json:"IP,omitempty"`            //ip地址
	Tags       ctype.StrArray   `gorm:"type:json"   json:"tag,omitempty"`         //标签
	Role       ctype.Role       `gorm:"size:4;default:1"   json:"role,omitempty"` //用户权限
	SignStatus ctype.SignStatus `gorm:"type:smallint(6)" json:"sign_status"`      //用户登录方式
}
