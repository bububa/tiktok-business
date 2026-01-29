package smartplus

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest Update an Upgraded Smart+ Ad API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SmartPlusAdID Ad ID.
	// To obtain ad IDs, use /smart_plus/ad/get/.
	SmartPlusAdID string `json:"smart_plus_ad_id,omitempty"`
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
	// CreativeList A list of creatives.
	// Size range: 1-50
	CreativeList []Creative `json:"creative_list,omitempty"`
	// AdTextList A list of ad texts.
	// Ad texts are shown to your audience as part of your ad creatives, to deliver the message you intend to communicate to them.
	AdTextList []AdText `json:"ad_text_list,omitempty"`
	// CallToActionList Call-to-action list
	CallToActionList []CallToAction `json:"call_to_action,omitempty"`
	// InteractiveAddOnList A list of interactive add-on (creative portfolio) IDs.
	InteractiveAddOnList []InteractiveAddOn `json:"interactive_add_on_list,omitempty"`
	// PageList A list of pages
	PageList []Page `json:"page_list,omitempty"`
	// LandingPageURLList A list of landing page URLs
	LandingPageURLList []LandingPageURL `json:"landing_page_url_list,omitempty"`
	// CustomProductPageList Details of the custom product page to use in the ad
	CustomProductPageList []CustomProductPage `json:"custom_product_page_url,omitempty"`
	// DeeplinkList A list of deeplinks
	DeeplinkList []Deeplink `json:"deeplink_list,omitempty"`
	// Disclaimer Disclaimer information
	Disclaimer *Disclaimer `json:"disclaimer,omitempty"`
	// AdConfiguration Additional configurations
	AdConfiguration *AdConfiguration `json:"ad_configuration,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse Update an Upgraded Smart+ Ad API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Ad `json:"data,omitempty"`
}
