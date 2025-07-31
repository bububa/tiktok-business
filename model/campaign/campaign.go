package campaign

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Campaign 推广系列
type Campaign struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignSystemOrigin 仅为内容加热推广系列返回本字段。
	// 非内容加热推广系列不返回本字段。
	//
	// 推广系列来源。
	//
	// 枚举值：
	// PROMOTE：该推广系列为通过 TikTok 移动应用创建的内容加热推广系列。
	//
	// 仅为内容加热推广系列（campaign_system_origin 为 PROMOTE）返回以下信息：
	// advertiser_id
	// campaign_id
	// create_time
	// campaign_name
	// operation_status
	// secondary_status
	CampaignSystemOrigin enum.CampaignSystemOrigin `json:"campaign_system_origin,omitempty"`
	// IsSearchCampaign 推广系列是否为搜索推广系列。
	// 枚举值：true，false。
	// 注意：目前不支持使用 API 接口创建搜索推广系列。
	IsSearchCampaign bool `json:"is_search_campaign,omitempty"`
	// IsSmartPerformanceCampaign 是否为自动化类型的推广系列。
	// 支持的值：
	// true：推广系列为 Smart+ 推广系列或智能效果网站推广系列。
	// false：推广系列为普通类型的推广系列。
	// 注意：若is_smart_performance_campaign 为 true 且 objective_type 为 WEB_CONVERSIONS，您可根据/campaign/spc/get/返回的spc_type查看该网站推广系列为 Smart+ 网站推广系列还是智能效果网站推广系列。
	IsSmartPerformanceCampaign bool `json:"is_smart_performance_campaign,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignType 推广系列类型。枚举值: REGULAR_CAMPAIGN（普通推广系列）、IOS14_CAMPAIGN（iOS 14专属推广系列）。
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// PostbackWindowMode 选定的模式决定您想要确保能接收到的 SKAN (SKAdNetwork) 4.0 回传。
	// 枚举值：
	// POSTBACK_WINDOW_MODE1：此模式确保能接收到第一次回传，对应的归因窗口期为 0-2 天。回传数据最长需要 4 天时间接收，推广系列配额释放的时间最长也为 4 天。
	// POSTBACK_WINDOW_MODE2：此模式确保能接收到第一次及第二次回传，对应的归因窗口期为 3-7 天。回传数据最长需要 13 天时间接收，推广系列配额释放的时间最长也为 13 天。
	// POSTBACK_WINDOW_MODE3：此模式确保能接收到三次回传，对应的归因窗口期为 8-35 天。回传数据最长需要 41 天时间接收，推广系列配额释放的时间最长也为 41 天。
	//
	// 注意：若您已为应用启用 Adjust，Airbridge，Appsflyer，Branch，Kochava，或 Singular 移动合作伙伴 (MMP) 跟踪，且您的 MMP SDK 版本已更新至支持 SKAN 4.0 的版本，您可在事件管理平台将应用迁移至 SKAN 4.0 。若想了解为应用启用 MMP 跟踪的步骤，请查看如何在 TikTok 广告管理平台中设置应用归因。若想了解将应用迁移至 SKAN 4.0 的步骤，请查看关于 SKAN 4.0 和 TikTok 以及如何迁移到 SKAN 4.0。
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
	// AppID 推广的App的ID。您可通过/app/list/获取app_id。
	AppID string `json:"app_id,omitempty"`
	// Budget 推广系列预算
	Budget model.Float64 `json:"budget,omitempty"`
	// BudgetMode 预算类型。
	// 枚举值：
	// BUDGET_MODE_INFINITE：不限预算。
	// BUDGET_MODE_TOTAL：总预算。
	// BUDGET_MODE_DAY：日预算。
	// BUDGET_MODE_DYNAMIC_DAILY_BUDGET：动态日预算，即一周七日的平均预算。单日实际消耗不超过设定的七日平均预算的125%。周消耗不会超过七日平均预算*7。
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// RtaID 实时 API ID，即实时 API 策略标识符
	RtaID string `json:"rta_id,omitempty"`
	// RtaBidEnabled 是否使用实时 API 出价。
	// 支持的值：
	// true：启用。
	// false：不启用。
	RtaBidEnabled bool `json:"rta_bid_enabled,omitempty"`
	// RtaProductSelectionEnabled 是否使用实时 API 自动选择商品。
	// 枚举值：
	// true：使用。
	// false：不使用。
	RtaProductSelectionEnabled bool `json:"rta_product_selection_enabled,omitempty"`
	// SecondaryStatus 推广系列状态（二级状态）。枚举值详见枚举值-二级状态-推广系列状态。
	// te 注意：沙箱环境下本字段不返回，因为推广系列未实际投放。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// OperationStatus 推广系列的操作状态。
	// ENABLE：推广系列处于启用（“开”）状态。
	// DISABLE：推广系列处于未启用（“关”）状态。
	OpeterationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// Objective 推广类型（应用或落地页），枚举值: APP(应用)，LANDING_PAGE(落地页) 。
	Objective enum.Objective `json:"objective,omitempty"`
	// ObjectiveType 推广目标，获取枚举值，参阅 枚举值-推广目标。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// AppPromotionType 应用推广类型。枚举值：
	// APP_INSTALL：应用安装。让新用户安装您的应用。
	// APP_RETARGETING：应用再营销。重新吸引现有用户在您的应用中执行操作。
	// APP_PREREGISTRATION：应用预注册。让新用户在您的应用发布前完成预注册。
	// APP_POSTS_PROMOTION：应用帖子推广。推广 TikTok 帖子，提升应用认知度并度量转化量。
	// 注意：目前不支持通过 API 创建应用帖子推广类型的推广系列，仅支持通过 API 更新或检索此类推广系列。
	AppPromotionType enum.AppPromotionType `json:"aepp_promotion_type,omitempty"`
	// VirtualObjectiveType 新推广目标类型。

	// 枚举值：SALES（销量）。
	//
	// 使用本字段创建推广目标为销量的推广系列，该目标合并了网站转化量目标和商品销量目标。了解销量目标的详情，参见销量目标。
	VirtetualObjectiveType enum.VirtualObjectiveType `json:"virtual_objective_type,omitempty"`
	// SalesDestination 销售目标页面，即想要推动销售的渠道。
	// 枚举值：
	// TIKTOK_SHOP：TikTok Shop。推动 TikTok Shop 上的销售。
	// WEBSITE：网站。推动网站上的销售。
	// APP：应用。推动应用上的销售（需要商品库）。
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// CampaignProductSource 仅当推广系列对应的objective_type为PRODUCT_SALES，且设置了商品来源时返回。
	// 推广系列的商品来源。
	// 枚举值：
	// CATALOG ：商品库。
	// STORE：TikTok Shop 店铺，或 TikTok 橱窗。
	CampaignProductSource enum.ProductSource `json:"campaign_product_source,omitempty"`
	// BudgetOptimizeOn 仅为开启了推广系列预算优化的推广系列返回本字段。
	// 是否开启推广系列预算优化。
	// 枚举值：true（开启）。
	// 详见推广系列预算优化。
	// 注意：对于 Smart+ 推广系列或智能效果网站推广系列，本字段值将为 true，因为这些推广系列均默认启用推广系列预算优化。
	BudgetOptimizeOn bool `json:"budget_optimize_on,omitempty"`
	// BidType 推广系列层级的竞价策略。当开启推广系列预算优化的时候返回
	BidType enum.BidType `json:"bid_type,omitempty"`
	// DeepBidType 深度事件出价类型。获取枚举值，请参阅枚举值-深度事件出价类型
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// RoasBid 用于价值优化的ROAS目标值。
	RoasBid model.Float64 `json:"roas_bid,omitempty"`
	// OptimizationGoal 优化目标。当开启推广系列预算优化的时候返回
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// IsNewStructure 该推广系列是否为新结构（推广系列与其下广告组及广告结构相同）te。
	IsNewStructure bool `json:"is_new_structure,omitempty"`
	// CreateTime 推广系列创建时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 推广系列修改时te间
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
	// IsAdvancedDedicatedCampaign 该推广系列是否为高级专属推广系列。高级专属推广系列采用捕捉实时信号的高级投放模型。
	// 枚举值：true，false。
	IsAdvancedDedicatedCampaign bool `json:"is_advanced_dedicated_campaign,omitempty"`
	// DisableSkanCampaign 是否关闭 SKAN（SKAdNetwork）归因（Apple 针对 iOS 推广系列的转化归因解决方案）。
	// 枚举值：
	// true：关闭 SKAN 归因。推广系列将不受专属推广系列配额限制的约束，您可为此推广系列拉取自归因平台（SAN）指标数据。但是，不支持为此推广系列拉取 SKAN 指标数据。了解自归因平台集成的详情。
	// false：启用 SKAN 归因。推广系列将受专属推广系列配额限制的约束，您可为此推广系列拉取 SKAN 指标数据。
	DisableSkanCampaign bool `json:"disable_skan_campaign,omitempty"`
	// BidAlignType 专属推广系列的归因类型。此类型决定广告组层级设置的目标 CPA （conversion_bid_price） 或目标 ROAS（roas_bid） 基于的归因网络类型。
	// 枚举值：
	// SAN：自归因平台归因。
	// SKAN：SKAN 归因。
	BidAlignType enum.BidAlignType `json:"bid_atlign_type,omitempty"`
	// CampaignAppProfilePageState 下载中间页使用情况。
	// 枚举值：INVALID，UNSET，ON，OFF。
	CampaignAppProfilePageState enum.AppProfilePageState `json:"campaign_app_profile_page_state,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：
	// 在推广系列商品来源为商品库和推广系列商品来源非商品库的购物广告中使用特殊广告分类目前为单独的白名单功能。如需使用这些功能，请联系您的TikTok销售代表。
	// 本字段对于注册地为美国或加拿大的广告主已全量。若您是注册地非美国或加拿大的广告主且想在定向美国或加拿大的广告中使用特殊广告分类，您需申请额外的白名单。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// RfCampaignType 仅当推广目标（objective_type）设置为RF_REACH时，可根据本字段判断具体的合约推广系列类型。
	// 枚举值：
	// STANDARD：覆盖和频次推广系列。
	// PULSE：TikTok Pulse 推广系列。
	// TOPVIEW： TopView 推广系列。
	RfCampaignType enum.RfCampaignType `json:"rf_campaign_type,omitempty"`
}
