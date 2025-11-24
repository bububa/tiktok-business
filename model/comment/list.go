package comment

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取评论列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CommentType 评论类型。
	// 枚举值：全部 ALL，原始评论 COMMENT，回复评论 REPLY。
	// 默认值：["ALL"]。
	CommentType []enum.CommentType `json:"comment_type,omitempty"`
	// SearchField 搜索字段。
	// 枚举值：
	// ADGROUP_ID：通过广告组 ID 筛选
	SearchField string `json:"search_field,omitempty"`
	// SearchValue 搜索值。
	// search_field 为 ADGROUP_ID 时，需通过本字段指定一个广告组 ID。
	// 若想获取广告账户中的广告组列表，可使用/adgroup/get/
	SearchValue string `json:"search_value,omitempty"`
	// SortField 排序字段。
	// 枚举值: 创建时间 CREATE_TIME，点赞数 LIKES，回复数 REPLIES。
	// 默认值：CREATE_TIME
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序类型。
	// 枚举值：ASC （升序）, DESC（降序）
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// StartTime 请求起始时间。
	// 格式：YYYY-MM-DD
	StartTime string `json:"start_time,omitempty"`
	// EndTime 请求结束时间。
	// 格式：YYYY-MM-DD
	EndTime string `json:"end_time,omitempty"`
	// Page 当前页数
	// 默认值：1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10。
	// 取值范围：1-100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.CommentType) > 0 {
		values.Set("comment_type", string(util.JSONMarshal(r.CommentType)))
	}
	values.Set("search_field", r.SearchField)
	values.Set("search_value", r.SearchValue)
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if r.SortField != "" {
		values.Set("sort_field", r.SortField)
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
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

// ListResponse 获取评论列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Comments 评论列表
	Comments []Comment `json:"comments,omitempty"`
}

// Comment 评论
type Comment struct {
	// CommentID 评论ID
	CommentID string `json:"comment_id,omitempty"`
	// AppID 应用 ID
	AppID string `json:"app_id,omitempty"`
	// Content 评论内容
	Content string `json:"content,omitempty"`
	// Likes 评论或评论回复获得的点赞数
	Likes int64 `json:"likes,omitempty"`
	// Replies 评论获得的回复数
	Replies int64 `json:"replies,omitempty"`
	// CommentType 评论类型。枚举值：原始评论 COMMENT，回复评论 REPLY
	CommentType enum.CommentType `json:"comment_type,omitempty"`
	// OriginalCommentID 原始评论ID。当 comment_type=REPLY 时有值
	OriginalCommentID string `json:"original_comment_id,omitempty"`
	// ReplyToCommentID 回复的原始评论 ID
	ReplyToCommentID string `json:"reply_to_comment_id,omitempty"`
	// CommentStatus 此评论或评论回复的可见状态。
	// 枚举值：
	// HIDDEN：评论或评论回复已隐藏，无法公开浏览
	// PUBLIC：评论或评论回复可被所有用户查看
	// 示例：PUBLIC
	CommentStatus enum.CommentStatus `json:"comment_status,omitempty"`
	// HitBlockedword 是否命中屏蔽词
	HitBlockedword bool `json:"hit_blockedword,omitempty"`
	// AdText 广告标题
	AdText string `json:"ad_text,omitempty"`
	// CreateTime 评论或评论回复发表的日期和时间，采用 Unix/Epoch 格式
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// CampaignID 评论所属推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 评论所属推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// AdgroupID 评论所属广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 评论所属广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// AdID 评论所属广告 ID
	AdID string `json:"ad_id,omitempty"`
	// AdName 评论所属广告名称
	AdName string `json:"ad_name,omitempty"`
	// TikTokItemID TikTok视频 ID
	TikTokItemID string `json:"tiktok_item_id,omitempty"`
	// IdentityID 认证身份ID，当使用Spark Ads(tiktok_item_id)时必填。如果您不使用Spark Ads，我们仍然强烈建议传入 ideneity_id 和 identity_type (CUSTOMIZED_USER类型)，以便更好地管理广告信息
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型，当使用Spark Ads(tiktok_item_id)时必填。枚举值：CUSTOMIZED_USER (自定义用户）， TT_USER (TikTok用户）。关于认证身份的更多信息，请访问认证身份
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IsPinned 评论是否置顶
	IsPinned bool `json:"is_pinned,omitempty"`
	// CanDelete 评论是否可删除
	CanDelete bool `json:"can_delete,omitempty"`
	// IsAuthTtba 用户是否已关联TTBA账户
	IsAuthTtba bool `json:"is_auth_ttba,omitempty"`
	// IsAuthCommentManageScope 用户是否有评论管理权限
	IsAuthCommentManageScope bool `json:"is_auth_comment_manage_scope,omitempty"`
	// VideoPlayURL 视频播放 URL
	VideoPlayURL string `json:"video_play_url,omitempty"`
	// VideoCoverURL 视频封面 URL
	VideoCoverURL string `json:"video_cover_url,omitempty"`
	// UserAvatarURL 用户头像 URL
	UserAvatarURL string `json:"user_avatar_url,omitempty"`
	// UserName 用户名
	UserName string `json:"user_name,omitempty"`
	// UserID 发表评论的 TikTok 用户的 ID
	UserID string `json:"user_id,omitempty"`
}
