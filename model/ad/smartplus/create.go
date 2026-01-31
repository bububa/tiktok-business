package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest Create an Upgraded Smart+ Ad API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID Ad group ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
	// OperationStatus 广告的操作状态。
	// ENABLE：广告处于启用（“开”）状态。
	// DISABLE：广告处于未启用（“关”）状态。
	// FROZEN：广告处于冻结状态，无法再次启用。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// CreativeList A list of creatives.
	// Size range: 1-50
	CreativeList []Creative `json:"creative_list,omitempty"`
	// AdTextList A list of ad texts.
	// Ad texts are shown to your audience as part of your ad creatives, to deliver the message you intend to communicate to them.
	AdTextList []AdText `json:"ad_text_list,omitempty"`
	// AutoMessageList Returned only when promotion_type is LEAD_GEN_CLICK_TO_TT_DIRECT_MESSAGE at the ad group level.
	// Details of the automatic message to use in a TikTok Direct Messaging Ad.
	AutoMessageList []AutoMessage `json:"auto_message_list,omitempty"`
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
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse Create an Upgraded Smart+ Ad API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Ad `json:"data,omitempty"`
}
