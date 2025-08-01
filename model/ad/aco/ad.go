package aco

import "github.com/bububa/tiktok-business/enum"

// Ad 智能创意广告
type Ad struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID 广告组ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// MediaInfoList 创意信息列表
	MediaInfoList []MediaInfo `json:"media_info_list,omitempty"`
	// TitleList 广告文案列表
	TitleList []Title `json:"title_list,omitempty"`
	// CallToActionList 行动引导文案列表
	CallToActionList []CallToAction `json:"call_to_action,omitempty"`
	// DeeplinkList 深度链接列表
	DeeplinkList []Deeplink `json:"deeplink_list,omitempty"`
	// DisplayNameList 显示名称
	DisplayNameList []DisplayName `json:"display_name_list,omitempty"`
	// AvatarIconList 广告头像图片列表
	AvatarIconList []AvatarIcon `json:"avatar_icon_list,omitempty"`
	// PageList 页面 ID 列表
	PageList []Page `json:"page_list,omitempty"`
	// CardList 卡片ID列表。长度范围：[0,1]
	CardList []Card `json:"card_list,omitempty"`
	// LandingPageURLs 落地页URL列表
	LandingPageURLs []LandingPageURL `json:"landing_page_urls,omitempty"`
	// CommonMaterial 常用素材
	CommonMaterial *CommonMaterial `json:"common_material,omitempty"`
}

// MediaInfo 创意信息
type MediaInfo struct {
	// MaterialID 素材ID
	MaterialID string `json:"material_id,omitempty"`
	// MaterialOperationStatus 素材状态。枚举值：ENABLE (素材可用)，DISABLE(素材不可用）。
	MaterialOperationStatus enum.EnableDisable `json:"material_operation_status,omitempty"`
	// MaterialName 素材（图片或视频）的名称。
	// 本字段用于决定素材在 TikTok 广告管理平台展示的名称。
	// 若素材类型为图片，则 TikTok 广告管理平台展示的名称与 material_name 的值一致。
	// 若素材类型为视频，则 TikTok 广告管理平台展示的名称为file_name 或 material_name 的值。
	// file_name不为空时，展示的素材名称为file_name的值。
	// file_name为空时，展示的素材名称为material_name的值。
	MaterialName string `json:"material_name,omitempty"`
	// IdentityID 认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。使用Spark Ads(tiktok_item_id)时，该字段必填。枚举值： AUTH_CODE, TT_USER。获取详细信息，参阅认证身份。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// TiktokItemID 用作 Spark Ad 的 TikTok 帖子 ID
	TiktokItemID string `json:"tiktok_item_id,omitempty"`
	// AigcDisclosureType 是否启用 AI 生成内容自主声明开关，以表明广告中包含 AI 生成内容。启用该开关后，当用户查看完整广告时，您的广告将带有“广告主标记为 AI 生成”的标签。
	//
	// 枚举值：
	// SELF_DISCLOSURE：启用开关，声明广告包含 AI 生成内容。
	// NOT_DECLARED：不声明广告包含 AI 生成内容。
	AigcDisclosureType enum.AigcDisclosureType `json:"aigc_disclosure_type,omitempty"`
	// ImageInfo 图片信息。
	// 若素材类型为视频，image_info 中将包含视频封面图的详情。
	// 若素材类型为图片，image_info 中将包含图片素材的详情。
	ImageInfo []ImageInfo `json:"image_info,omitempty"`
	// VideoInfo 仅当素材类型为视频时返回。
	// 视频信息。
	VideoInfo *VideoInfo `json:"video_info,omitempty"`
	// MediaInfo 创意信息
	MediaInfo *MediaInfo `json:"media_info,omitempty"`
}

// ImageInfo 图片信息。
type ImageInfo struct {
	// WebURI 图片ID
	WebURI string `json:"web_uri,omitempty"`
	// FileName 使用/ad/aco/create/创建智能创意广告时指定的用于生成广告名称的图片名。
	// 若您创建智能创意广告时未在image_info对象数组中传入file_name，本字段的值将为空字符串。
	FileName string `json:"file_name,omitempty"`
}

// VideoInfo 视频信息。
type VideoInfo struct {
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// FileName 使用/ad/aco/create/创建智能创意广告时指定的用于生成广告名称的视频名称。
	// 若您创建智能创意广告时未在video_info对象中传入file_name，本字段的值将为空字符串。
	FileName string `json:"file_name,omitempty"`
}

// Title 广告文案
type Title struct {
	// Title 广告文t案
	Title string `json:"title,omitempty"`
	// MaterialID 素材ID
	MaterialID string `json:"material_id,omitempty"`
	// MaterialOperationStatus 素材状态。枚举值：ENABLE (素材可用)，DISABLE(素材不可用）。
	MaterialOperationStatus enum.EnableDisable `json:"material_operation_status,omitempty"`
}

