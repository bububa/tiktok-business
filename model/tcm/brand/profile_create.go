package brand

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建 TTO Brand Profile API Request
type CreateRequest struct {
	// TTOTCMAccountID TTO Creator Marketplace 账户 ID
	TTOTCMAccountID string `json:"tto_tcm_account_id,omitempty"`
	// BrandName 品牌名称。
	// 长度限制：60个字符。
	// 注意：设置后不可更新。
	BrandName string `json:"brand_name,omitempty"`
	// BrandIndustryID 品牌行业 ID。
	// 注意：设置后不可更新。
	BrandIndustryID string `json:"brand_industry_id,omitempty"`
	// BrandWebsite 品牌网站链接。
	// 链接应以 https:// 开头。
	BrandWebsite string `json:"brand_website,omitempty"`
	// LogoURL 品牌 logo 图片 URL。
	// 支持 JPG、PNG 或 GIF，宽高比 1:1，文件大小不超过 5MB。
	// 注意：设置后不可更新。
	LogoURL string `json:"logo_url,omitempty"`
	// TikTokAccountURL TikTok 账户 URL。
	// 链接应以 https://www.tiktok.com/@ 开头。
	TikTokAccountURL string `json:"tiktok_account_url,omitempty"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建 TTO Brand Profile API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Profile `json:"data,omitempty"`
}
