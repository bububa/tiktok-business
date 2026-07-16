package brand

// Profile TTO Brand Profile
type Profile struct {
	// BrandProfileID Brand Profile ID
	BrandProfileID string `json:"brand_profile_id,omitempty"`
	// BrandName 品牌名称
	BrandName string `json:"brand_name,omitempty"`
	// BrandIndustryID 品牌行业 ID
	BrandIndustryID string `json:"brand_industry_id,omitempty"`
	// BrandIndustry is returned by the Brand Profile list endpoint.
	BrandIndustry string `json:"brand_industry,omitempty"`
	// BrandWebsite 品牌网站链接
	BrandWebsite string `json:"brand_website,omitempty"`
	// LogoURL TTO Creator Marketplace 账户的品牌 logo 图片 URL
	LogoURL string `json:"logo_url,omitempty"`
	// TikTokAccountURL TikTok 账户 URL
	TikTokAccountURL string `json:"tiktok_account_url,omitempty"`
}
