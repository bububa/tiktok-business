package pixel

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Pixel pixel info
type Pixel struct {
	// PixelID Pixel ID
	PixelID string `json:"pixel_id,omitempty"`
	// PixelCode Pixel 代码，即您在事件管理平台看到的 ID
	PixelCode string `json:"pixel_code,omitempty"`
	// PixelName Pixel 名称
	PixelName string `json:"pixel_name,omitempty"`
	// PixelCategory Pixel 追踪事件类型。
	// 枚举值：
	// ONLINE_STORE ：电商事件。
	// FILLING_FORM ：表单事件。
	// CONTACTS ：联系事件。
	// LANDING_PAGE ：APK 下载事件。
	// CUSTOMIZE_EVENTS ：自定义事件。
	//
	// 注意：无论是否为 Pixel 设置了本字段，都不再影响 Pixel 可以追踪的 Pixel 事件。本字段将在下一个 API 版本中废弃。
	PixelCategory enum.PixelCategory `json:"pixel_category,omitempty"`
	// PixelScript Pixel 代码，由系统生成，您需要将 Pixel 代码安装到您的网站以进行事件追踪。见 帮助中心
	PixelScript string `json:"pixel_script,omitempty"`
	// CreateTime 创建时间的时间戳（如 1581679723600）
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// PixelSetupMode Pixel安装模式。枚举值: STANDARD (标准模式), DEVELOPER (开发者模式)。2022年9月15日后使用/pixel/create/接口创建的所有Pixel都是开发者模式。
	PixelSetupMode enum.PixelSetupMode `json:"pixel_setup_mode,omitempty"`
	// PartnerName 合作伙伴名
	PartnerName string `json:"partner_name,omitempty"`
	// AdvancedMatchingFields 高级匹配设置
	AdvancedMatchingFields *AdvancedMatchingFields `json:"advanced_matching_fields,omitempty"`
	// AutomaticAdvancedMatchingFields 该 Pixel 的自动高级匹配设置。
	// 自动高级匹配功能自动使用通过 TikTok Pixel 收集的客户信息，如电子邮箱、姓名和电话号码，从而改善匹配率、转化监测和再营销
	AutomaticAdvancedMatchingFields *AutomaticAdvancedMatchingFields `json:"automatic_advanced_matching_fields,omitempty"`
	// EnableFirstPartyCookies 是否允许第一方 Cookie。
	// 允许第一方 Cookie 将与 TikTok 共享来自你网站的第一方 Cookie 数据，可改善衡量效果并帮助触达更多用户。了解关于通过 TikTok Pixel 使用 Cookie。
	// 默认值： true。
	EnableFirstPartyCookies bool `json:"enable_first_party_cookies,omitempty"`
	// EnableExpandedDataSharing 是否启用扩展数据共享。
	// 启用扩展数据共享将允许此 Pixel 收集来自你的网站的更多信息，包括页面内容（如元标记和商品详情）、访客操作（如点击、按钮互动和使用时长）和网站性能（如页面加载速度和响应能力）。此数据可用于提升 TikTok 的广告投放成效和推广系列成效。你可以通过 TikTok Pixel 助手查看收集的所有数据。了解通过 TikTok Pixel 实现扩展型数据共享功能。
	// 默认值： true。
	EnableExpandedDataSharing bool `json:"enable_expanded_data_sharing,omitempty"`
	// AssetOwnership 资产所有权数据
	AssetOwnership *AssetOwnership `json:"asset_ownership,omitempty"`
	// Events Pixel 事件列表。
	// Pixel 事件数据每 2-4 小时刷新一次。如果事件在四小时内发生，本字段的值可能为空列表（[]）。您可以在四小时后重试，以获取这些事件的数据。
	Events []Event `json:"events,omitempty"`
}

// AdvancedMatchingFields 高级匹配设置
type AdvancedMatchingFields struct {
	// PhoneNumber 电话号码是否用于高级匹配
	PhoneNumber bool `json:"phone_number"`
	// Email 邮件地址是否用于高级匹配
	Email bool `json:"email"`
}

// AutomaticAdvancedMatchingFields 该 Pixel 的自动高级匹配设置。
// 自动高级匹配功能自动使用通过 TikTok Pixel 收集的客户信息，如电子邮箱、姓名和电话号码，从而改善匹配率、转化监测和再营销
type AutomaticAdvancedMatchingFields struct {
	// PhoneNumber 电话号码是否用于高级匹配
	PhoneNumber bool `json:"phone_number"`
	// Email 邮件地址是否用于高级匹配
	Email bool `json:"email"`
	// Name 是否为姓名启用自动高级匹配
	Name bool `json:"name"`
	// Address 是否为地址启用自动高级匹配
	Address bool `json:"address"`
	// ExternalID 是否为外部 ID 启用自动高级匹配
	ExternalID bool `json:"external_id"`
}

