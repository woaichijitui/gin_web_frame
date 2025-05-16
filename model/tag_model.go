package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName      string    `gorm:"size:32;not null;unique"` // 标签名称
	TagDesc      string    `gorm:"size:128;not null"`       // 标签描述
	ArticleCount int       // 关联文章数量 目前不是实时更新的
	Articles     []Article `gorm:"many2many:article_tag_relations" json:"-"` // 关联文章
}
