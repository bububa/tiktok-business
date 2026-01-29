package smartplus

import "github.com/bububa/tiktok-business/enum"

// Creative 广告创意
type Creative struct {
	// AdMaterialID An ad-specific material ID generated when a particular creative is used in an ad.
	// This ID differs from the creative ID you receive when uploading the creative to your ad account’s Creative Library.
	AdMaterialID string `json:"ad_material_id,omitempty"`
	// MaterialOperationStatus The status of the creative.
	// Enum values:
	// ENABLE: The creative is enabled (in 'ON' status).
	// DISABLE: The creative is disabled (in 'OFF' status).
	MaterialOperationStatus enum.OperationStatus `json:"material_operation_status,omitempty"`
	// CreativeInfo Creative information
	CreativeInfo *CreativeInfo `json:"creative_info,omitempty"`
}

// CreativeInfo Creative information
type CreativeInfo struct {
	// AdFormat The ad format.
	// Enum values:
	// SINGLE_VIDEO: Single Video.
	// CAROUSEL_ADS: Standard Carousel.
	// CATALOG_CAROUSEL: Catalog Carousel.
	AdFormat enum.AdFormat `json:"ad_format,omitempty"`
	// MaterialName Material name
	MaterialName string `json:"material_name,omitempty"`
	// VideoInfo Video information
	VideoInfo *VideoInfo `json:"video_info,omitempty"`
	// ImageInfo Image information
	ImageInfo []ImageInfo `json:"image_info,omitempty"`
	// MusicInfo Music information
	MusicInfo *MusicInfo `json:"music_info,omitempty"`
	// AigcDisclosureType Whether to turn on the AIGC (Artificial Intelligence Generated Content) self-disclosure toggle to indicate the ad contains AI-generated content. After the toggle is turned on, your ad will carry an "Advertiser labeled as Al-generated" label when viewed in full.
	// Enum values:
	// SELF_DISCLOSURE: To turn on the toggle to declare that the ad contains AI-generated content.
	// NOT_DECLARED: To not declare that the ad contains AI-generated content.
	AigcDisclosureType enum.AigcDisclosureType `json:"aigc_disclosure_type,omitempty"`
	// TiktokItemID The ID of the TikTok post to be used as an ad (Spark Ads).
	TiktokItemID *string `json:"tiktok_item_id,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值： CUSTOMIZED_USER (自定义用户），AUTH_CODE (授权用户)， TT_USER (TikTok 企业号用户）， BC_AUTH_TT（添加到商务中心的TikTok用户）。
	// 关于认证身份的更多信息，请访问认证身份。
	// 注意：类型为 AUTH_CODE 、 TT_USER 或 BC_AUTH_TT 的认证身份仅可用于创建最多10,000个广告。若达到此限制，您可以通过删除使用该身份的广告来释放与该身份关联的广告配额。若想查询此类身份的剩余广告配额，您可以调用/asset/bind/quota/并查看返回的 available_quota 字段值。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityID 认证身份ID。
	// 如果您不使用Spark Ads，传入 identity_id 和 identity_type (CUSTOMIZED_USER类型)可帮助您更好地管理广告信息。
	// 注意：类型为 AUTH_CODE 、 TT_USER 或 BC_AUTH_TT 的认证身份仅可用于创建最多10,000个广告。若达到此限制，您可以通过删除使用该身份的广告来释放与该身份关联的广告配额。若想查询此类身份的剩余广告配额，您可以调用/asset/bind/quota/并查看返回的 available_quota 字段值。
	IdentityID *string `json:"identity_id,omitempty"`
	// IdentityAuthorizedBcID 当identity_type 为 BC_AUTH_TT时必填。
	// 与添加到商务中心的TikTok用户这一认证身份绑定的商务中心的ID。
	IdentityAuthorizedBcID *string `json:"identity_authorized_bc_id,omitempty"`
}

// VideoInfo Video information
type VideoInfo struct {
	// VideoID Video ID
	VideoID string `json:"video_id,omitempty"`
	// FileName Video name
	FileName string `json:"file_name,omitempty"`
}

// ImageInfo Image information
type ImageInfo struct {
	// WebURI Image ID
	WebURI string `json:"web_uri,omitempty"`
}

// MusicInfo Music information
type MusicInfo struct {
	// MusicID The ID of the piece of music to use in the Carousel Ads
	MusicID string `json:"music_id,omitempty"`
}
