package post

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AuthorizeStatusRequest 获取 TikTok 帖子的授权状态 API Request
type AuthorizeStatusRequest struct {
	// BusinessID TikTok 账户的应用唯一识别 ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// ItemID TikTok 帖子 ID。
	// 若想获取 TikTok 账号的公开帖子，可使用/business/video/list/。
	// 示例：7123456789123456789。
	// 注意：若 TikTok 帖子无授权码，无论是未生成过授权码或生成的授权码已删除，将出现报错。
	ItemID string `json:"item_id,omitempty"`
}

// Encode implements GetRequest
func (r *AuthorizeStatusRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("item_id", r.ItemID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AuthorizeStatusResponse 获取 TikTok 帖子的授权状态 API Response
type AuthorizeStatusResponse struct {
	model.BaseResponse
	Data *AuthorizeStatus `json:"data,omitempty"`
}

// AuthorizeStatus TikTok 帖子的授权状态
type AuthorizeStatus struct {
	// ItemID TikTok 帖子 ID
	ItemID string `json:"item_id,omitempty"`
	// AuthCode TikTok 帖子的授权码
	AuthCode string `json:"auth_code,omitempty"`
	// AuthCodeStartTime 授权码生效时间，格式为 YYYY-MM-DD HH:MM:SS（UTC+0）
	AuthCodeStartTime model.DateTime `json:"auth_code_start_time,omitzero"`
	// AuthCodeEndTime 授权码过期时间，格式为 YYYY-MM-DD HH:MM:SS（UTC+0）
	AuthCodeEndTime model.DateTime `json:"auth_code_end_time,omitzero"`
	// AuthorizationDays 授权的有效期，单位为天
	AuthorizationDays int `json:"authorization_days,omitempty"`
	// AuthCodeStatus 授权码的状态。
	// 枚举值：
	// NOT_USED：授权码已生成，但尚未使用。
	// IN_USE：授权码已生成，且已经使用。
	// EXPIRED：授权码已生成，但已经失效。
	// DELETED：授权码已删除
	AuthCodeStatus enum.AuthCodeStatus `json:"auth_code_status,omitempty"`
}
