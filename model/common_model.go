package models

// PageInfo 接收前端的分页需求
type PageInfo struct {
	Page  int    `form:"page" binding:"required,min=1"  msg:"参数错误"` //那一页
	Key   string `form:"key"  `                                     //模糊查询的关键字
	Limit int    `form:"limit" binding:"required,min=1" msg:"参数错误"` //每页显示数量
	Sort  string `form:"sort"`                                      //时间排序 desc
}

// RemoveRequest 接收要删除的id
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
