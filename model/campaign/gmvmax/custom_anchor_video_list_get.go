package gmvmax

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/file/video"
	"github.com/bububa/tiktok-business/model/identity"
	"github.com/bububa/tiktok-business/util"
)

// CustomAnchorVideoListGetRequest Get customized TikTok posts API Request
type CustomAnchorVideoListGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID The ID of a TikTok Shop.
	// To obtain a TikTok Shop that is available for GMV Max Campaigns, use /gmv_max/store/list/ and confirm that the returned is_gmv_max_available is true.
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID The ID of the Business Center that is authorized to access the TikTok Shop (store_id).
	// To obtain the ID, check the store_authorized_bc_id for the store_id returned from /gmv_max/store/list/.
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// CreativeSource Creative source type.
	// Enum value:
	// CUSTOMIZED: customized post.
	CreativeSource string `json:"creative_source,omitempty"`
	// SpuIDList A list of Product SPU (standard product unit) IDs to filter the results by.
	// To obtain the list of SPU IDs for products within a TikTok Shop, use /store/product/get/. Set ad_creation_eligible to GMV_MAX and select item_group_id values where status is AVAILABLE and gmv_max_ads_status is UNOCCUPIED.
	SpuIDList []string `json:"spu_id_list,omitempty"`
	// SortField The sorting field.
	// Enum values:
	// GMV: GMV.
	// POST_TIME: Post time.
	// VIDEO_VIEWS: Video views.
	// VIDEO_LIKES: Video likes.
	// CLICK_THROUGH_RATE: Click-through Rate (CTR).
	// PRODUCT_CLICKS: Product clicks.
	// Default value: GMV.
	SortField string `json:"sort_field,omitempty"`
	// SortType The sorting order.
	// Enum values:
	// ASC: ascending.
	// DESC: descending.
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Keyword A search keyword to filter posts (videos).
	// You can specify a post ID or caption.
	// To search by post caption (text), provide a text string. This search type supports multiple-language searches and fuzzy matches.
	// Example: summer.
	// To search by post ID (item_id), provide a numeric string with at least 19 characters. This search type supports exact matches.
	// Example: 1234567891234567891.
	Keyword string `json:"keyword,omitempty"`
	// NeedAuthCodeVideo Whether to include posts associated with AUTH_CODE identities in the results.
	// Supported values: true, false.
	// Default value: false.
	// Note: If need_auth_code_video is false or not passed and identity_list is not specified, you will retrieve an empty post list.
	NeedAuthCodeVideo bool `json:"need_auth_code_video,omitempty"`
	// IdentityList A list of TT_USER, BC_AUTH_TT, or TTS_TT identities to filter the results by.
	// To obtain a list of identities available for Product GMV Max Campaigns, use /gmv_max/identity/get/ and select identities with product_gmv_max_available as true.
	// Max size: 20.
	// Note: If need_auth_code_video is false or not passed and identity_list is not specified, you will retrieve an empty post list.
	IdentityList []identity.Identity `json:"identity_list,omitempty"`
	// CampaignID The ID of a Product GMV Max Campaign to filter the returns by.
	// You can use this field to filter campaign-level customized posts. If this field is not specified, the results will only include shop-level customized posts.
	// To retrieve existing Product GMV Max Campaigns within your ad account, use /gmv_max/campaign/get/ and set filtering to {"gmv_max_promotion_types":["PRODUCT_GMV_MAX"]}.
	CampaignID string `json:"campaign_id,omitempty"`
	// Page Current page number.
	// Value range: ≥ 1.
	// Default value: 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size.
	// Value range: 1-50.
	// Default value:10.
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAnchorVideoListGetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CustomAnchorVideoListGetResponse Get customized TikTok posts API Response
type CustomAnchorVideoListGetResponse struct {
	model.BaseResponse
	Data *CustomAnchorVideoListGetResult `json:"data,omitempty"`
}

type CustomAnchorVideoListGetResult struct {
	// ItemList A list of TikTok videos (posts) available for Product GMV Max campaigns using the specified TikTok Shop.
	ItemList []AnchorVideo `json:"item_list,omitempty"`
	// PageInfo Pagination information.
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// AnchorVideo TikTok videos (posts) available for Product
type AnchorVideo struct {
	// ItemID TikTok 帖子 ID
	ItemID string `json:"item_id,omitempty"`
	// Text TikTok 帖子描述。
	// 示例：summer
	Text string `json:"text,omitempty"`
	// SpuIDList 与该 TikTok 帖子绑定的商品 SPU ID 列表
	SpuIDList []string `json:"spu_id_list,omitempty"`
	// CanChangeAnchor Whether you can change the product anchor link in this video.
	// Supported values:
	// true: You can add a product anchor link to create a customized post.
	// false: You cannot use the video to create a customized post.
	// Note: If can_change_anchor is false, you cannot pass the video to custom_anchor_video_list or item_list in /campaign/gmv_max/create/ or /campaign/gmv_max/update/ to create customized posts.
	CanChangeAnchor bool `json:"can_change_anchor,omitempty"`
	// IdentityInfo 与该 TikTok 帖子绑定的认证身份的有关信息
	IdentityInfo *identity.Identity `json:"identity_info,omitempty"`
	// VideoInfo 帖子中视频的详情
	VideoInfo *video.Video `json:"video_info,omitempty"`
}
