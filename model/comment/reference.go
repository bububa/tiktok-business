package comment

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// ReferenceRequest 获取关联评论 API Request
type ReferenceRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CommentID 评论 ID
	CommentID string `json:"comment_id,omitempty"`
	// CommentType 评论类型。
	// 枚举值：原始评论 COMMENT，回复评论 REPLY。
	CommentType enum.CommentType `json:"comment_type,omitempty"`
	// OriginalCommentID 原始评论 ID。评论类型为 REPLY 时需要填写
	OriginalCommentID string `json:"original_comment_id,omitempty"`
	// Page 当前页数
	// 默认值：1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10。
	// 取值范围：1-100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *ReferenceRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("comment_id", r.CommentID)
	values.Set("comment_type", string(r.CommentType))
	if r.OriginalCommentID != "" {
		values.Set("original_comment_id", r.OriginalCommentID)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}
