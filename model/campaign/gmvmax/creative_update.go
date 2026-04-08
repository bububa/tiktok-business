package gmvmax

import "github.com/bububa/tiktok-business/util"

// CreativeUpdateRequest Remove or add back creatives in a GMV Max Campaign API Request
type CreativeUpdateRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"avertiser_id,omitempty"`
	// CampaignID Campaign ID.
	// Note:
	// The campaign must be in an active status. To retrieve active GMV Max Campaigns in your ad account, use /gmv_max/report/get/ with filtering set to {"campaign_statuses":["STATUS_DELIVERY_OK"]}.
	// If the campaign is a Product GMV Max Campaign, the product_video_specific_type of the campaign must be AUTO_SELECTION.
	CampaignID string `json:"campaign_id,omitempty"`
	// Action The type of action to perform.
	// Enum values:
	// REMOVE: To remove creatives from a GMV Max Campaign.
	// ADD: To add previously removed creatives back into a GMV Max Campaign.
	// Note:
	// Once the action is performed, wait 20 minutes before verifying the updated statuses using /gmv_max/report/get/.
	// If a creative is successfully removed from a Product GMV Max Campaign, its creative_delivery_status in /gmv_max/report/get/ will be EXCLUDED.
	Action string `json:"action,omitempty"`
	// Items The list of TikTok posts (videos) to be operated on.
	// Max size: 400.
	// Note:
	// This endpoint allows for the removal of up to 10,000 posts from a GMV Max Campaign, with a limit of 400 posts per request.
	// For example, if you need to exclude 800 posts, send a first request with 400 posts in this field, followed by a second request with the remaining 400.
	Items []CreativeUpdateItem `json:"items,omitempty"`
}

type CreativeUpdateItem struct {
	// ItemID The ID of the TikTok post to operate on.
	// When action is REMOVE, use /gmv_max/report/get/ with filtering set to {"creative_delivery_statuses":["IN_QUEUE","LEARNING","DELIVERING","NOT_DELIVERYING","UNAVAILABLE","NOT_ACTIVE"]} to retrieve eligible creatives.
	// When action is ADD, use /gmv_max/report/get/ with filtering set to {"creative_delivery_statuses":["EXCLUDED"]} to retrieve previously removed creatives.
	// Note: Currently, you can only obtain the item_id of TikTok posts in Live GMV Max Campaigns on TikTok Ads Manager.
	ItemID string `json:"item_id,omitempty"`
	// SpuIDList Required for a Product GMV Max Campaign.
	// The list of Product SPU IDs that the TikTok post is associated with.
	// To obtain the list of Product SPU IDs associated with a TikTok post in a Product GMV Max Campaign, use /campaign/gmv_max/info/.
	SpuIDList []string `json:"spu_id_list,omitempty"`
}

// Encode implements PostRequest
func (r *CreativeUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
