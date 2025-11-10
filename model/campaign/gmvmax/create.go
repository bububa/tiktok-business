package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/identity"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建 GMV Max 推广系列 API Request
type CreateRequest struct {
	// CampaignName GMV Max 推广系列的名称
	CampaignName string `json:"campaign_name,omitempty"`
	// RequestID 请求 ID。该 ID 支持接口幂等，避免重复请求。
	// 若您在 10 秒的缓存时间内 传入相同的请求 ID 重试多次请求，只有一次请求会成功。若您在缓存时间已过后，发送带有已过期的请求 ID 的重复请求，服务器将该请求视作新请求进行处理。
	// 该 ID 和返回参数中的 request_id 不同。返回的 request_id 用于唯一标识一次 HTTP 请求。
	// 本字段的值需为字符串格式的 64 位整数值。
	RequestID string `json:"request_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 有权限访问该 TikTok Shop（store_id）的商务中心的 ID
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// AdvertiserID 广告主 ID。
	// 若想确认广告账号拥有您想要推广的 TikTok Shop 的 GMV Max 专属授权，可调用/gmv_max/store/list/，并确认返回的 is_gmv_max_available 为 true 且返回的 exclusive_authorized_advertiser_info 中的 advertiser_id 与您想要用来创建 GMV Max 推广系列的广告账户一致。
	// 若不一致，则可使用 /gmv_max/exclusive_authorization/create/授予此广告账户该 TikTok Shop 的专属 GMV Max 授权。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ShoppingAdsType GMV Max 推广系列类型。
	// 枚举值：
	// PRODUCT：商品 GMV Max 推广系列。
	// LIVE：直播 GMV Max 推广系列。
	// 若想了解如何创建此类推广系列，可查看创建商品 GMV Max 推广系列和创建直播 GMV Max 推广系列。
	ShoppingAdsType enum.ShoppingAdsType `json:"shopping_ads_type,omitempty"`
	// ProductSpecificType 仅当 shopping_ads_type 为 PRODUCT 时生效。
	// 选择商品的维度。
	// 枚举值：
	// ALL：允许 TikTok 从 TikTok Shop 所有商品中动态选择。
	// CUSTOMIZED_PRODUCTS：选择 TikTok Shop 中自定义数量的特定商品。
	// 若想确认您可以推广 TikTok Shop 的所有商品，可调用 /gmv_max/store/shop_ad_usage_check/ 并查看返回的 promote_all_products_allowed
	ProductSpecificType enum.ProductSpecificType `json:"product_specific_type,omitempty"`
	// ItemGroupIDs shopping_ads_type 为 PRODUCT 且product_specific_type为CUSTOMIZED_PRODUCTS时必填。
	// 商品 SPU（标准化产品单元）ID 列表。
	// 最大数量：50。
	// 若想获取某个 TikTok Shop 中商品的 SPU ID 列表，可使用 /store/product/get/。将 ad_creation_eligible 设置为 GMV_MAX 并从返回中挑选 status 为 AVAILABLE 且 gmv_max_ads_status 为 UNOCCUPIED 的 item_group_id 值
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// OptimizationGoal 优化目标。
	// 枚举值：
	// VALUE：总收入。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// DeepBidType 出价策略。
	// 枚举值：
	// VO_MIN_ROAS: 最低 ROAS。在有 ROI 目标的条件下出价。
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// RoasBid 仅当 deep_bid_type 为 VO_MIN_ROAS 时返回本字段。
	// ROI 目标。
	RoasBid float64 `json:"roas_bid,omitempty"`
	// Budget 日预算
	Budget float64 `json:"budget,omitempty"`
	// ScheduleType 排期类型。
	// 枚举值：
	// SCHEDULE_FROM_NOW：在排定的开始时间之后持续投放推广系列。
	// SCHEDULE_START_END:：在排定的开始时间和结束时间之间投放推广系列。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 广告投放起始时间(UTC+0)。
	// 格式：YYYY-MM-DD HH:MM:SS。
	ScheduleStartTime string `json:"schedule_start_time,omitempty"`
	// ScheduleEndTime 广告投放结束时间(UTC+0)。
	// 格式： YYYY-MM-DD HH:MM:SS
	ScheduleEndTime string `json:"schedule_end_time,omitempty"`
	// ProductVideoSpecificType 视频选择模式。
	// 枚举值：
	// AUTO_SELECTION：自动选择视频。
	// CUSTOM_SELECTION：手动选择视频。
	// UNSET：未设置。
	ProductVideoSpecificType enum.ProductVideoSpecificType `json:"product_video_specific_type,omitempty"`
	// IdentityList shopping_ads_type 为 LIVE时，本字段必填。
	// shopping_ads_type 为 PRODUCT 且 product_video_specific_type 为 AUTO_SELECTION时，本字段为可选字段。
	// 要与 GMV Max 推广系列绑定的认证身份（即 TikTok 账号）列表。
	// shopping_ads_type 为 PRODUCT 且 product_video_specific_type 为 AUTO_SELECTION 时，需通过本字段指定 0-20 个认证身份作为视频创意的来源。
	// 若想获取可用于使用同一 TikTok Shop 的商品 GMV Max 推广系列的认证身份列表，可使用 /gmv_max/identity/get/，并选择 product_gmv_max_available 为 true 的认证身份。
	// 若未通过 identity_list 指定认证身份，商品图片将用作广告创意。
	// shopping_ads_type 为 LIVE 时，需通过本字段指定 1 个认证身份作为直播来源。
	// 若想获取可用于直播 GMV Max 推广系列的认证身份，可调用 /gmv_max/identity/get/ 并检查返回的 live_gmv_max_available 是否为 true
	IdentityList []identity.Identity `json:"identity_list,omitempty"`
	// AffiliatePostsEnabled 此字段仅在满足以下所有条件时生效且使用默认值 true：
	// shopping_ads_type 为 PRODUCT
	// product_video_specific_type 为 AUTO_SELECTION
	// 是否为商品 GMV Max 推广系列启用联盟作品。
	// 联盟帖子是由联盟方（即参加 TikTok Shop 联盟项目的达人）制作并被授权用于 TikTok 店铺广告的 TikTok 帖子。了解 TikTok 店铺广告的联盟创意的详情。
	// 支持的值： true， false。
	AffiliatePostsEnabled *bool `json:"affiliate_posts_enabled,omitempty"`
	// ItemList shopping_ads_type 为 PRODUCT 且 product_video_specific_type 为 CUSTOM_SELECTION 时必填。
	// 要与商品 GMV Max 推广系列绑定的已授权 TikTok 作品（即 TikTok 帖子）或自定义的作品。
	// 已授权的作品代表来自您的 TikTok 账号或通过视频代码授权您使用，能够充分展示你所选商品的现有作品。
	// 若您想为此商品 GMV Max 推广系列仅手动选择已授权的 TikTok 作品，可通过 item_list 指定这些帖子，无需传入 custom_anchor_video_list。
	// 自定义的作品代表创建推广系列时通过向所指定视频添加自定义商品链接而新建的帖子。这些帖子仅用于此商品 GMV Max 推广系列。
	// 若您想为此商品 GMV Max 推广系列手动选择自定义的 TikTok 作品，需通过 item_list 指定要用于自定义作品的视频，且同时在 custom_anchor_video_list 中传入这些视频。
	// 最大数量：50。
	// 若想获取可用于使用同一 TikTok Shop 的商品 GMV Max 推广系列的 TikTok 帖子列表，可使用 /gmv_max/video/get/。
	Items []TikTokItem `json:"items,omitempty"`
	// CustomAnchorVideoList 与该 GMV Max 推广系列绑定的自定义 TikTok 作品（即 TikTok 帖子）。
	// 自定义的作品代表创建推广系列时通过向所指定视频（即 custom_anchor_video_list 中的 item_id）添加商品（即 custom_anchor_video_list 中的 spu_id_list）的自定义商品链接而新建的帖子。这些帖子仅用于此商品 GMV Max 推广系列。
	// 手动选择视频时：若 shopping_ads_type 为 PRODUCT 且 product_video_specific_type 为 CUSTOM_SELECTION，需通过本字段指定自定义的作品中手动选择的视频。
	// 自动选择视频时：若 shopping_ads_type 为 PRODUCT 且 product_video_specific_type 为 AUTO_SELECTION，需通过本字段指定自定义的作品自动选择的视频范围。
	// 最大数量：200。
	// 注意：创建商品 GMV Max 推广系列时最多仅可添加 200 个自定义的作品，但单个商品 GMV Max 推广系列最多可包含 2,000 个自定义的作品。若您想在单个商品 GMV Max 推广系列中使用 200 个以上的自定义的作品，可使用 /campaign/gmv_max/update/ 中的 custom_anchor_video_list 字段增量更新推广系列，每次请求可新增 200 个作品。
	CustomAnchorVideoList []AnchorVideo `json:"custom_anchor_video_list,omitempty"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建 GMV Max 推广系列 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
