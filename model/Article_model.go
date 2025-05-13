package models

import (
	"gin_web_frame/model/ctype"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string         `gorm:"size:256"`  // 标题
	Content     string         `gorm:"type:text"` // 内容
	Cover       string         // 封面
	Description string         `gorm:"size:256"` // 描述
	AuthorId    uint           // 作者ID
	Status      int8           // 状态 0:待审核 1:审核不通过 2:发布 3:用户删除
	CommentNum  int            // 评论数
	LikeNum     int            // 点赞数
	CollectNum  int            // 收藏数
	ViewNum     int            // 浏览数
	ShareNum    int            // 分享数
	Tags        ctype.StrArray // 标签ID

}
