package campaign

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/ad/aco"
	"github.com/bububa/tiktok-business/util"
)

// CopyTaskCreateRequest 创建推广系列异步复制任务 API Request
type CopyTaskCreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 您想要复制的推广系列的 ID。
	// 注意：
	// 不支持复制与自定义 TikTok 作品绑定的 GMV Max 推广系列。若想检查 GMV Max 推广系列是否绑定了自定义 TikTok 作品，可使用 /campaign/gmv_max/info/并检查返回的 custom_anchor_video_list 是否为 []。
	// 使用未在 MMP（移动监测合作伙伴）侧启用 SAN （自归因） 模块的 App 的 iOS 专属推广系列不支持复制。若想确保您的 App 启用 TikTok 自归因集成，可查看如何将新应用集成到 SAN 和如何将现有应用迁移到 SAN。
	// 自 2024 年 9 月 30 日起，包含投放至 TikTok 版位的试玩广告的推广系列不再支持复制。
	// 仅支持复制使用以下推广目标的推广系列：REACH，TRAFFIC，VIDEO_VIEWS，ENGAGEMENT，APP_PROMOTION，LEAD_GENERATION 和WEB_CONVERSIONS。
	// 不支持复制推广目标为 ENGAGEMENT 的推广系列中promotion_type 为LIVE_SHOPPING 的广告组。
	// 不支持复制推广目标为 ENGAGEMENT 的推广系列中optimization_goal 为PAGE_VISIT 的广告组。
	// 不支持复制 Smart+ 推广系列。
	// 您想要复制的推广系列需包含至少一个未删除的广告组，且其中包含至少一个未删除的广告。
	// 您想要复制的推广系列最多包含 20 个未删除的广告组，每个广告组中最多包含 50 个未删除的广告。
	// 不支持复制推广系列中的程序化创意广告组，但支持复制智能创意广告组。
	// 若您想要复制的推广系列中包含已与拆分对比测试绑定的广告组，则将仅复制广告组，而不会复制与广告组绑定的拆分对比测试组。
	CampaignID string `json:"campaign_id,omitempty"`
	// RequestID 请求 ID。该 ID 支持接口幂等，避免重复请求。
	// 若您在 10 秒的缓存时间内传入相同的请求 ID 重试多次请求，只有一次请求会成功。若您在缓存时间已过后，发送带有已过期的请求 ID 的重复请求，服务器将该请求视作新请求进行处理。
	// 该 ID 和返回参数中的 request_id 不同。返回的request_id用于唯一标识一次 HTTP 请求。
	// 本字段的值需为字符串格式的 64 位整数值。
	RequestID string `json:"request_id,omitempty"`
	// OperationStatus 推广系列的操作状态。
	// ENABLE：推广系列处于启用（“开”）状态。
	// DISABLE：推广系列处于未启用（“关”）状态。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// VirtualObjectiveType 新推广目标类型。
	// 枚举值：SALES（销量）。
	// 使用本字段创建推广目标为销量的推广系列，该目标合并了网站转化量目标和商品销量目标。了解销量目标的详情，参见销量目标。
	VirtetualObjectiveType enum.VirtualObjectiveType `json:"virtual_objective_type,omitempty"`
	// SalesDestination 销售目标页面，即想要推动销售的渠道。
	// 枚举值：
	// TIKTOK_SHOP：TikTok Shop。推动 TikTok Shop 上的销售。
	// WEBSITE：网站。推动网站上的销售。
	// APP：应用。推动应用上的销售（需要商品库）。
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// CampaignName 新推广系列的名称。
	// 长度限制：512 字符，且不能包含 emoji。
	// 中文或日文每个字占用 2 个字符，每个英语字母占用 1 个字符。
	// 若未指定本字段，本字段将默认设置为"COPIED_{{原推广系列名称} }"。例如，若原推广系列名称为"FIRST_CAMPAIGN"，且未指定本字段，则新推广系列的名称将默认设置为"COPIED_FIRST_CAMPAIGN"。
	CampaignName string `json:"campaign_name,omitempty"`
	// IsAdvancedDedicatedCampaign 该推广系列是否为高级专属推广系列。高级专属推广系列采用捕捉实时信号的高级投放模型。
	// 枚举值：true，false。
	IsAdvancedDedicatedCampaign bool `json:"is_advanced_dedicated_campaign,omitempty"`
	// Budget 推广系列预算
	// 若未指定本字段，本字段将默认设置为原推广系列的预算。
	// 若想了解最低预算以及如何设置预算类型，请参阅预算。若想获取某一币种对应的每日预算取值范围，请参阅币种-每日预算取值范围。
	Budget float64 `json:"budget,omitempty"`
	// RtaID 实时 API ID，即实时 API 策略标识符
	// 若想获取与您的广告账户绑定的实时 API ID，请联系您的 TikTok 销售代表。
	// 若想了解为广告启用实时 API 的前提及启用步骤，请查看实时 API。
	// 注意：为广告启用实时 API 目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。您需首先完成实时 API 的入门流程，具备接收发送至您的系统的实时 API 请求并进行回复的能力。
	RtaID string `json:"rta_id,omitempty"`
	// PONumber PO（采购订单）号。
	// PO 号可用于结算单追踪和对账。
	// 了解月度结算单上的 PO 号。
	PONumber string `json:"po_number,omitempty"`
	// ScheduleType 本字段用于为新推广系列中的所有广告组设置统一的排期。
	// 若想为新推广系列中的每个广告组设置单独的排期，您可传入 adgroup_list 中的schedule_type 字段。
	// 所有新广告组的排期类型。
	// 枚举值：
	// SCHEDULE_START_END：设置投放广告组的开始时间和结束时间。您需同时传入 schedule_start_time 和 schedule_end_time 。
	// SCHEDULE_FROM_NOW：设置持续投放广告组的开始时间。您只需要设置一个开始时间，结束时间将自动设置为开始时间十年后的时间。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 传入 schedule_type 时必填。
	// 所有新广告组的投放起始时间（UTC+0），格式为"YYYY-MM-DD HH:MM:SS"。
	// 起始时间可以早于当前时间，但是最多只能早12个小时，且不可晚于"2028-01-01 00:00:00"。
	ScheduleStartTime string `json:"schedule_start_time,omitempty"`
	// ScheduleEndTime schedule_type 设置为 SCHEDULE_START_END 时必填。
	// schedule_type 设置为 SCHEDULE_FROM_NOW 时，不支持本字段。
	// 所有新广告组的投放结束时间（UTC+0），格式为"YYYY-MM-DD HH:MM:SS"。
	// 结束时间不可晚于"2038-01-01 00:00:00"。
	ScheduleEndTIme string `json:"schedule_end_time,omitempty"`
	// DeepCopyMode 复制模式，决定新推广系列中广告组和广告的创建方式。
	// 枚举值：
	// DEFAULT：默认复制模式。该模式将从原推广系列复制所有未删除的广告组和广告到新生成的推广系列中。将忽略 adgroup_list 。
	// CUSTOM： 自定义复制模式。该模式仅从原推广系列复制所指定的未删除的广告组和广告到新生成的推广系列中。
	// 在此模式下，您需同时传入 adgroup_list 。最多支持复制 20 个广告组，每个个广告组中最多复制 50 个广告。
	// 默认值： DEFAULT。
	DeepCopyMode string `json:"deep_copy_mode,omitempty"`
	// AdgroupList deep_copy_mode 设置为 DEFAULT 或不传入时， 忽略本字段。
	// deep_copy_mode 设置为 CUSTOM 时，本字段必填。
	// 为新推广系列中的广告组自定义的设置。
	// 最大数量：20。
	AdgroupList []Adgroup `json:"adgroup_list,omitempty"`
}

