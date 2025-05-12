### 	笔记模型设计

```
Article struct {
Id int64 `db:"id"` // 主键ID
Title string `db:"title"` // 标题
Content string `db:"content"` // 内容
Cover string `db:"cover"` // 封面
Description string `db:"description"` // 描述
AuthorId int64 `db:"author_id"` // 作者ID
Status int64 `db:"status"` // 状态 0:待审核 1:审核不通过 2:可见 3:用户删除
CommentNum int64 `db:"comment_num"` // 评论数
LikeNum int64 `db:"like_num"` // 点赞数
CollectNum int64 `db:"collect_num"` // 收藏数
ViewNum int64 `db:"view_num"` // 浏览数
ShareNum int64 `db:"share_num"` // 分享数
TagIds string `db:"tag_ids"` // 标签ID
PublishTime time.Time `db:"publish_time"` // 发布时间
CreateTime time.Time `db:"create_time"` // 创建时间
UpdateTime time.Time `db:"update_time"` // 最后修改时间
}

```

