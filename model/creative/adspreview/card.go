package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// CardPreviewRequest 预览已有创新互动样式 API Request
type CardPreviewRequest struct {
	CreateRequest
	// CardID 创意素材包 ID。
	// 传入以下类型的创意素材包的 ID：
	// 展示卡片（又称图片卡片）。
	// 倒计时贴纸。
	// 礼品码贴纸。
	//
	// 若想创建素材包，可使用 /creative/portfolio/create/。
	// 若想了解如何获取展示卡片的 ID，请参考卡片。
	// 若想了解如何获取倒计时贴纸或礼品码贴纸素材包的 ID，请参考贴纸。
	CardID string `json:"card_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *CardPreviewRequest) Encode() []byte {
	r.PreviewType = enum.AdPreviewType_CARD
	return util.JSONMarshal(r)
}
