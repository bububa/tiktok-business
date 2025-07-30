package enum

// CampaignSystemOrigin 推广系列来源
type CampaignSystemOrigin string

const (
	// CampaignSystemOrigin_PROMOTE：该推广系列为通过 TikTok 移动应用创建的内容加热推广系列。
	CampaignSystemOrigin_PROMOTE CampaignSystemOrigin = "PROMOTE"
	// CampaignSystemOrigin_TT_ADS_PLATFORM：该推广系列为通过 API 或TikTok 广告管理平台创建的非内容加热推广系列。
	CampaignSystemOrigin_TT_ADS_PLATFORM CampaignSystemOrigin = "TT_ADS_PLATFORM"
)
