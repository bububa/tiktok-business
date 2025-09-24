package customconversion

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// CustomConversion 自定义转化
type CustomConversion struct {
	// CustomConversionID 自定义转化的 ID
	CustomConversionID string `json:"custom_conversion_id,omitempty"`
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
	// ActiveStatus 自定义转化的活跃状态。
	// 枚举值:
	// NO_RECENT_ACTIVITY: 近期无活动。已收到回传活动，但并非是在过去 7 天内收到。
	// ACTIVE: 活跃。在过去 7 天内收到回传活动。
	// WAITING_FOR_ACTIVITY: 不活跃。过去 90 天内未收到回传活动。
	ActiveStatus enum.CustomConversionActiveStatus `json:"active_status,omitempty"`
	// OptimizationEvent 要与自定义转化关联的标准事件。
	// 枚举值参阅optimization_event 的可选值。
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// Rules 自定义转化的转化规则详细信息。
	// 只有满足这些条件的操作才会计为自定义转化。
	//
	// 数量范围: 1-10
	Rules []CustomConversionRule `json:"rules,omitempty"`
	// CreateTime 自定义转化的创建时间，格式为YYYY-MM-DD HH:MM:SS(UTC+0)。
	// 示例: 2025-01-01 00:00:01。
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// UpdateTime 自定义转化上次更新的时间，格式为YYYY-MM-DD HH:MM:SS(UTC+0)。
	// 示例: 2025-01-01 00:00:01。
	UpdateTime model.DateTime `json:"update_time,omitzero"`
}

// CustomConversionRule 自定义转化的转化规则详细信息。
type CustomConversionRule struct {
	// Parameter 参数。
	// 枚举值:
	// BRAND: brand。
	// CONTENT_CATEGORY: content_category。
	// CONTENT_ID: content_id。
	// CONTENT_NAME: content_name。
	// CONTENT_TYPE: content_type。
	// CURRENCY: currency。
	// DESCRIPTION: description。
	// EVENT_ID: event_id。
	// NUM_ITEMS: num_items。
	// PRICE: price。
	// SEARCH_STRING: search_string。
	// STATUS: status。
	// URL: url。
	// VALUE: value。
	Parameter enum.CustomConversionRuleParameter `json:"parameter,omitempty"`
	// Operator 运算符。
	// 当 parameter 为 BRAND、CONTENT_CATEGORY、CONTENT_ID、CONTENT_NAME、CONTENT_TYPE、DESCRIPTION、EVENT_ID、SEARCH_STRING、STATUS 或 URL 时，枚举值为:
	// CONTAIN: 包含。
	// NOT_CONTAIN: 不包含。
	// START_WITH: 开头是。
	// EQUAL: 等于。
	// NOT_EQUAL: 不等于。
	// 当 parameter 为 NUM_ITEMS、PRICE 或 VALUE 时，枚举值为:
	// EQUAL: 等于。
	// NOT_EQUAL: 不等于。
	// LT: 小于。
	// GT: 大于。
	// LTE: 小于或等于。
	// GTE: 大于或等于。
	// 当 parameter 为 CURRENCY 时，枚举值为:
	// EQUAL: 等于。
	// NOT_EQUAL: 不等于。
	Operator enum.FilterOperator `json:"operator,omitempty"`
	// Values 参数的值。
	// 当 parameter 为 BRAND、CONTENT_CATEGORY、CONTENT_ID、CONTENT_NAME、CONTENT_TYPE、DESCRIPTION、EVENT_ID、SEARCH_STRING、STATUS、URL 时，该字段将是一个字符串数组，每个字符串的最大长度为 100 个字符，最多 100 个条目。
	// 当 parameter 为 NUM_ITEMS 时，该字段将是一个介于 0 和 1,000,000,000 之间的整数。
	// 当 parameter 为 PRICE、VALUE 时，该字段将是一个保留两位小数的浮点数，取值范围为 0.00到 1,000,000,000.00。
	// 当 parameter 为 CURRENCY 时，该字段将是一个字符串数组。要了解枚举值，请参阅的可选值。
	Values []string `json:"values,omitempty"`
}
