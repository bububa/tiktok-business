package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// MaterialStatusUpdateRequest Disable or enable creatives in an Upgraded Smart+ Ad API Request
type MaterialStatusUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SmartPlusAdID Ad ID
	SmartPlusAdID string `json:"smart_plus_ad_id,omitempty"`
	// AdMaterialIDs A list of ad material IDs for creatives to operate on.
	// An ad material ID is an ad-specific material ID generated when a particular creative is used in an ad.
	// This ID differs from the creative ID you receive when uploading the creative to your ad account’s Creative Library.
	//
	// To obtain ad material IDs associated with an ad, use /smart_plus/ad/get/.
	//
	// Max size: 20.
	//
	// The number of creatives you can operate on varies based on the type of creatives you want to operate on:
	// If you want to operate on the following types of creatives, you can only specify one ID in ad_material_ids:
	// a carousel in Spark Ad Standard Carousel created through Sparks Ads Push
	// a carousel in non-Spark Ad Standard Carousel
	// If you want to operate on the following types of creatives, you can specify up to 20 IDs in ad_material_ids:
	// non-Spark Ad Video
	// Spark Ad Video created through Sparks Ads Push
	// TikTok video posts in Spark Ad Video created through Sparks Ads Pull
	// TikTok photo posts in Spark Ad Standard Carousel created through Sparks Ads Pull
	AdMaterialIDs []string `json:"ad_material_ids,omitempty"`
	// OperationStatus The operation to perform on the creatives.
	// Enum values: DISABLE (pause), ENABLE (enable).
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
}

// Encode implements PostRequest
func (r *MaterialStatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
