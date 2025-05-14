package models

// PageInfo 接收前端的分页需求
type PageInfo struct {
	Page  int    `form:"page"`  //那一页
	Key   string `form:"key"`   //模糊查询的关键字
	Limit int    `form:"limit"` //每页显示数量
	Sort  string `form:"sort"`  //时间排序 desc
}

// RemoveRequest 接收要删除的id
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
