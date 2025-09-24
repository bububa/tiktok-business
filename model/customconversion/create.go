package customconversion

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建自定义转化
type CreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Name 自定义转化的名称。
	// 名称必须以字母开头，可以包含字母、数字、空格、下划线或横线。
	//
	// 长度限制: 50 个字符，不能包含表情符号。
	//
	// 注意:
	//
	// 不要指定标准事件的名称。
	// 不允许在广告账户内重复使用自定义转化名称。
	Name string `json:"name,omitempty"`
	// Description 自定义转化的描述。
	// 长度限制: 500 个字符，不能包含表情符号。
	Description string `json:"description,omitempty"`
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
	// OptimizationEvent 要与自定义转化关联的标准事件。
	// 枚举值参阅optimization_event 的可选值。
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// CustomEventType 仅当 event_source_type 为 PIXEL 且已为 Pixel (event_source_id) 配置自定义事件时生效。
	// 要与自定义转化关联的自定义事件。
	// 当为自定义转化指定自定义事件时，需要同时通过 optimization_event 指定要映射到的标准事件。
	// 自定义事件是 TikTok 合作伙伴可以在预定义的标准事件列表之外自行定义的操作。要了解有关自定义事件的更多信息，参阅关于 TikTok 广告管理平台中的自定义事件。
	// 要获取为 Pixel 配置的自定义事件，可使用/pixel/list/并检查返回的custom_event_type。
	CustomEventType string `json:"custom_event_type,omitempty"`
	// Rules 自定义转化的转化规则详细信息。
	// 只有满足这些条件的操作才会计为自定义转化。
	//
	// 数量范围: 1-10
	Rules []CustomConversionRule `json:"rules,omitempty"`
}

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建自定义转化 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// CustomConversionID 自定义转化的 ID
		CustomConversionID string `json:"custom_conversion_id,omitempty"`
	} `json:"data"`
}
