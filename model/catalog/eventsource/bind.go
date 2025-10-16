package eventsource

import "github.com/bububa/tiktok-business/util"

// BindRequest 绑定商品库与事件源 API Request
type BindRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// AppID 移动应用ID。您可以在TikTok Ads Manager中进入资产 > 事件 > 应用事件中获取。您至少需要传入移动应用ID(app_id)或者网站Pixel代码(pixel_code)两者之一。若您想要同时绑定应用事件和网页事件，可以两者都传入。
	AppID string `json:"app_id,omitempty"`
	// PixelCode 网站Pixel代码。您可以在TikTok Ads Manager中进入资产 > 事件 > 网页事件中获取。您至少需要传入移动应用ID(app_id)或者网站Pixel代码(pixel_code)两者之一。若您想要同时绑定应用事件和网页事件，可以两者都传入。
	PixelCode string `json:"pixel_code,omitempty"`
}

// Encode implements PostRequest interface
func (r *BindRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
