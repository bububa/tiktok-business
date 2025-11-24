package comment

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest 更新评论状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CommentIDs 评论 ID
	CommentIDs []string `json:"comment_ids,omitempty"`
	// Operation 操作类型。枚举值：隐藏 HIDDEN， 公开 PUBLIC
	Operation enum.CommentStatus `json:"operation,omitempty"`
}

// Encode implements PostRequest
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
