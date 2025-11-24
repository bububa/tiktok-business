package comment

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// ReplyListRequest 获取评论的所有回复 API Request
type ReplyListRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// VideoID 自有 TikTok 视频的唯一标识符，用于列出对应评论（及评论回复）。可通过/business/video/list/ 接口的item_id字段获取
	VideoID string `json:"video_id,omitempty"`
	// CommentID 自有 TikTok 视频评论的唯一标识符，用于列出对应回复。可通过/business/comment/list/ 接口的comment_id字段获取所列的各个评论的 ID
	CommentID string `json:"comment_id,omitempty"`
	// Status 用于筛选结果的评论（及评论回复）可见状态。
	// 枚举值：
	// PUBLIC：所有 TikTok 用户均可公开查看的评论（及评论回复）。
	// ALL：视频发布者的隐藏评论（及评论回复）以及所有 TikTok 用户均可公开查看的评论（及评论回复）。
	// 默认值： ALL
	Status enum.CommentStatus `json:"status,omitempty"`
	// SortField 对评论（及评论回复）进行排序的特定字段。
	// 枚举值：
	// likes：按sort_order指定的排序方法，同时依据点赞数对评论（及评论回复）排序。
	// replies：按sort_order指定的排序方法，同时依据回复数对评论（及评论回复）排序。
	// create_time：按sort_order指定的排序方法，同时依据评论时间对评论（及评论回复）排序。
	// 评论默认以随机顺序返回。
	SortField string `json:"sort_field,omitempty"`
	// SortType 指定评论（及评论回复）排序方法的特定字段。
	// 枚举值：
	// asc：参照sort_field指定的排序依据，对评论（及评论回复）升序排序。
	// desc：参照sort_field指定的排序依据，对评论（及评论回复）降序排序。
	// smart：利用算法自动进行评论（及评论回复）排序。
	// 默认值： smart
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Cursor 分页光标。如果返回参数has_more 为true，请将返回中此参数的光标值需传递给后续请求，以获取下一页数据。
	// 默认值：0
	Cursor int64 `json:"cursor,omitempty"`
	// MaxCount 每页返回的最大评论数。
	// 默认值：20。
	// 最小值：1。
	// 最大值：30。
	// 注意: 受信任与安全政策影响，接口可能会返回少于 max_count 的评论数，即使返回参数has_more 为 true
	MaxCount int `json:"max_count,omitempty"`
}

// Encode implements GetRequest
func (r *ReplyListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("video_id", r.VideoID)
	values.Set("comment_id", r.CommentID)
	if r.Status != "" {
		values.Set("status", string(r.Status))
	}
	if r.SortField != "" {
		values.Set("sort_field", r.SortField)
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.FormatInt(r.Cursor, 10))
	}
	if r.MaxCount > 0 {
		values.Set("max_count", strconv.Itoa(r.MaxCount))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}
