package req

import "gin_web_frame/model/ctype"

type ArticleCreateReq struct {
	Category    string         `json:"category" binding:"max=32" msg:"内容超过范围"`               // 分类
	Title       string         `json:"title" binding:"required,max=256" msg:"内容超过范围"`        // 标题
	Content     string         `json:"content,omitempty" `                                         // 内容
	Cover       string         `json:"cover,omitempty"`                                            // 封面
	Description string         `json:"description,omitempty" binding:"max=256" msg:"内容超过范围"` // 描述
	Tags        ctype.StrArray `json:"tags,omitempty" `                                            // 标签
}
