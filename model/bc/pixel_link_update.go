package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// PixelLinkUpdateRequest 将pixel和广告账户绑定/解绑 API Request
type PixelLinkUpdateRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// PixelCode 待转移的Pixel代码
	PixelCode string `json:"pixel_code,omitempty"`
	// AdvertiserIDs 广告账户列表
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// RelationStatus 枚举值: LINK, UNLINK
	RelationStatus enum.PixelLinkRelationStatus `json:"relation_status,omitempty"`
}

// Encode implements PostRequest
func (r *PixelLinkUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