type Adgroup struct {
	// AdgroupID 传入 adgroup_list 时必填。
	// 您想要复制的广告组。该广告组需属于原推广系列（campaign_id）。
	// 若广告组未启用程序化创意且未启用智能创意，您需传入 adgroup_id 对象数组中的ad_list 字段 。
	// 若广告组为智能创意广告组，您无需传入 adgroup_id 对象数组中的ad_list 字段 ，但可传入可选参数 smart_creative_info，自定义在新广告组中生成智能创意广告所基于的广告创意。
	// 注意:
	// 自 2025 年 7 月 19 日起，您将无法创建或复制使用子版位 TikTok Lite （TIKTOK_LITE）的广告组。为确保流畅、不受影响的广告创建体验，推荐您从广告组创建流程中移除此子版位或避免复制包含此类广告组的推广系列。
	// 自 2024 年 11 月 30 日起，您将无法创建或复制优化目标为“安装与应用内事件数据”且版位为仅 TikTok 或自动版位的广告组。此变动将影响使用以下配置的广告组：
	// 推广系列层级：objective_type 为 APP_PROMOTION
	// 广告组层级：
	// optimization_goal 为 INSTALL 且指定了 secondary_optimization_event 合法值
	// placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 为["PLACEMENT_TIKTOK"]，或 placement_type 为PLACEMENT_TYPE_AUTOMATIC
	// 现有的投放在 TikTok 版位的“安装与应用内事件数据”广告组将不受影响。此外，Pangle 版位和全球化应用组合版位（即 placement_type 为 PLACEMENT_TYPE_NORMAL 且 placements 的值仅包含PLACEMENT_PANGLE 和 PLACEMENT_GLOBAL_APP_BUNDLE）也不受影响。请注意此变更，并在必要时对您的集成进行适当的调整。
	// 使用未在 MMP 侧启用 SAN 模块的 App 的广告组不支持复制。若想确保您的 App 启用 TikTok 自归因集成，可查看如何将新应用集成到 SAN 和如何将现有应用迁移到 SAN。
	AdgroupID string `json:"adgroup_id,omitempty"`
	// OperationStatus 新广告组创建时的启用或关闭状态。
	// 枚举值：
	// ENABLE ：广告组创建时处于启用状态。
	// DISABLE ：广告组创建时处于关闭状态。
	// 默认值：ENABLE。
	// 若想在广告组创建后更新其启用或关闭状态，请使用/adgroup/status/update/。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// AdgroupName 新广告组的名称。
	// 长度限制：512 字符，且不能包含 emoji。
	// 中文或日文每个字占用 2 个字符，每个英语字母占用 1 个字符。
	// 若未指定本字段，本字段将默认设置为原广告组的名称。
	AdgroupName string `json:"adgroup_name,omitempty"`
	// AutomatedKeywordsEnabled 仅在满足推广系列层级 is_search_campaign 为 true 时生效。
	// 是否启用自动关键词，即在添加广告后，让系统自动生成关键词。这样可以拓展优质流量，以便提升成效。在报表中查看成效出色的自动关键词。
	// 支持的值：true，false。
	// 默认值：false。
	// 注意:
	// 搜索广告推广系列中启用自动关键词目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	AutomatedKeywordsEnabled bool `json:"automated_keywords_enabled,omitempty"`
	// LocationIDs 新广告组定向的地域 ID 列表。
	// 最大数量：3,000。若同时传入 location_ids 和 zipcode_ids，则单个广告组中地域 ID 和邮政编码 ID 的数量之和不可超过 3,000。
	// 若想根据您的版位和推广目标来查询您可以投放的地区以及相应 ID，可以使用/tool/region/或/tool/targeting/search/接口。ID列表也可以通过附录-地域定向查看。
	// 注意：
	// 定向地域不允许重叠。例如，不支持同时定向美国和加利福尼亚州。
	// 若您在创建广告组时通过 location_ids 或 zipcode_ids 定向了美国地域，您可随后将定向地域更新为美国的其他地域，但不允许从定向地域中完全去除美国地域，从而只定向非美国地域。
	// 若您要复制的广告组的 tiktok_subplacements 值中包含 TIKTOK_LITE，则不允许从新广告组的定向地域（location_ids）中完全去除日本和韩国地域。
	LocationIDs []string `json:"location_ids,omitempty"`
	// ZipcodeIDs 新广告组定向的地域的邮政编码 ID。
	// 最大数量：3,000。若同时传入 location_ids 和 zipcode_ids，则单个广告组中地域 ID 和邮政编码 ID 的数量之和不可超过 3,000。
	// 您可以通过/tool/targeting/search/ 返回的geo_id（对应的geo_type 为ZIP_CODE）获取可投地域的邮政编码 ID，返回结果基于您设置的版位，推广目标和关键词。
	// 注意：
	// 邮政编码定向目前仅支持定向美国、加拿大、巴西、印度尼西亚、泰国、越南的地域。
	// 巴西、印度尼西亚、泰国、越南的邮政编码定向目前为白名单功能。如需使用此功能，请联系您的 TikTok 销售代表。
	// 不支持在启用了特殊广告分类（special_industries）的推广系列中使用邮政编码定向。
	// 不支持在 objective_type 为 RF_REACH 的推广系列中使用邮政编码定向。
	// 邮政编码定向仅支持 TikTok 版位。因此，placements 的值中不包含 PLACEMENT_TIKTOK。
	// 定向地域不允许重叠。例如，不支持同时定向美国和加利福尼亚州。
	// 若您在创建广告组时通过 location_ids 或 zipcode_ids 定向了美国地域，您可随后将定向地域更新为美国的其他地域，但不允许从定向地域中完全去除美国地域，从而只定向非美国地域。
	// 欲获取邮政编码 ID 的信息，您可使用/tool/targeting/info/。
	ZipcodeIDs []string `json:"zipcode_ids,omitempty"`
	// AudienceIDs 新广告组定向的受众 ID 列表。
	AudienceIDs []string `json:"audience_ids,omitempty"`
	// ExcludedAudienceIDs 新广告组的排除受众 ID 列表
	ExcludedAudienceIDs []string `json:"excluded_audience_ids,omitempty"`
	// AgeGroups 新广告组定向的受众年龄区间。
	// 枚举值详见枚举值-受众年龄区间。
	// 在某些场景下，不允许创建定向美国、拉丁美洲、欧洲经济区、英国、瑞士或加拿大 13-17 年龄区间（AGE_13_17）受众的广告组。若想了解详情，请查看 TikTok 广告的新增年龄定向限制。
	// 若您使用了与美国、拉丁美洲、欧洲经济区、英国、瑞士或加拿大 13-17 年龄区间定向不兼容的定向设置，且将 age_groups 设置为 [] 或不传入字段 age_groups ，则该字段将默认为["AGE_18_24", "AGE_25_34", "AGE_35_44", "AGE_45_54", "AGE_55_100"]。
	// 注意：若您通常将 age_groups 设置为 [] 或不传入字段 age_groups，推荐您在该字段中明确指定您想要定向的具体年龄区间，例如AGE_18_24。
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// Budget 广告组预算。
	// 当开启了推广系列预算优化(budget_optimize_on)时，返回0.0。
	// 对于 TopView 广告，本字段代表预估折后预算。
	Budget float64 `json:"budget,omitempty"`
	// ScheduleType 广告投放时间类型。
	// 枚举值: SCHEDULE_FROM_NOW，SCHEDULE_START_END。如果您选择了SCHEDULE_START_END，您需要明确投放开始和结束时间。 如果您选择了SCHEDULE_FROM_NOW，您只需要明确投放开始时间。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime string `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime string `json:"schedule_end_time,omitzero"`
	// AdList deep_copy_mode 设置为 DEFAULT 或不传入时，忽略本字段。
	// deep_copy_mode 设置为 CUSTOM，且原广告组（adgroup_id）未启用程序化创意且未启用智能创意时，本字段必填。
	// 若不想自定义新生成的广告组中广告的设置，可在本对象数组中仅传入 ad_id 字段。
	// 若想自定义新生成广告组中广告的设置，可在本对象数组中同时传入 ad_id 字段和其他设置字段。
	// 为新生成的未启用程序化创意且未启用智能创意的广告组中的广告自定义的设置。
	// 最大数量：50。
	AdList []Ad `json:"ad_list,omitempty"`
}

