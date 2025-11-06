package video

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SettingsRequest 获取 TikTok 账号帖子隐私设置 API Request
type SettingsRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
}

// Encode implements GetRequest
func (r *SettingsRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SettingsResponse 获取 TikTok 账号帖子隐私设置 API Response
type SettingsResponse struct {
	model.BaseResponse
	Data *Settings `json:"data,omitempty"`
}

// Settings TikTok 账号帖子隐私设置
type Settings struct {
	// PrivacyLevelOptions 该 TikTok 账号可选的隐私级别设置 。
	// 若该账号为公开账号，支持的枚举值为：
	// PUBLIC_TO_EVERYONE：所有人。
	// MUTUAL_FOLLOW_FRIENDS：好友（您回关的粉丝）。
	// SELF_ONLY：仅自己。
	// 若该账号为私密账号，支持的枚举值为：
	// FOLLOWER_OF_CREATOR：粉丝。
	// MUTUAL_FOLLOW_FRIENDS：好友（您回关的粉丝）。
	// SELF_ONLY：仅自己。
	PrivacyLevelOptions []enum.PrivacyLevelOption `json:"privacy_level_options,omitempty"`
	// CommentDisabled 该 TikTok 账号的帖子是否可以使用“允许发表评论”设置。
	// 支持的值：
	// true：不可以。不支持为该 TikTok 账号的帖子启用此设置。
	// 若您已在 TikTok 账号的评论隐私设置中将“允许评论”设置为“不允许”，本字段将为 true。
	// false：可以。支持为该 TikTok 账号的帖子启用或关闭此设置。
	CommentDisabled bool `json:"comment_disabled,omitempty"`
	// DuetDisabled 该 TikTok 账号的帖子是否可以使用“允许合拍”设置。
	// 支持的值：
	// true：不可以。不支持为该 TikTok 账号的帖子启用此设置。
	// 若您已在 TikTok 账号的合拍隐私设置中将“允许合拍”设置为“仅自己”或 TikTok 账号为私密账号，本字段将为 true。
	// false：可以。支持为该 TikTok 账号的帖子启用或关闭此设置。
	DuetDisabled bool `json:"duet_disabled,omitempty"`
	// StitchDisabled 该 TikTok 账号的帖子是否可以使用“允许拼接”设置。
	// 支持的值：
	// true：不可以。不支持为该 TikTok 账号的帖子启用此设置。
	// 若您已在 TikTok 账号的拼接隐私设置中将“允许拼接”设置为“仅自己”或 TikTok 账号为私密账号，本字段将为 true。
	// false：可以。支持为该 TikTok 账号的帖子启用或关闭此设置。
	StitchDisabled bool `json:"stitch_disabled,omitempty"`
	// MaxVideoPostDurationSec 该 TikTok 账号能否发布的视频帖子最大时长，单位为秒。
	// 若您试图通过/business/video/publish/发布超过最大时长的视频帖子，发布将失败。
	MaxVideoPostDurationSec int64 `json:"max_video_post_duration_sec,omitempty"`
}
