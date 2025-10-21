package feed

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// LogRequest 获取更新源日志 API Request
type LogRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedID 更新源ID
	FeedID string `json:"feed_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *LogRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	values.Set("feed_id", r.FeedID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// LogResponse 获取更新源日志 API Response
type LogResponse struct {
	model.BaseResponse
	Data struct {
		// FeedLogs 日志列表
		FeedLogs []Log `json:"feed_logs,omitempty"`
	} `json:"data"`
}

// Log 日志
type Log struct {
	// UpdateStatus 更新信息
	UpdateStatus *UpdateStatus `json:"update_status,omitempty"`
	// UpdateTime 更新时间信息
	UpdateTime *UpdateTime `json:"update_time,omitempty"`
}

// UpdateStatus 更新信息
type UpdateStatus struct {
	// AddCount 增加的记录数
	AddCount int `json:"add_count,omitempty"`
	// ErrorCount 更新出错的记录数
	ErrorCount int `json:"error_count,omitempty"`
	// RemoveCount 删除的记录数
	RemoveCount int `json:"remove_count,omitempty"`
	// ProcessStatus 更新操作状态。 枚举值: PROCESSING(处理中), SUCCESS(成功), FAILED(失败), WAITING(等待中)
	ProcessStatus string `json:"process_status,omitempty"`
	// UpdateCount 已更新的记录数
	UpdateCount int `json:"update_count,omitempty"`
	// WarnCount 获得警告数量
	WarnCount int `json:"warn_count,omitempty"`
}

// UpdateTime 更新时间信息
type UpdateTime struct {
	// EndTime 更新操作结束的日期和时间。例如, "2021-05-23 16:33:30"
	EndTime model.DateTime `json:"end_time,omitzero"`
	// StartTime 更新操作开始的时间。例如, "2021-05-23 16:33:30"
	StartTime model.DateTime `json:"start_time,omitzero"`
}
