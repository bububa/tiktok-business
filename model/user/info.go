package user

import "github.com/bububa/tiktok-business/model"

// InfoResponse 获取用户信息 API Response
type InfoResponse struct {
	model.BaseResponse
	Data *User `json:"data,omitempty"`
}

// User 用户信息
type User struct {
	// DisplayName 用户昵称
	DisplayName string `json:"display_name,omitempty"`
	// Email 电子邮箱。
	// 注意：电子邮箱已作数据脱敏处理。
	Email string `json:"email,omitempty"`
	// CoreUserID 应用开发者在TikTok for Business平台的用户 ID
	CoreUserID string `json:"core_user_id,omitempty"`
	// CreateTime 用户创建时间
	CreateTime int64 `json:"create_time,omitempty"`
	// AvatarURL 用户头像的URL
	AvatarURL string `json:"avatar_url,omitempty"`
}
