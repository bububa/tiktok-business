package model

// PageInfo 分页信息
type PageInfo struct {
	// Page 当前页数
	Page int `json:"page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalNumber 总结果数
	TotalNumber int `json:"total_number,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
}
