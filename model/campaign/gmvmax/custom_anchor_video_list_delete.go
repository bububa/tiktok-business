package gmvmax

import (
	"github.com/bububa/tiktok-business/util"
)

// CustomAnchorVideoListDeleteRequest Delete customized TikTok posts API Request
type CustomAnchorVideoListDeleteRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID The ID of a TikTok Shop.
	// To obtain a TikTok Shop that is available for GMV Max Campaigns, use /gmv_max/store/list/ and confirm that the returned is_gmv_max_available is true.
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID The ID of the Business Center that is authorized to access the TikTok Shop (store_id).
	// To obtain the ID, check the store_authorized_bc_id for the store_id returned from /gmv_max/store/list/.
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// CustomAnchorVideoList The details of the videos and the products you want to unlink from them.
	// Max size: 200.
	CustomAnchorVideoList []AnchorVideo `json:"custom_anchor_video_list,omitempty"`
	// CampaignID Required when the specified TikTok videos (item_id) have been used in a campaign to create campaign-level customized posts.
	// The ID of a Product GMV Max Campaign.
	// To retrieve existing Product GMV Max Campaigns within your ad account, use /gmv_max/campaign/get/ and set filtering to {"gmv_max_promotion_types":["PRODUCT_GMV_MAX"]}.
	CampaignID string `json:"campaign_id,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAnchorVideoListDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
