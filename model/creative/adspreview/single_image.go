package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// SingleImagePreviewRequest 预览图片 API Request
type SingleImagePreviewRequest struct {
	CreateRequest
	// ImageID 图片ID。不同版位对素材的要求不同。详情请参阅广告格式介绍。
	ImageID string `json:"image_id,omitempty"`
	// DisplayName 品牌名称。长度限制：1-40个英文字符或1-20个中文/日文/韩文字符
	DisplayName string `json:"display_name,omitempty"`
	// AdText 广告文案。字符数限制: 1-100。
	AdText string `json:"ad_text,omitempty"`
	// CallToAction 行动引导文案，引导用户在看到你的广告后完成你期望的行为。对于TikTok广告，本字段或者 call_to_action_id 必须传入一个，如果同时传入，此字段将被忽略。枚举值见枚举值 - 行动引导文案。直播购物事件下，此字段的值必须为 WATCH_LIVE。
	CallToAction enum.CallToAction `json:"call_to_action"`
	// CallToActionID 行动引导文案素材包ID。行动引导文案素材包指的是一组自动优化的行动引导文案。如果本字段和 call_to_action 都传入，call_to_action 会被忽略。关于自动优化行动引导文案的更多信息，请参见智能推荐CTA > 动态优选CTA。
	CallToActionID string `json:"call_to_action_id,omitempty"`
	// Placements 版位。枚举值见枚举值 - 版位。不同版位对素材的要求不同。详情请参阅广告格式介绍。
	Placements []enum.Placement `json:"placements,omitempty"`
	// Device 设备型号列表。枚举值见获取设备型号。
	Device []string `json:"device,omitempty"`
	// Language 广告所用语言。 枚举值: ENGLISH, CHINESE, JAPANESE。默认值: ENGLISH。
	Language string `json:"language,omitempty"`
}

// Encode implements PostRequest interface
func (r *SingleImagePreviewRequest) Encode() []byte {
	r.PreviewType = enum.AdPreviewType_SINGLE_IMAGE
	return util.JSONMarshal(r)
}
