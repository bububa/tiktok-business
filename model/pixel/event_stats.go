package pixel

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// EventStatsRequest 获取Pixel事件数据 API Request
type EventStatsRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PixelIDs Pixel ID 列表。
	// 最大数量：10。
	PixelIDs []string `json:"pixel_ids,omitempty"`
	// DateRange 时间范围，不得超过 30 天。
	// 注意：从此接口获取的事件数据与事件管理平台显示的事件数据可能会存在轻微差异。这是因为此接口获取的事件数据基于 UTC 时间，而事件管理平台显示的事件数据基于广告主的时区。
	DateRange *EventStatsDateRange `json:"date_range,omitempty"`
}

type EventStatsDateRange struct {
	// StartDate 起始时间（UTC 时间）。
	// 示例： "2020-02-16"。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间（UTC 时间）。
	// 示例： "2020-02-17"。
	EndDate string `json:"end_date,omitempty"`
}

// Encode implements GetRequest interface
func (r *EventStatsRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("pixel_ids", string(util.JSONMarshal(r.PixelIDs)))
	values.Set("date_range", string(util.JSONMarshal(r.DateRange)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// EventStatsResponse 获取Pixel事件数据 API Response
type EventStatsResponse struct {
	model.BaseResponse
	Data struct {
		// List 数据列表
		List []PixelEventStats `json:"list,omitempty"`
	} `json:"data"`
}

type PixelEventStats struct {
	// PixelID Pixel ID
	PixelID string `json:"pixel_id,omitempty"`
	// Statistics 统计信息
	Statistics []EventStats `json:"statistics,omitempty"`
}

// EventStats 统计信息
type EventStats struct {
	// PixelEventType 仅当事件为标准网站事件（Pixel 事件）时返回。
	// 标准网站事件类型。
	// 本字段的枚举值参见Pixel 事件类型- Pixel 事件小节中的“用于广告创编的事件名称”列。
	PixelEventType enum.EventType `json:"pixel_event_type,omitempty"`
	// CustomEventType 仅当事件为自定义网站事件时返回。
	// 自定义网站事件类型。 自定义事件是 TikTok 合作伙伴可以在预定义的标准事件列表外自定义的行为。若想了解自定义事件的详情，请查看标准事件和自定义事件。
	// 不支持使用/pixel/event/create/创建自定义网站事件，但支持使用事件 1.0 接口/pixel/track/或/pixel/batch/，或事件 2.0 接口/event/track/直接回传自定义网站事件。 例如，若回传 SHOPPING 事件，则该事件视作自定义事件，因为该事件类型未在事件 API 1.0 - 报告网站事件和参数 - 支持的网站事件或事件 API 2.0 - 支持的事件 - 网站标准事件中的“事件”列列出。
	CustomEventType string `json:"custom_event_type,omitempty"`
	// AttributedCount 归因事件数（归因到您的广告上的事件数）
	AttributedCount int64 `json:"attributed_count,omitempty"`
	// PreviewCount 预览事件数（广告预览产生的事件数）
	PreviewCount int64 `json:"preview_count,omitempty"`
	// TotalCount 总事件数，即从所有数据源接收到的该事件类型（pixel_event_type 或 custom_event_type）的网站事件总数。
	// 公式： total_count= browser_event_total_count + server_event_total_count.
	TotalCount int64 `json:"total_count,omitempty"`
	// BrowserEventTotalCount 从网络浏览器接收到的该事件类型的网站事件数。
	BrowserEventTotalCount int64 `json:"browser_event_total_count,omitempty"`
	// ServerEventTotalCount 从服务器接收到的该事件类型的网站事件数。
	ServerEventTotalCount int64 `json:"server_event_total_count,omitempty"`
}
