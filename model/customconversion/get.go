package customconversion

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取自定义转化的详情 API Request
type GetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomConversionID 自定义转化的 ID
	CustomConversionID string `json:"custom_conversion_id,omitempty"`
	// EventSourceType 事件源类型。
	// 枚举值:
	// PIXEL: Pixel。
	// APP: 应用。
	EventSourceType enum.EventSourceType `json:"event_source_type,omitempty"`
	// EventSourceID 事件源 ID。
	// 当 event_source_type 为 PIXEL 时，指定一个 Pixel ID。
	// 要获取广告账户中的 Pixel ID，可使用/pixel/list/并检查返回的pixel_id。
	// 当 event_source_type 为 APP 时，指定自归因网络应用的应用平台 ID。
	// 要获取广告账户中自归因应用的应用平台 ID，可使用/app/list/。确认应用的 self_attribution_enabled 返回值为 true 并检查返回的 app_platform_id。
	EventSourceID string `json:"event_source_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("custom_conversion_id", r.CustomConversionID)
	values.Set("event_source_type", string(r.EventSourceType))
	values.Set("event_source_id", r.EventSourceID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取自定义转化的详情 API Response
type GetResponse struct {
	model.BaseResponse
	Data *CustomConversion `json:"data,omitempty"`
}
