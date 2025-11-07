package gmvmax

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
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// OperationStatus 推广系列的操作状态。
	// ENABLE：推广系列处于启用（“开”）状态。
	// DISABLE：推广系列处于未启用（“关”）状态。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// ObjectiveType 推广目标，获取枚举值，参阅 枚举值-推广目标。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// SecondaryStatus 推广系列状态（二级状态）。枚举值详见枚举值-二级状态-推广系列状态。
	// te 注意：沙箱环境下本字段不返回，因为推广系列未实际投放。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// StoreID TikTok Shop ID
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 有权限访问该 TikTok Shop（store_id）的商务中心的 ID。
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// ShoppingAdsType GMV Max 推广系列类型。
	// 枚举值：
	// PRODUCT：商品 GMV Max 推广系列。
	// LIVE：直播 GMV Max 推广系列。
	ShoppingAdsType enum.ShoppingAdsType `json:"shopping_ads_type,omitempty"`
	// ProductSpecificType 选择商品的维度。
	// 枚举值：
	// ALL：允许 TikTok 从 TikTok Shop 所有商品中动态选择。
	// CUSTOMIZED_PRODUCTS：选择 TikTok Shop 中自定义数量的特定商品。
	// UNSET：未设置。
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ItemGroupIDs 仅当 shopping_ads_type 为 PRODUCT 且 product_specific_type 为CUSTOMIZED_PRODUCTS时返回本字段。
	// 商品 SPU ID 列表。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// OptimizationGoal 优化目标。
	// 枚举值：
	// VALUE：总收入。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// ROIProtectionEnabled 该推广系列是否符合 ROI（投资回报率）保障资格。
	// 投资回报率保障功能在您遵循 GMV Max 指南和最佳实践且推广系列的投资回报率仍低于特定阈值的情况下，自动向您发放广告赠款。若想了解 ROI 保障功能的符合资格及不符合的场景，请查看关于 GMV Max 推广系列的 ROI 保障功能。
	// 支持的值： true， false。
	ROIProtectionEnabled bool `json:"roi_protection_enabled,omitempty"`
	// ROIProtectionCompensationStatus 该推广系列的 ROI（投资回报率）保障状态。
	// 投资回报率保障功能在您遵循 GMV Max 指南和最佳实践且推广系列的投资回报率仍低于特定阈值的情况下，自动向您发放广告赠款。若想了解 ROI 保障功能的符合资格及不符合的场景，请查看关于 GMV Max 推广系列的 ROI 保障功能。
	// 枚举值：
	// IN_EFFECT：符合 ROI 保障资格。如果您的广告在 24 小时内促成 20 次以上转化，但 ROI 结果低于目标的 90%，此推广系列有资格获得广告费用赠款。为确保推广系列持续符合 ROI 保障资格，请勿修改 ROI、暂停广告投放或启用最大投放量优化模式。推广系列资格每 24 小时重置一次。
	// NOT_ELIGIBLE：不符合 ROI 保障资格。如果因为修改了 ROI、暂停了推广系列、启用了最大投放量优化模式，或者您的店铺、广告账号、直播或商品存在问题，导致没有达成目标 ROI，那么此推广系列不符合广告费用赠款资格。推广系列资格每 24 小时重置一次。
	ROIProtectionCompensationStatus enum.ROIProtectionCompensationStatus `json:"roi_protection_compensation_status,omitempty"`
	// DeepBidType 出价策略。
	// 枚举值：
	// VO_MIN_ROAS: 最低 ROAS。在有 ROI 目标的条件下出价。
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// RoasBid 仅当 deep_bid_type 为 VO_MIN_ROAS 时返回本字段。
	// ROI 目标。
	RoasBid model.Float64 `json:"roas_bid,omitempty"`
	// Budget 日预算
	Budget model.Float64 `json:"budget,omitempty"`
	// ScheduleType 排期类型。
	// 枚举值：
	// SCHEDULE_FROM_NOW：在排定的开始时间之后持续投放推广系列。
	// SCHEDULE_START_END:：在排定的开始时间和结束时间之间投放推广系列。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime model.DateTime `json:"schedule_start_time,omitzero"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime model.DateTime `json:"schedule_end_time,omitzero"`
	// Placements 广告位（版位），即推广系列投放的应用。
	// 枚举值：
	// PLACEMENT_TIKTOK: TikTok。
	// PLACEMENT_PANGLE: Pangle。
	// 注意：
	// 商品 GMV Max 推广系列自动通投至 TikTok 和 Pangle 版位以进行充分流量探索。了解 Pangle 版位详情。
	// 若想为您的商品 GMV Max 推广系列退出 Pangle 版位，可联系您的 TikTok 代表。
	Placements []enum.Placement `json:"placements,omitempty"`
	// LocationIDs 定向地域 ID。使用tool/region 获取定向地域ID
	LocationIDs []string `json:"location_ids,omitempty"`
	// AgeGroups 定向受众年龄
	// 枚举值：详见枚举值-受众年龄区间
	AgeGroups []enum.Age `json:"age_groups,omitempty"`
	// ProductVideoSpecificType 视频选择模式。
	// 枚举值：
	// AUTO_SELECTION：自动选择视频。
	// CUSTOM_SELECTION：手动选择视频。
	// UNSET：未设置。
	ProductVideoSpecificType enum.ProductVideoSpecificType `json:"product_video_specific_type,omitempty"`
	// IdentityList 与该 GMV Max 推广系列绑定的认证身份（即 TikTok 账号）列表
	IdentityList []Identity `json:"identity_list,omitempty"`
	// AffiliatePostsEnabled 是否为商品 GMV Max 推广系列启用联盟作品。
	// 联盟帖子是由联盟方（即参加 TikTok Shop 联盟项目的达人）制作并被授权用于 TikTok 店铺广告的 TikTok 帖子。了解 TikTok 店铺广告的联盟创意的详情。
	// 支持的值： true， false
	AffiliatePostsEnabled bool `json:"affiliate_posts_enabled,omitempty"`
	// Items 与该 GMV Max 推广系列绑定的已授权 TikTok 作品（即 TikTok 帖子）或自定义的作品。
	// 已授权的作品代表来自您的 TikTok 账号或通过视频代码授权您使用，能够充分展示你所选商品的现有作品。
	// 自定义的作品代表创建推广系列时通过向所指定视频添加自定义商品链接而新建的帖子。这些帖子仅用于此商品 GMV Max 推广系列
	Items []TikTokItem `json:"items,omitempty"`
	// CampaignCustomAnchorVideoID 此商品 GMV Max 推广系列中创建的自定义的作品的合集 ID。
	// 您可将本字段传入 /gmv_max/custom_anchor_video_list/get/ 中从而获取合集中的自定义的作品的详情
	CampaignCustomAnchorVideoID string `json:"campaign_custom_anchor_video_id,omitempty"`
	// CustomAnchorVideoList 与该 GMV Max 推广系列绑定的自定义 TikTok 作品（即 TikTok 帖子）。
	// 自定义的作品代表带有自定义商品链接的作品
	CustomAnchorVideoList []AnchorVideo `json:"custom_anchor_video_list,omitempty"`
	// CreateTime 推广系列创建时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 推广系列修改时te间
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
}
