package rf

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/adgroup"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建覆盖和频次广告组 API Request
type CreateRequest struct {
	// RequestID 请求ID，用于创建同名广告组。该ID支持接口幂等，避免重复请求。
	// 该ID和返回参数中的 request_id 不同。返回的request_id用于唯一标识一次 HTTP 请求。
	// 本字段的值需为字符串格式的64位整数值。
	// 示例： "123456789"
	RequestID string `json:"request_id,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 广告组所属的推广系列
	CampaignID string `json:"campaign_id,omitempty"`
	// ShareDisabled 是否禁止将本广告组中的广告分享到第三方平台。默认值：False
	ShareDisabled bool `json:"share_disabled,omitempty"`
	// AdgroupName 广告组名称。长度不得超过100字符或包含表情符号
	AdgroupName string `json:"adgroup_name,omitempty"`
	// PromotionType 推广类型。枚举值: APP_ANDROID, APP_IOS, WEBSITE, WEBSITE_OR_DISPLAY。如果推广目标（objective_type）是RF_REACH，本字段需设置为WEBSITE_OR_DISPLAY
	PromotionType enum.PromotionType `json:"promotion_type"`
	// OptimizationEvent 本广告组的转化事件。当推广对象是App且有多个追踪URL时必填。支持的App事件参见附录-转化事件。您也可以使用/app/external_action/接口查看您的App支持的转化事件
	OptimizationEvent enum.OptimizationEvent `json:"optimization_event,omitempty"`
	// AppID 推广的移动应用ID
	AppID string `json:"app_id,omitempty"`
	// CommentDisabled 是否禁止用户在TikTok上对您的广告发表评论。默认值: False
	CommentDisabled bool `json:"comment_disabled,omitempty"`
	// AudienceIDs 想要定向的受众ID列表。若要将受众在覆盖和频次广告中定向或排除，需要在创建或更新受众时，将audience_sub_type设置为REACH_FREQUENCY。
	// 注意：不能将相似受众用于覆盖和频次广告，否则会出错
	AudienceIDs []string `json:"audience_ids,omitempty"`
	// ExcludedAudienceIDs 要排除的受众ID列表。若要将受众在覆盖和频次广告中定向或排除，需要在创建或更新受众时，将audience_sub_type设置为REACH_FREQUENCY。
	// 注意：不能将相似受众用于覆盖和频次广告，否则会出错。
	ExcludedAudienceIDs []string `json:"excluded_audience_ids,omitempty"`
	// AgeGroups 受众年龄区间。
	// 若定向地区包括以色列和巴西，则此字段不可设置为AGE_13_17。
	// 枚举值详见枚举值-受众年龄区间。
	// 在某些场景下，不允许创建定向美国 13-17 年龄区间（AGE_13_17）受众的广告组。若想了解详情，请查看 TikTok 广告的新增年龄定向限制。
	// 若您使用了与美国 13-17 年龄区间定向不兼容的定向设置，且将 age_groups 设置为 [] 或不传入字段 age_groups，将出现报错。您需在该字段中明确指定您想要定向的具体年龄区间，例如AGE_18_24，从而避免发生报错。
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// Gender 受众性别。枚举值详见枚举值 -受众性别
	Gender enum.GENDER `json:"gender,omitempty"`
	// Languages 受众语言。枚举值详见枚举值-受众语言
	Languages []string `json:"languages,omitempty"`
	// LocationIDs 定向地域ID。所定向地域可为国家或地区层级，省、指定市场区域或都会区层级，或城市层级地域。
	// 只支持定向到单个国家或地区。如果传入了多个ID，则这些ID必须属于同一国家或地区。
	// 支持投放覆盖和频次广告的国家和地区请参见关于覆盖和频次广告。
	LocationIDs []string `json:"location_ids,omitempty"`
	// IsHfss 广告推广对象是否是HFSS食品（高脂肪、高盐和高糖的食品）。注意：欧洲国家不允许向未成年人推送 HFSS 食品广告
	IsHfss bool `json:"is_hfss,omitempty"`
	// OperatingSystems 操作系统。枚举值: ANDROID, IOS, PC
	OperatingSystems []enum.OperatingSystem `json:"operating_systems,omitempty"`
	// NetworkTypes 网络情况。默认值：unlimited (不限）。枚举值详见枚举值-网络类型
	NetworkTypes []enum.NetworkType `json:"network_types,omitempty"`
	// DeviceModelIDs 想要定向的设备机型ID列表。目前只支持定向到品牌层级。传入本字段时，设备价格字段device_price_ranges不可同时传入。所有设备机型ID的枚举值可通过获取设备机型列表接口获取。
	// 注意: 创建广告时需确保对应设备处于活跃状态（ /tool/device_model/ 接口返回中is_active 为 True）。
	DeviceModelIDs []string `json:"device_model_ids,omitempty"`
	// DevicePriceRanges 设备价格区间（10000代表1000+）。该数字必须是50的倍数。
	// 重要提示: 最终的实际上限将在您设定的上限基础上增加50，您可以在TikTok广告管理平台的广告组设置中看到实际上限。例如，如果您设置的价格区间是[0, 250]，实际区间则为[0, 300]。
	DevicePriceRanges []int `json:"device_price_ranges,omitempty"`
	// CarrierIDs 运营商 ID列表。只有在对应的in_use字段为true时，该运营商的定向才有效。所选运营商需要和定向国家对应。您可以使用获取运营商 接口获取运营商ID
	CarrierIDs []string `json:"carrier_ids,omitempty"`
	// InterestCategoryIDs 兴趣分类。可使用/tool/interest_category/ 接口获取兴趣分类的完整列表
	InterestCategoryIDs []string `json:"interest_category_ids,omitempty"`
	// RfPurchasedType 购买方式。枚举值：
	// FIXED_SHOW： 按固定展示量购买。
	// FIXED_REACH： 按固定覆盖人数购买。
	// FIXED_BUDGE: 按固定预算购买。
	RfPurchasedType enum.RfPurchasedType `json:"rf_purchased_type,omitempty"`
	// Budget 预算金额。请参考库存预估接口的返回结果填写
	Budget float64 `json:"budget,omitempty"`
	// PurchasedImpression 要购买的展示量，以1000为单位。请参考库存预估接口的返回结果填写
	PurchasedImpression int64 `json:"purchased_impression,omitempty"`
	// PurchasedReach 要购买的覆盖人数，以1000为单位。请参考库存预估接口的返回结果填写。
	PurchasedReach int64 `json:"purchased_reach,omitempty"`
	// ScheduleStartTime 排期开始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	// schedule_end_time 与 schedule_start_time 之间的时间间隔不能超过90天。
	// 本字段的值应设置为投放地时区的YYYY-MM-DD HH:00:00。
	// 例如，若您是 UTC-5 时区的广告主，并希望创建一个排期结束时间为2023年11月1日的覆盖和频次广告组，本字段需设置为 "2023-11-01 05:00:00"
	ScheduleStartTime string `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 排期结束时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	// schedule_end_time 与 schedule_start_time 之间的时间间隔不能超过90天。
	// 本字段的值应设置为投放地时区的YYYY-MM-DD HH:59:59。
	// 例如，若您是 UTC-5 时区的广告主，并希望创建一个排期结束时间为2023年11月10日的覆盖和频次广告组，本字段需设置为 "2023-11-10 04:59:59"。
	ScheduleEndTime string `json:"schedule_end_time,omitzero"`
	// Frequency 频次。本字段与 frequency_schedule 共同定义用户看到广告的频次上限（仅适用于覆盖和频次广告）。
	// 例如，frequency = 2 且frequency_schedule = 3， 表示每 3 天至多 2 次展示。
	// 需满足以下条件：
	// 1 ≤ frequency ≤ frequency_schedule
	// 1 ≤ frequency_schedule ≤ min(schedule_end_time - schedule_start_time, 30)
	// 由frequency 和 frequency_schedule 定义的频次上限不可超过每天 4 次展示。
	Frequency int `json:"frequency,omitempty"`
	// FrequencySchedule 频次天数。与 frequency 共同定义用户看到广告的频次上限（仅适用于覆盖和频次广告）。请查看frequency字段的描述
	FrequencySchedule int `json:"frequency_schedule,omitempty"`
	// OptimizationGoal 优化目标。 必须与推广系列层级的 objective_type 保持对应。枚举值: REACH, VIDEO_VIEW, CLICK, POST_ENGAGEMENT, INSTALL
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// CPVVideoDuration 目标视频播放时长。
	// 枚举值：SIX_SECONDS(视频播放6秒)，TWO_SECONDS(视频播放2秒)
	CPVVideoDuration enum.CPVVideoDuration `json:"cpv_video_duration,omitempty"`
	// BrandSafetyType 覆盖和频次广告的品牌安全类型。默认值：NO_BRAND_SAFETY。枚举值：
	// NO_BRAND_SAFETY：不使用任何品牌安全解决方案，即使用全部库存。您的广告可能会在某些包含成人主题的内容周围展示。
	// EXPANDED_INVENTORY：使用TikTok的品牌安全解决方案，即使用扩展库存。您的广告将不会出现在不当内容或包含成人主题的内容后面。在下个API版本中EXPANDED_INVENTORY将替代NO_BRAND_SAFETY，并成为默认的品牌安全选项。
	// STANDARD_INVENTORY：使用TikTok的品牌安全解决方案中的标准库存。您的广告将展示在适合大多数品牌的内容周围。
	// LIMITED_INVENTORY：使用TikTok的品牌安全解决方案中的有限库存。您的广告将展示在不包含成人主题的内容周围。
	// THIRD_PARTY：使用第三方品牌安全解决方案。
	// 您可以使用/tool/region/接口查询品牌安全设置对应的广告可投放国家或地区。
	// 注意：
	// 出价前第三方品牌安全解决方案目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	BrandSafetyType enum.BrandSafetyType `json:"brand_safe_type,omitempty"`
	// ExcludeCategoryIDs 仅当brand_safety_type 为 STANDARD_INVENTORY或LIMITED_INVENTORY时可用。
	// 内容排除类别ID。 您可以使用/tool/content_exclusion/get/来获取内容类别ID列表（excluded_category_list），ID可以用于指定某内容类别，避免您的广告出现在该内容类型周围。
	ExcludeCategoryIDs []string `json:"excluded_category_ids,omitempty"`
	// VideoDownloadDisabled 用户是否可以在 TikTok 上下载您的广告视频（该字段设置后不允许修改）。 默认值: false
	VideoDownloadDisabled bool `json:"video_download_disabled,omitempty"`
	// FeedType 信息流类型。枚举值: STANDARD_FEED, TOP_FEED。STANDARD_FEED指库存高，且CPM价格正常。TOP_FEED 指库存受限，且CPM价格较高
	FeedType enum.FeedType `json:"feed_type,omitempty"`
	// DeliveryMode 广告投放的排序与排期策略。
	// 枚举值:
	// STANDARD：标准投放。每个广告均匀投放，预计会获得大致相同的访问量。
	// SCHEDULE：计划投放。为每个广告设置投放的特定时间段。若您设置为此值，您需同时传入对象数组schedules。
	// SEQUENCE：顺序投放。为广告设置特定的投放顺序。若您设置为此值，您需同时传入字段 expected_orders 。
	// OPTIMIZE（已废弃）：优选投放。
	// 默认值：STANDARD。
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// ScheduleInfos 广告投放信息。
	// 若 delivery_mode 设置为 SEQUENCE，则本数组中包含的对象数量代表您要设置对应投放顺序的广告数量。
	ScheduleInfos []adgroup.ScheduleInfo `json:"schedule_infos,omitempty"`
	// ContextualTagIDs 当brand_safety_type设置为THIRD_PARTY时，不支持本字段。
	// 当feed_type设置为TOP_FEED时，不支持本字段。
	// 内容相关定向标签ID列表。您可使用 /tool/contextual_tag/get/获取可用的内容相关定向标签。
	// 注意：
	// 如果您在/campaign/create/中将rf_campaign_type设置为PULSE，则在创建广告组时，必须填入PREMIUM标签类型下MAX_PULSE或CUSTOM内容分组类型的内容相关定向标签。GENERAL类型的内容相关定向标签非必填。请参阅TikTok Pulse以了解使用 TikTok Pulse 的更多信息。
	// 目前，您可以在 AE（阿联酋）、AU（澳大利亚）、BR（巴西）、CA（加拿大）、DE（德国）、ES（西班牙）、FR（法国）、GB（英国）、IT（意大利）、MX（墨西哥）、SA（沙特阿拉伯）、TR（土耳其） 和 US（美国）使用 Max Pulse，以及在US(美国）、CA（加拿大）、BR（巴西）、AU（澳大利亚）、GB（英国）、FR（法国）、IT（意大利）、ES（西班牙）、DE（德国）使用类别分组。Max Pulse 和类别分组为白名单功能。如果您想在上述市场使用该两种功能，请联系您的 TikTok 销售代表。
	// 该功能目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	ContextualTagIDs []string `json:"contextual_tag_ids,omitempty"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建覆盖和频次广告组 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Adgroup `json:"data,omitempty"`
}
