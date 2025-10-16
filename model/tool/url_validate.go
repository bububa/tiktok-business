package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// URLValidateRequest 获取 URL 的验证结果 API Request
type URLValidateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// URL 您想要验证的 URL
	URL string `json:"url,omitempty"`
}

// Encode implements GetRequest interface
func (r *URLValidateRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("url", r.URL)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// URLValidateResponse 获取 URL 的验证结果 API Response
type URLValidateResponse struct {
	model.BaseResponse
	Data struct {
		// LandingPageURLInfo 与该 URL 有关的信息
		LandingPageURLInfo struct {
			// ValidateInfo 该 URL 的验证结果信息
			ValidateInfo *URLValidateInfo `json:"validate_info,omitempty"`
		} `json:"landing_page_url_info"`
	} `json:"data"`
}

type URLValidateInfo struct {
	// IsSchemeLink 该 URL 是否为自定义网址架构
	IsSchemeLink bool `json:"is_scheme_link,omitempty"`
	// IsValidURL 该 URL 是否为合法的 URL （iOS 通用链接，Android 应用链接 ，或自定义网址架构）
	IsValidURL bool `json:"is_valid_url,omitempty"`
	// IsValidScheme 该 URL 是否为合法的自定义网址架构
	IsValidScheme bool `json:"is_valid_scheme,omitempty"`
}
