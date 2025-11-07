package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新 GMV Max 推广系列 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID。
	// 若想确认广告账号拥有您想要推广的 TikTok Shop 的 GMV Max 专属授权，可调用/gmv_max/store/list/，并确认返回的 is_gmv_max_available 为 true 且返回的 exclusive_authorized_advertiser_info 中的 advertiser_id 与您想要用来创建 GMV Max 推广系列的广告账户一致。
	// 若不一致，则可使用 /gmv_max/exclusive_authorization/create/授予此广告账户该 TikTok Shop 的专属 GMV Max 授权。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID GMV Max 推广系列 ID。
	// 若想创建 GMV Max 推广系列并获取对应 ID，可使用/campaign/gmv_max/create/。
	// 若想筛选广告账户中的现有 GMV Max 推广系列，可使用/gmv_max/campaign/get/。
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName GMV Max 推广系列的名称
	CampaignName string `json:"campaign_name,omitempty"`
	// RoasBid 仅当 deep_bid_type 为 VO_MIN_ROAS 时生效。
	// ROI 目标。
	// 若想获取推荐 ROI 目标，可使用 /gmv_max/bid/recommend/ 并查看返回的 roas_bid
	RoasBid float64 `json:"roas_bid,omitempty"`
	// Budget 日预算
	Budget float64 `json:"budget,omitempty"`
	// ItemGroupIDs shopping_ads_type 为 PRODUCT 且product_specific_type为CUSTOMIZED_PRODUCTS时必填。
	// 商品 SPU（标准化产品单元）ID 列表。
	// 最大数量：50。
	// 若想获取某个 TikTok Shop 中商品的 SPU ID 列表，可使用 /store/product/get/。将 ad_creation_eligible 设置为 GMV_MAX 并从返回中挑选 status 为 AVAILABLE 且 gmv_max_ads_status 为 UNOCCUPIED 的 item_group_id 值
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
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
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新 GMV Max 推广系列 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
