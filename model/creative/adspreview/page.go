package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// PagePreviewRequest 预览页面 API Request
type PagePreviewRequest struct {
	CreateRequest
	// PageID 页面 ID。
	// 若想创建即时体验页面，可使用即时体验页面编辑器并将 businessType设置为 6。
	// 创建即时体验页面后，可使用/page/get/获取页面 ID。
	PageID string `json:"page_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *PagePreviewRequest) Encode() []byte {
	r.PreviewType = enum.AdPreviewType_PAGE
	return util.JSONMarshal(r)
}