type Ad struct {
	// AdID 广告 ID
	AdID string `json:"ad_id,omitempty"`
	// OperationStatus 广告的操作状态。
	// ENABLE：广告处于启用（“开”）状态。
	// DISABLE：广告处于未启用（“关”）状态。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
	// IdentityType 认证身份类型。枚举值： CUSTOMIZED_USER (自定义用户），AUTH_CODE (授权用户)， TT_USER (TikTok用户）， BC_AUTH_TT（添加到商务中心的TikTok用户）。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID 认证身份ID。
	// 如果您不使用Spark Ads，传入 identity_id 和 identity_type (CUSTOMIZED_USER类型)可帮助您更好地管理广告信息。
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityAuthorizedBcID 当identity_type 为 BC_AUTH_TT时返回。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
	// AdFormat 广告样式。枚举值：SINGLE_IMAGE， SINGLE_VIDEO， LIVE_CONTENT, CAROUSEL_ADS（非视频购物类型的轮播广告），CATALOG_CAROUSEL（视频购物类型的轮播广告） 。
	AdFormat enum.AdFormat `json:"ad_format,omitempty"`
	// VideoID 视频 ID。
	// 若广告为通过从关联的企业号、创作者授权的帖子或商务中心中创作者授权的账户提取原生视频而创建的 Spark Ads，则本字段值为 null。此时您可以通过tiktok_item_id字段获取用作 Spark Ads 的 TikTok 帖子 ID。
	// 若想了解 Spark Ads 详情，请查看创建 Spark Ads。
	VideoID string `json:"video_id,omitempty"`
	// ImageIDs 图片 ID 列表
	ImageIDs []string `json:"image_ids,omitempty"`
	// MusicID TikTok轮播广告中使用的音乐的ID
	MusicID string `json:"music_id,omitempty"`
	// TiktokItemID 用作 Spark Ads 的 TikTok 帖子 ID，
	// 若想了解 Spark Ads 详情，请查看创建 Spark Ads。
	// 注意：本字段为以下类型的 Spark Ads 返回：
	// 通过 Spark Ads 提取创建的 Spark Ads。
	// 通过 Spark Ads 推送创建且通过广告审核的 Spark Ads。
	TiktokItemID string `json:"tiktok_item_id,omitempty"`
	// AdText 广告文案，将作为广告创意的一部分展示给您的受众，向他们传达您想要推广的信息。关键词匹配类型为精确匹配
	AdText string `json:"ad_text,omitempty"`
	// AigcDisclosureType 是否启用 AI 生成内容自主声明开关，以表明广告中包含 AI 生成内容。启用该开关后，当用户查看完整广告时，您的广告将带有“广告主标记为 AI 生成”的标签。
	// 枚举值：
	// SELF_DISCLOSURE：启用开关，声明广告包含 AI 生成内容。
	// NOT_DECLARED：不声明广告包含 AI 生成内容。
	AigcDisclosureType enum.AigcDisclosureType `json:"aigc_disclosure_type,omitempty"`
	// CreativeAutoEnhancementStrategyList 应用于你的广告的自动优化策略列表。
	// 你传入的优化功能将在你的推广系列投放期间自动实时应用。
	// 枚举值：
	// VIDEO_QUALITY：视频质量。通过提高分辨率和清晰度来提升整体视觉效果。
	// MUSIC_REFRESH：音乐焕新。替换为当前 TikTok 上流行的音乐，紧跟潮流趋势。
	// IMAGE_QUALITY：图片质量。通过提高分辨率和清晰度来提升整体视觉效果。
	// IMAGE_RESIZE：调整尺寸。调整图片大小，充分利用全屏功能。
	CreativeAutoEnhancementStrategyList []enum.CreativeAutoEnhancementStrategy `json:"creative_auto_enhancement_strategy_list,omitempty"`
	// TrackingPixelID 正在监测的 Pixel ID。你可以使用此字段追踪发生在外部网站的事件
	TrackingPixelID string `json:"tracking_pixel_id,omitempty"`
	// TrackingAppID 正在监测的应用的 ID
	TrackingAppID string `json:"tracking_app_id,omitempty"`
	// SmartCreativeInfo deep_copy_mode 设置为 DEFAULT 或未传入时，忽略本字段。
	// deep_copy_mode 设置为 CUSTOM 且原广告组（adgroup_id）为智能创意广告组时，本字段可选。
	// 若您希望新广告组中生成智能创意广告所基于的广告创意与原广告组相同，则无需传入本字段。
	// 若您希望自定义新广告组中生成智能创意广告所基于的广告创意，则需传入本字段。
	// 新智能创意广告组中生成智能创意广告所基于的广告创意。
	// 注意：智能创意仅支持objective_type （推广目标）设置为 APP_PROMOTION，WEB_CONVERSIONS，TRAFFIC 或 LEAD_GENERATION 的推广系列 。
	SmartCreativeInfo *aco.Ad `json:"smart_creative_info,omitempty"`
}

// Encode implements PostEncode
func (r *CopyTaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CopyTaskCreateResponse 创建推广系列异步复制任务 API Response
type CopyTaskCreateResponse struct {
	model.BaseResponse
	Data *CopyTaskCreateResult `json:"data"`
}

type CopyTaskCreateResult struct {
	// TaskID 推广系列异步复制任务 ID。
	// 若想获取任务的结果，需使用 /campaign/copy/task/check/。
	// 若发生报错（adgroup_error_list 不为空），则不会生成任务 ID，本字段的值将为空字符串（""）。
	TaskID string `json:"task_id,omitempty"`
	// AdgroupErrorList 广告组复制过程中发生的报错。
	// 若无报错，此字段的值将为空列表（[]）
	AdgroupErrorList []AdgroupError `json:"adgroup_error_list,omitempty"`
}

type AdgroupError struct {
	// AdgroupID 复制失败的原广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// ErrorMessage 复制该广告组（adgroup_id）时发生的报错
	ErrorMessage string `json:"error_message,omitempty"`
}
