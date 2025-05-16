package req

import (
	"gin_web_frame/model/ctype"
)

type ArticleCreateReq struct {
	Category    string         `json:"category" binding:"max=32" msg:"内容超过范围"  struct:"category"`                 // 分类
	Title       string         `json:"title" binding:"required,max=256" msg:"内容超过范围" struct:"title"`              // 标题
	Content     string         `json:"content,omitempty"  struct:"content"`                                       // 内容
	Cover       string         `json:"cover,omitempty" struct:"cover"`                                            // 封面
	Description string         `json:"description,omitempty" binding:"max=256" msg:"内容超过范围" struct:"description"` // 描述
	Status      int8           `json:"status,omitempty"  struct:"status"`                                         // 状态
	Tags        ctype.StrArray `json:"tags,omitempty"`                                                            // 标签
}

type TagReq struct {
	TagName string `json:"tag_name" binding:"min=1,max=32" msg:"格式错误"` // 标签名称
	TagDesc string `json:"tag_desc,omitempty" binding:"max=128"`       // 标签描述
}
