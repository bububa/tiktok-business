package diagnostic

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// EventSourceMetricRequest 获取商品库事件趋势及匹配率 API Request
type EventSourceMetricRequest struct {
	// BcID 商务中心ID。该商务中心需为所指定商品库的所属商务中心，或所指定商品库已作为资产分享给该商务中心
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID。您需有访问该商品库的权限。
	CatalogID string `json:"catalog_id,omitempty"`
	// EventSourceType 事件源类型。
	// 枚举值：
	// APP: App。需同时传入参数 app_id。
	// PIXEL: Pixel。需同时传入参数 pixel_code。
	EventSourceType enum.EventSourceType `json:"event_source_type,omitempty"`
	// AppID event_source_type 为 APP 时必填。
	// App ID。
	// 若想获取与商品库（catalog_id）绑定的 App ID，可使用 /catalog/eventsource_bind/get/。
	AppID string `json:"app_id,omitempty"`
	// PixelCode event_source_type 为 PIXEL 时必填。
	// Pixel 代码。
	// 若想获取与商品库（catalog_id）绑定的 Pixel 代码，可使用 /catalog/eventsource_bind/get/。
	PixelCode string `json:"pixel_code,omitempty"`
	// EventType 您想要获取对应数据的事件类型。
	// 枚举值：
	// VIEW_CONTENT：查看详情。
	// ADD_TO_CART：加入购物车。
	// PURCHASE：支付完成。
	// 默认值： VIEW_CONTENT。
	EventType enum.EventType `json:"event_type,omitempty"`
	// TimeRange 您想要获取对应数据的时间范围。
	// 枚举值：
	// YESTERDAY：近1天。
	// LAST_7_DAYS：近7天。
	// LAST_30_DAYS：近30天。
	// 默认值：LAST_7_DAYS
	TimeRange string `json:"time_range,omitempty"`
}

// Encode implements GetRequest interface
func (r *EventSourceMetricRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("catalog_id", r.CatalogID)
	values.Set("bc_id", r.BcID)
	values.Set("event_source_type", string(r.EventSourceType))
	if r.AppID != "" {
		values.Set("app_id", r.AppID)
	}
	if r.PixelCode != "" {
		values.Set("pixel_code", r.PixelCode)
	}
	if r.EventType != "" {
		values.Set("event_type", string(r.EventType))
	}
	if r.TimeRange != "" {
		values.Set("time_range", r.TimeRange)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// EventSourceMetricResponse 获取商品库事件趋势及匹配率 API Response
type EventSourceMetricResponse struct {
	model.BaseResponse
	Data struct {
		// List 接收到的事件的趋势及可用于在营销的事件数量
		List []EventSourceMetric `json:"list,omitempty"`
	} `json:"data"`
}

// EventSourceMetric 接收到的事件的趋势及可用于在营销的事件数量
type EventSourceMetric struct {
	// AvailableType 事件类别。
	// 枚举值：
	// EVENT_RECEIVED：所有事件。
	// EVENT_WITH_CONTENT_ID：带有内容 ID（content ID）的事件。
	// EVENT_WITH_CONTENT_ID_MATCHING_INVENTORY：内容 ID 能匹配库存的事件。
	AvailableType enum.EventSourceAvailableType `json:"available_type,omitempty"`
	// EventDetails 事件类别（available_type）的详情
	EventDetails []EventDetail `json:"event_details,omitempty"`
}

// EventDetail 事件类别（available_type）的详情
type EventDetail struct {
	// Date 日期，格式为"YYYY-MM-DD"。
	// 示例： "2023-12-30"。
	Date string `json:"date,omitempty"`
	// Count 在该日期接收到的属于此事件类别的事件数量。
	// 示例： "245076"。
	// 例如， available_type 为 EVENT_WITH_CONTENT_ID_MATCHING_INVENTORY 时，本字段的值为"245076"代表当日收到了 245,076 个内容 ID 与您的商品库中的商品（SKU ID）匹配的事件。
	Count model.Int64 `json:"count,omitempty"`
	// Percentage 当日收到的属于此事件类别的事件数量占当日收到的总事件数量的百分比。
	// available_type 为 EVENT_WITH_CONTENT_ID_MATCHING_INVENTORY 时，本字段的值代表商品库匹配率。匹配率的计算方法是：与商品库中的商品匹配的信号事件数除以与 TikTok 共享的信号事件总数。
	// 我们建议您将匹配率保持在 90% 以上，以确保视频购物广告获得理想的投放效果。
	// 示例："100"。
	// 注意：available_type 为 EVENT_RECEIVED 时，本字段的值始终为"100"。
	Percentage model.Float64 `json:"percentage,omitempty"`
}
