package model

// PageQuery 表示分页查询的参数
type PageQuery struct {
	// 数据起始行
	Offset int64 `form:"offset"`
	// 数据总量
	Count int64 `form:"count"`
}

// PageResult 表示分页查询的结果
type PageResult struct {
	// 数据总数
	Total int64 `json:"total"`
	// 查询到的数据列表
	Data interface{} `json:"data"`
}
