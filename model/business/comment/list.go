package comment

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取自有视频的评论 API Request
type ListRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// VideoID 自有 TikTok 视频的唯一标识符，用于列出对应评论（及评论回复）。可通过/business/video/list/ 接口的item_id字段获取
	VideoID string `json:"video_id,omitempty"`
	// CommentIDs 评论或评论回复对应的ID列表，用于筛选返回结果。
	// 最大数量：30
	CommentIDs []string `json:"comment_ids,omitempty"`
	// IncludeReplies 是否在结果中返回第一层级评论对应的评论回复。
	// 若将本字段设置为true，将返回每条第一层级评论的对应回复，每条第一层级评论返回的回复数量最大为3，评论和回复默认根据智能排序算法进行排序。
	// 默认值：false。
	// 若想获取单条评论的所有回复，需使用/business/comment/reply/list/。
	IncludeReplies bool `json:"include_replies,omitempty"`
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
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("video_id", r.VideoID)
	if len(r.CommentIDs) > 0 {
		values.Set("comment_ids", string(util.JSONMarshal(r.CommentIDs)))
	}
	if r.IncludeReplies {
		values.Set("include_replies", "true")
	}
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

// ListResponse 获取自有视频的评论 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// Cursor 用于获取下一页结果的光标，可传递到后续接口请求的 cursor 查询参数
	Cursor int64 `json:"cursor,omitempty"`
	// HasMore 是否有下一页数据
	HasMore bool `json:"has_more,omitempty"`
	// Comments 针对某一视频的当前页面评论（或评论回复或评论和评论回复）列表
	Comments []Comment `json:"comments,omitempty"`
}

// Comment 评论
type Comment struct {
	// CommentID 评论或评论回复的唯一标识符
	CommentID string `json:"comment_id,omitempty"`
	// VideoID 评论或评论回复对应的自有视频的唯一标识符
	VideoID string `json:"video_id,omitempty"`
	// UserID 发表评论或评论回复的用户在开发者应用和 TikTok 账户范围内的唯一标识符。
	// 示例 :ce9b7be629b0447af777a0ad945783f61b7f790982b49fbda04a23b4722eaffe
	// 注意：本字段将在下个 API 版本中废弃。推荐您使用 unique_identifier 识别 TikTok 用户，因为该用户标识符在多个 API 之间保持一致
	UserID string `json:"user_id,omitempty"`
	// UniqueIdentifier 发表评论或评论回复的用户的全球统一标识符。该标识符在多个 API 之间保持一致，方便集成和参数的互相参照。
	// 示例：+ABc1D2/E0fGhijklMNOrD34vDW/zDiZUNRH5Jyahue6OcA7b8ZpwnqXV9u1Uh+k
	UniqueIdentifier string `json:"unique_identifier,omitempty"`
	// CreateTime 评论或评论回复发表的日期和时间，采用 Unix/Epoch 格式
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// Text 评论或评论回复的文本内容
	Text string `json:"text,omitempty"`
	// Likes 评论或评论回复获得的点赞数
	Likes int64 `json:"likes,omitempty"`
	// Replies 评论获得的回复数
	Replies int64 `json:"replies,omitempty"`
	// Owner 评论或评论回复是否由发布对应视频的用户本人创建
	Owner bool `json:"owner,omitempty"`
	// Liked 发布对应视频的用户是否点赞了此评论或评论回复
	Liked bool `json:"liked,omitempty"`
	// Pinned 发布对应视频的用户是否置顶了此评论或评论回复
	Pinned bool `json:"pinned,omitempty"`
	// Status 此评论或评论回复的可见状态。
	// 枚举值：
	// HIDDEN：评论或评论回复已隐藏，无法公开浏览
	// PUBLIC：评论或评论回复可被所有用户查看
	// 示例：PUBLIC
	Status enum.CommentStatus `json:"status,omitempty"`
	// Username 发布此评论或评论回复的 TikTok 账户的唯一用户名
	Username string `json:"username,omitempty"`
	// DisplayName 发布此评论或评论回复的 TikTok 账户的显示名称
	DisplayName string `json:"display_name,omitempty"`
	// ProfileImage 发布此评论或评论回复的 TikTok 账户主页头像的临时URL。过期日期和时间以 Epoch/Unix 时间格式包含在x-expires查询参数中，以秒为单位
	ProfileImage string `json:"profile_image,omitempty"`
	// ParentCommentID 仅为评论回复返回此字段。
	// 收到此条评论回复的评论的ID
	ParentCommentID string `json:"parent_comment_id,omitempty"`
	// ReplyList 仅当请求参数include_replies设置为true时，为第一层级评论返回此对象数组。
	// 有关该评论收到的回复的信息。
	// 最大数量：3
	ReplyList []Comment `json:"reply_list,omitempty"`
}
