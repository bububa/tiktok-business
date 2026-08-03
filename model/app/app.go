package app

import "github.com/bububa/tiktok-business/enum"

// App contains app management information returned by TikTok.
type App struct {
	AdvertiserID                     string                   `json:"advertiser_id,omitempty"`
	AppID                            string                   `json:"app_id,omitempty"`
	TikTokAppID                      string                   `json:"tiktok_app_id,omitempty"`
	AppName                          string                   `json:"app_name,omitempty"`
	AppPlatformID                    string                   `json:"app_platform_id,omitempty"`
	AppType                          enum.AppType             `json:"app_type,omitempty"`
	Platform                         enum.OperatingSystem     `json:"platform,omitempty"`
	DownloadURL                      string                   `json:"download_url,omitempty"`
	IconURL                          string                   `json:"icon_url,omitempty"`
	PartnerID                        string                   `json:"partner_id,omitempty"`
	PartnerName                      string                   `json:"partner_name,omitempty"`
	EnableRetargeting                enum.AppRetargetingState `json:"enable_retargeting,omitempty"`
	TrackingURL                      *TrackingURL             `json:"tracking_url,omitempty"`
	SkanAllowed                      bool                     `json:"skan_allowed,omitempty"`
	SelfAttributionEnabled           bool                     `json:"self_attribution_enabled,omitempty"`
	AdvancedDedicatedCampaignAllowed bool                     `json:"advanced_dedicated_campaign_allowed,omitempty"`
}

// TrackingURL contains default attribution tracking URLs for an app.
type TrackingURL struct {
	ClickURL                 string `json:"click_url,omitempty"`
	ImpressionURL            string `json:"impression_url,omitempty"`
	RetargetingClickURL      string `json:"retargeting_click_url,omitempty"`
	RetargetingImpressionURL string `json:"retargeting_impression_url,omitempty"`
}
