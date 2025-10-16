package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OpenURLRequest 获取 TikTok 应用内链接 API Request
type OpenURLRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// URL 您想要获取对应应用内链接的TikTok公开网址。
	URL string `json:"url,omitempty"`
	// URLType TikTok公开网址类型。枚举值：
	// USER_PROFILE：用户档案页面。
	// VIDEO：视频详情页面。
	// HASHTAG_CHALLENGE：挑战赛页面。
	// MUSIC：音乐页面。
	// STICKER：贴纸（特效）页面。
	// STICKER_SHOOTER：贴纸开拍页面。
	URLType enum.TikTokURLType `json:"url_type,omitempty"`
}

// Encode implements GetRequest interface
func (r *OpenURLRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("url", r.URL)
	values.Set("url_type", string(r.URLType))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// OpenURLResponse 获取 TikTok 应用内链接 API Response
type OpenURLResponse struct {
	model.BaseResponse
	Data *OpenURLResult `json:"data,omitempty"`
}

type OpenURLResult struct {
	// OpenURL TikTok应用内链接
	OpenURL string `json:"open_url,omitempty"`
	// SupportedRegions 贴纸（特效）可以使用的国家和地区。仅当url_type 为STICKER 或STICKER_SHOOTER时返回
	SupportedRegions []string `json:"supported_regions,omitempty"`
}
