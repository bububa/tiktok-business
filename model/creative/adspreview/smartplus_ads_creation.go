package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model/ad/smartplus"
	"github.com/bububa/tiktok-business/util"
)

// SmartPlusAdsCreationPreviewRequest Preview Upgraded Smart+ Ads API Request
type SmartPlusAdsCreationPreviewRequest struct {
	CreateRequest
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CatalogEnabled Whether to use catalog in the campaign.
	// Supported values: true, false.
	CatalogEnabled bool `json:"catalog_enabled,omitempty"`
	// CatalogID Required when catalog_enabled is true.
	// The ID of the catalog to use in the ad group.
	// To retrieve the catalogs within your Business Center, use /catalog/get/.
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogAuthorizedBcID Required when catalog_enabled is true.
	// The ID of the Business Center that the catalog (catalog_id) belongs to.
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// CreativeList A list of creatives.
	// Size range: 1-50
	CreativeList []smartplus.Creative `json:"creative_list,omitempty"`
	// AdTextList A list of ad texts.
	// Ad texts are shown to your audience as part of your ad creatives, to deliver the message you intend to communicate to them.
	AdTextList []smartplus.AdText `json:"ad_text_list,omitempty"`
	// CallToActionList Call-to-action list
	CallToActionList []smartplus.CallToAction `json:"call_to_action,omitempty"`
	// AdConfiguration Additional configurations
	AdConfiguration *smartplus.AdConfiguration `json:"ad_configuration,omitempty"`
}

// Encode implements PostRequest interface
func (r *SmartPlusAdsCreationPreviewRequest) Encode() []byte {
	r.PreviewType = enum.AdPreviewType_ADS_CREATION
	return util.JSONMarshal(r)
}

func (r *SmartPlusAdsCreationPreviewRequest) Gateway() string {
	return "smart_plus/ad/preview/"
}
