package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AdPreviewRequest 预览已有广告 API Request
type AdPreviewRequest struct {
	CreateRequest
	// AdID 您想要预览的广告 ID
	AdID string `json:"ad_id,omitempty"`
	// Device 设备型号列表。
	// 枚举值见/tool/device_model/。
	Device []string `json:"device,omitempty"`
	// Language 广告所用语言。
	// 枚举值: ENGLISH, CHINESE, JAPANESE。
	// 默认值: ENGLISH。
	Language string `json:"language,omitempty"`
	// Placement 广告版位。
	// 枚举值：
	// PLACEMENT_TIKTOK：仅在objective_type 为REACH， TRAFFIC, VIDEO_VIEWS，ENGAGEMENT，APP_PROMOTION， LEAD_GENERATION，WEB_CONVERSIONS 或 PRODUCT_SALES时支持。
	// PLACEMENT_PANGLE：仅在objective_type 为TRAFFIC，LEAD_GENERATION（仅当 promotion_type为LEAD_GENERATION时），APP_PROMOTION 或 WEB_CONVERSIONS 时支持。
	// ALL：仅在objective_type 为TRAFFIC，LEAD_GENERATION（仅当 promotion_type为LEAD_GENERATION时），APP_PROMOTION 或 WEB_CONVERSIONS 时支持。
	//
	// 注意：若想为广告设置非 TikTok 版位，需在传入本字段进行设置。否则，本字段将默认设置为 TikTok 版位 (PLACEMENT_TIKTOK) 。
	Placement enum.Placement `json:"placement,omitempty"`
}

// Encode implements PostRequest interface
func (r *AdPreviewRequest) Encode() []byte {
	r.PreviewType = enum.AdPreviewType_AD
	return util.JSONMarshal(r)
}
