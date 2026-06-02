package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ShopVideoVideoAnchorsRequest Get product linkage details of videos in customized posts API Request
type ShopVideoVideoAnchorsRequest struct {
	// AdvertiserID Advertiser ID.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID The ID of a TikTok Shop.
	// To obtain a TikTok Shop that is available for GMV Max Campaigns, use /gmv_max/store/list/ and confirm that the returned is_gmv_max_available is true.
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID The ID of the Business Center that is authorized to access the TikTok Shop (store_id).
	// To obtain the ID, check the store_authorized_bc_id for the store_id returned from /gmv_max/store/list/.
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// ItemIDs The list of TikTok videos for which you want to retrieve product linkage information.
	// To retrieve TikTok videos used in customized posts, use /gmv_max/creation/custom_anchor_video_list/get/.
	// Max size: 20.
	ItemIDs []string `json:"item_ids,omitempty"`
	// CampaignID Required when the specified TikTok videos (item_id) have been used in a campaign to create campaign-level customized posts.
	// The ID of a Product GMV Max Campaign.
	// To retrieve existing Product GMV Max Campaigns within your ad account, use /gmv_max/campaign/get/ and set filtering to {"gmv_max_promotion_types":["PRODUCT_GMV_MAX"]}.
	CampaignID string `json:"campaign_id,omitempty"`
}

// Encode implements PostRequest
func (r *ShopVideoVideoAnchorsRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ShopVideoVideoAnchorsResponse Get product linkage details of videos in customized posts API Response
type ShopVideoVideoAnchorsResponse struct {
	model.BaseResponse
	Data struct {
		// VideoList The product linkage details for the specified videos.
		VideoList []ShopVideoAnchor `json:"video_list,omitempty"`
	} `json:"data"`
}

// ShopVideoAnchor The product linkage details for the specified video.
type ShopVideoAnchor struct {
	// ItemID The ID of the TikTok video
	ItemID string `json:"item_id,omitempty"`
	// ProductList The list of products that the video is linked to
	ProductList []ShopVideoAnchorProduct `json:"product_list,omitempty"`
}

// ShopVideoAnchorProduct The product that the video is linked to
type ShopVideoAnchorProduct struct {
	// SpuID The SPU ID of the product linked to the TikTok video.
	SpuID string `json:"spu_id,omitempty"`
	// Title The title of the product
	Title string `json:"title,omitempty"`
	// Picture The URL of the product image
	Picture string `json:"picture,omitempty"`
	// AnchorSource The source of the product anchor in the video.
	// Enum values:
	// CUSTOM: custom anchor. The product anchor was added to the video through Shop Creative Hub for GMV Max. The linked product will appear in ads traffic.
	// ORGANIC: organic anchor. The product anchor was added to the video through the TikTok app. The linked product will appear in both ads and organic traffic.
	AnchorSource string `json:"anchor_source,omitempty"`
}
