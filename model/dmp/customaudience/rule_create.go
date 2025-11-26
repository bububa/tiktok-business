package customaudience

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// RuleCreateRequest 通过规则创建受众 API Request
type RuleCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceName 受众名称。
	// 长度限制：128 个字符。
	CustomAudienceName string `json:"custom_audience_name,omitempty"`
	// AudienceType 受众类型。
	// 枚举值：
	// ENGAGEMENT：广告互动受众（资产类型为广告组）。
	// ENGAGEMENT_ORGANIC_VIDEO：自然流量视频互动受众（资产类型为公开视频）。
	// ENGAGEMENT_LIVE_VIDEO：直播视频互动受众（资产类型为直播视频）。
	// APP：应用事件受众。
	// PIXEL：网站访客受众。
	// LEAD_GENERATION：线索收集受众。
	// BUSINESS_ACCOUNT：企业号受众。
	// TIKTOK_SHOP ：商店活动受众。
	// OFFLINE：线下活动受众。
	AudienceType enum.AudienceType `json:"audience_type,omitempty"`
	// AudienceSubType 受众子类型，表明该受众可用于的广告类型。
	// 枚举值：
	// NORMAL：常规受众，即可用于非覆盖和频次广告的受众。
	// REACH_FREQUENCY：覆盖和频次广告受众，即只可用于覆盖和频次广告。
	// 默认值：NORMAL。
	AudienceSubType enum.AudienceSubType `json:"audience_sub_type,omitempty"`
	// RetentionInDays 保留该受众的天数。
	// 取值范围：1-365。
	// 注意：
	// 若设置了本字段，受众将在距离创建时间的指定保留天数后过期，对该受众的任何操作均不会重置过期日期。
	// 若未设置本字段，受众将在连续12个月未应用于任何活跃广告组且未进行任何修改后过期，将该受众应用于活跃广告组或修改该广告组会重置过期日期。欲了解会重置过期日期的操作，请查看帮助中心文档受众过期政策。
	RetentionInDays int `json:"retention_in_days,omitempty"`
	// IsAutoRefresh 是否开启受众自动刷新功能。
	// 支持的值：true，false。
	// 默认值：true。
	// 如果开启该功能，受众将会根据您选择的天数进行自动更新。如果关闭该功能，受众将会保持不变。
	IsAutoRefresh *bool `json:"is_auto_refresh,omitempty"`
	// IdentityID 当 audience_type 为 ENGAGEMENT_LIVE_VIDEO 或 ENGAGEMENT_ORGANIC_VIDEO 时本字段必填。
	// 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 当 audience_type 为 ENGAGEMENT_LIVE_VIDEO 或 ENGAGEMENT_ORGANIC_VIDEO 时本字段必填。
	// 认证身份类型。
	// 枚举值：TT_USER，BC_AUTH_TT。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID 当 audience_type为BC_AUTH_TT时本字段必填。
	// 与添加到商务中心的 TikTok 用户这一认证身份绑定的商务中心的 ID
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// RuleSpec 受众的规则
	RuleSpec *RuleSpec `json:"rule_spec,omitempty"`
}

// RuleSpec 受众的规则
type RuleSpec struct {
	// IncludeRuleSet 包含规则集的详情
	IncludeRuleSet *RuleSet `json:"include_rule_set,omitempty"`
	// ExcludeRuleSet 排除规则集的详情
	ExcludeRuleSet *RuleSet `json:"exclude_rule_set,omitempty"`
}

// RuleSet 规则集
type RuleSet struct {
	// Operator 包含规则集中的包含规则之间的操作符。
	// 枚举值：OR。
	// 若您指定了一条以上的包含规则，则包含规则间将通过 OR 逻辑组合，从而扩大受众范围。
	Operator enum.FilterSetOperator `json:"operator,omitempty"`
	// Rules 包含规则集中的包含规则列表
	Rules []Rule `json:"rules,omitempty"`
}

// Rule 规则
type Rule struct {
	// EventSourceIDs 当audience_type不为ENGAGEMENT或LEAD_GENERATION时本字段必填。
	// 规则的事件源 ID 列表。
	// 对于广告互动受众，需要填入广告组 ID。不填默认使用所有事件源 ID。
	// 对于自然流量视频互动受众，使用 TikTok 帖子 ID 作为事件源 ID。您可以使用/identity/video/get/ 接口获取 TikTok 帖子 ID 。最多填入10个 TikTok 帖子 ID 。
	// 对于直播互动受众，使用直播视频 ID 作为事件源 ID。 您可以使用/identity/live/get/接口获取直播视频 ID。最多填入10个直播视频 ID。
	// 对于应用事件受众，需要填入 App ID。
	// 对于网站访客受众，需要填入Pixel ID。
	// 对于线索收集受众，不能填入该字段，否则会出现报错。不填默认使用所有事件源 ID。
	// 对于商店活动受众，填入 TikTok Shop ID。您可以使用 /store/list/接口获取 TikTok Shop ID。
	// 对于企业号受众，需要填入广告主的核心用户 ID。您可以使用 /user/info/ 接口获取核心用户 ID。
	// 对于线下活动受众，使用线下事件组 ID 作为事件源 ID。您可以使用/offline/get/ 获取线下事件组 ID。
	EventSourceIDs []string `json:"event_source_ids,omitempty"`
	// RetentionDays 受众回溯期。
	// 枚举值详情可查看枚举值 - 回溯期。
	// 注意：
	// 如果audience_type为BUSINESS_ACCOUNT，且filters对象中的value不为BUSINESS ACCOUNT PROFILE FOLLOW，则retention_days必须为7，14或30。
	// 如果audience_type为ENGAGEMENT_LIVE_VIDEO 或ENGAGEMENT_ORGANIC_VIDEO，retention_days必须为7，14或30。
	// 为避免retention_days和retention_in_days字段的混淆，我们将在下个API版本修改retention_days的字段名。
	RetentionDays int `json:"retention_days,omitempty"`
	// FilterSet 每个包含规则集中的筛选条件集
	FilterSet *model.FilterSet `json:"filter_set,omitempty"`
}

// Encode implements PostRequest
func (r *RuleCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// RuleCreateResponse 通过规则创建受众 API Response
type RuleCreateResponse struct {
	model.BaseResponse
	Data struct {
		// CustomAudienceID 自定义受众 ID
		CustomAudienceID string `json:"custom_audience_id,omitempty"`
	} `json:"data"`
}
