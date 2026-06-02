package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CustomAnchorVideoListCreateRequest Create shop-level customized TikTok posts API Request
type CustomAnchorVideoListCreateRequest struct {
	// AdvertiserID Advertiser ID.
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID The ID of a TikTok Shop.
	// To obtain a TikTok Shop that is available for GMV Max Campaigns, use /gmv_max/store/list/ and confirm that the returned is_gmv_max_available is true.
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID The ID of the Business Center that is authorized to access the TikTok Shop (store_id).
	// To obtain the ID, check the store_authorized_bc_id for the store_id returned from /gmv_max/store/list/.
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// CustomAnchorVideoList The details of your videos and the products you want to link them to.
	// By linking videos to a product at the TikTok Shop level, you can generate shop-level customized TikTok posts. The system will automatically add these shop-level customized TikTok posts to your Product GMV Max Campaign (with shopping_ads_type set to PRODUCT) if you use the Autoselect creative mode (with product_video_specific_type as AUTO_SELECTION) and your campaign includes the linked product.
	// Max size: 200.
	CustomAnchorVideoList []AnchorVideo `json:"custom_anchor_video_list,omitempty"`
}

// Encode implements PostRequest interface
func (r *CustomAnchorVideoListCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CustomAnchorVideoListCreateResponse Create shop-level customized TikTok posts API Response
type CustomAnchorVideoListCreateResponse struct {
	model.BaseResponse
	Data struct {
		// FailureList The list of shop-level customized TikTok posts that could not be successfully created.
		FailureList []CustomAnchorVideoListCreateError `json:"failure_list,omitempty"`
	} `json:"data"`
}

type CustomAnchorVideoListCreateError struct {
	AnchorVideo
	// Reason The reason for the failure.
	// Enum values:
	// IDENTITY_NO_PERMISSION: You do not have the required permissions for this identity.
	// INVALID_PARAMETER: Invalid parameters are submitted.
	// AUTH_CODE_CAN_NOT_CHANGE_ANCHOR: The anchor in the video cannot be changed without user permission.
	// ITEM_NOT_FOUND: The TikTok video cannot be found.
	// NATIVE_ANCHOR_EXISTS: The product link you want to add is the same as the product link already present in the video.
	Reason string `json:"reason,omitempty"`
	// ErrorMessage The detailed error message
	ErrorMessage string `json:"error_message,omitempty"`
}

func (r CustomAnchorVideoListCreateError) Error() string {
	return r.ErrorMessage
}
