package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName    string `gorm:"size:64,not null,unique"` // 标签名称
	TagDesc    string `gorm:"size:128,not null"`       // 标签描述
	ArticleIds []uint `gorm:"type:json"`               // 关联文章
	Total      int    `gorm:"not null"`                // 文章数量
}
