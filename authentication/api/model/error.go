package model

// Error 表示请求错误的结果
type Error struct {
	// 错误代码
	Code int64 `json:"code"`
	// 错误的描述
	Error string `json:"error"`
}