// AssetOwnership 资产所有权数据
type AssetOwnership struct {
	// AssetRelationStatus 资产关系状态。
	// 枚举值:
	// null: 该 Pixel 为当前广告主（advertiser_id）所有，且尚未作为资产转移至商务中心中。
	// TRANSFERRED: 该 Pixel 已作为资产转移至商务中心中，但仍为当前广告主（advertiser_id）所有。若想转移 Pixel，可使用/bc/pixel/transfer/。
	// SHARED: 该 Pixel 已作为资产分享给当前广告主（advertiser_id）。对于由他人分享的资产 Pixel ，你可以修改，但是不能删除。
	// UNBOUND: 该 Pixel 之前作为资产分享给当前广告主（advertiser_id），但已与当前广告主解绑。在统计数据中，状态为UNBOUND 的资产不会被计入。
	AssetRelationStatus enum.PixelAssetRelationStatus `json:"asset_relation_status,omitempty"`
	// OwnershipStatus 该 Pixel 是否为当前广告主（advertiser_id）所有。
	// 枚举值：true，false。
	// 若 asset_relation_status 字段的值为 null 或 TRANSFERRED，本字段的值将为 true。
	// 若 asset_relation_status 字段的值为 SHARED 或UNBOUND，本字段的值将为 false。
	OwnershipStatus bool `json:"ownership_status,omitempty"`
	// OwnerBcID 以资产形式拥有该 Pixel 的商务中心 ID。
	// 若 asset_relation_status 为TRANSFERRED，本字段的值代表以资产形式将该 Pixel 转移至的商务中心的 ID。
	// 若 asset_relation_status 不为 TRANSFERRED，本字段的值将为 null。
	OwnerBcID string `json:"owner_bc_id,omitempty"`
}

// Event Pixel 事件
type Event struct {
	// CurrencyValue 转化价值（若创建时不传入则返回""）
	CurrencyValue string `json:"currency_value,omitempty"`
	// Currency 转化价值对应币种，枚举值: INR(印度卢比),JPY(日元),USD(美元)
	Currency string `json:"currency,omitempty"`
	// Name 事件名称
	Name string `json:"name,omitempty"`
	// Deprecated 该事件是否已被废弃
	Deprecated bool `json:"deprecated,omitempty"`
	// EventType 仅当事件为标准网站事件（Pixel 事件）时返回。
	// 标准网站事件类型。
	// 本字段的枚举值参见Pixel 事件类型- Pixel 事件小节中的“用于广告创编的事件名称”列。
	EventType enum.OptimizationEvent `json:"event_type,omitempty"`
	// OptimizationEvent 标准网站事件（event_type）对应的转化事件类型。您可在使用/adgroup/create/创建广告组时将转化事件传入optimization_event字段。
	// 本字段的枚举值参见Pixel 事件类型- Pixel 事件小节中的“用于广告创编的事件名称”列。
	// 注意：标准网站事件PAGE_VIEW 不可用于优化。因此， event_type 为 PAGE_VIEW 时，本字段值将为 null。
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// CustomEventType 仅当事件为自定义网站事件时返回。
	// 自定义网站事件类型。 自定义事件是 TikTok 合作伙伴可以在预定义的标准事件列表外自定义的行为。若想了解自定义事件的详情，请查看标准事件和自定义事件。
	// 不支持使用/pixel/event/create/创建自定义网站事件，但支持使用事件 1.0 接口/pixel/track/或/pixel/batch/，或事件 2.0 接口/event/track/直接回传自定义网站事件。 例如，若回传 SHOPPING 事件，则该事件视作自定义事件，因为该事件类型未在事件 API 1.0 - 报告网站事件和参数 - 支持的网站事件或事件 API 2.0 - 支持的事件 - 网站标准事件中的“事件”列列出。
	CustomEventType string `json:"custom_event_type,omitempty"`
	// EventCode 事件代码
	EventCode string `json:"event_code,omitempty"`
	// EventID 事件 ID
	EventID string `json:"event_id,omitempty"`
	// StatisticType 统计类型，枚举值: EVERY_TIME（每一次）,ONCE（仅一次）
	StatisticType enum.PixelEventStatisticType `json:"statistic_type,omitempty"`
	// Rules 追踪事件规则，子字段规则及含义见 Pixel事件类型-事件规则类型
	Rules []EventRule `json:"rules,omitempty"`
}

// EventRule 追踪事件规则，子字段规则及含义见 Pixel事件类型-事件规则类型
type EventRule struct {
	// Operator 操作符，枚举值: OPERATORTYPE_CONTAINS,OPERATORTYPE_DOES_NOT_EQUAL,OPERATORTYPE_EQUALS
	Operator enum.PixelEventRuleOperator `json:"operator,omitempty"`
	// Trigger 事件触发类型，枚举值: TRIGGERTYPE_CLICK,TRIGGERTYPE_PAGEVIEW
	Trigger enum.PixelEventRuleTrigger `json:"trigger,omitempty"`
	// Value 变量值
	Value string `json:"value,omitempty"`
	// Variable 变量类型，枚举值: ELEMENT,PAGE_HOSTNAME,PAGE_PATH,PAGE_URL
	Variable enum.PixelEventRuleVariable `json:"variable,omitempty"`
}