// CallToAction 行动引导文案
type CallToAction struct {
	// CallToAction 行动引导文案。枚举值可参见枚举值 - 行动引导文案。
	CallToAction enum.CallToAction `json:"call_to_action,omitempty"`
}

// Deeplink 深度链接
type Deeplink struct {
	// Deeplink 深度链接：用户下载App后点击链接跳转至的特定页面。
	Deeplink string `json:"deeplink,omitempty"`
	// DeeplinkType 深度链接类型。 不同推广目标的允许值不同：NORMAL(普通 DeepLink，支持 Traffic、Conversion)， DEFERRED_DEEPLINK(延迟 DeepLink，支持 App Install )。默认值：NORMAL。对于iOS 14专属推广系列中包含的广告，此字段无法设置为DEFERRED_DEEPLINK。
	DeeplinkType enum.DeeplinkType `json:"deeplink_type,omitempty"`
}

// DisplayName 显示名称
type DisplayName struct {
	// AppName App名称
	AppName string `json:"app_name,omitempty"`
	// LandingPageName 落地页名称
	LandingPageName string `json:"landing_page_name,omitempty"`
}

// AvatarIcon 广告头像图片
type AvatarIcon struct {
	// AvatarIcon 广告头像图片信息
	AvatarIcon *ImageInfo `json:"avatar_icon,omitempty"`
}

// Page 页面 ID
type Page struct {
	// PageID 页面 ID
	PageID string `json:"page_id,omitempty"`
}

// Card 卡片ID
type Card struct {
	// CardID 图片卡片ID、礼品码卡片ID、倒计时贴纸ID或下载卡ID
	CardID string `json:"card_id,omitempty"`
}

// LandingPageURL 落地页URL
type LandingPageURL struct {
	// LandingPageURL 落地页URL
	LandingPageURL string `json:"landing_page_url,omitempty"`
}

// CommonMaterial 常用素材
type CommonMaterial struct {
	// AdName 广告名称
	AdName string `json:"ad_name,omitempty"`
	// CallToActionID 行动引导文案素材包ID
	CallToActionID string `json:"call_to_action_id,omitempty"`
	// CreativeAuthorized 是否允许在我们的 TikTok For Business 创意中心展示您的广告。仅对非美国广告主有效。默认值：false。
	// 注意：creative_authorized仅可用于视频广告。
	CreativeAuthorized bool `json:"creative_authorized,omitempty"`
	// PlayableURL 试玩素材 URtL
	PlayableURL string `json:"playable_url,omitempty"`
	// FallbackType 失败路径。如果用户没有安装过应用，您可以选择让用户返回去安装应用，或者去您想推广的网页。 不适用于 Deferred DeepLink。 枚举值：APP_INSTALL， WEBSITE， UNSET。
	// 注意：不适用于延迟Deeplink。
	FallbackType enum.FallbackType `json:"fallback_type,omitempty"`
	// TrackingInfo 监测信息
	TrackingInfo *TrackingInfo `json:"tracking_info,omitempty"`
	// IdentityID 认证身份ID，当不使用Spark Ads(tiktok_item_id)时必填。
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型，当不使用Spark Ads(tiktok_item_id)时必填。仅支持CUSTOMIZED_USER
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IsSmartCreative 广告组是否启用智能创意
	IsSmartCreative bool `json:"is_smart_creative,omitempty"`
}

// TrackingInfo 监测信息
type TrackingInfo struct {
	// ImpressionTrackingURLs 默认展示监测 URL
	ImpressionTrackingURLs []string `json:"impression_tracking_urls,omitempty"`
	// ClickTrackingURLs 点击监测URL
	ClickTrackingURLs []string `json:"click_tracking_urls,omitempty"`
	// TrackingPixelID 正在追踪的pixel ID。你可以使用此字段追踪发生在外部网站的事件。
	// 注意： 该字段处于测试阶段。如果您要使用该字段，请联系您的TikTok客户代表。
	TrackingPixelID string `json:"tracking_pixel_id,omitempty"`
	// TrackingAppID 正在追踪的App的ID。
	// 注意：当对应的推广系列推广目标为APP_PROMOTION，PRODUCT_SALES，或 ENGAGEMENT时，本字段为白名单功能。如需在这些推广目标下使用此功能，请联系您的TikTok销售代表。
	TrackingAppID string `json:"tracking_aptp_id,omitempty"`
	// TrackingOfflineEventSetIDs 正在追踪的线下事件组ID列表。
	// 点击这里了解如何创建和管理线下事件组。
	// 注意：线下事件追踪目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	TrackingOfflineEventSetIDs []string `json:"tracking_offline_event_set_ids,omitempty"`
}
