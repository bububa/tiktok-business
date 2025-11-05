package ttvideo

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusRequest 获取授权状态 API Request
type StatusRequest struct {
	// VideoID 创作者账户中创建的视频的 ID。
	// 视频需上传且标记为品牌推广内容，且创作者需已通过/tcm/order/create/v2/生成的邀请链接（invite_link）或推广系列代码（campaign_code）将视频与 TikTok Creator Marketplace（TTCM）订单绑定。
	// 您可通过/tcm/order/get/v2/ 或 /tcm/report/get/v2/接口获取 video_id。
	VideoID string `json:"video_id,omitempty"`
	// TCMAccountID TikTok Creator Marketplace 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
}

// Encode implements GetRequest
func (r *StatusRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("video_id", r.VideoID)
	values.Set("tcm_account_id", r.TCMAccountID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// StatusResponse 获取授权状态 API Response
type StatusResponse struct {
	model.BaseResponse
	Data *Video `json:"data,omitempty"`
}

type Video struct {
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// AuthStatus 授权状态。枚举值：
	// WAITING：创作者尚未批准或拒绝您的授权申请。
	// 创作者在 TikTok App 中收到您的授权申请通知后，可以在 App 中的 System Notification 界面选择批准或拒绝您的授权申请。若创作者在"Branded content and ad"界面启用了"Ad promotion"（广告推广）设置，auth_status 仍将显示为WAITING，表示客户尚未批准您的申请。不过，此时您将可以获取到auth_code，从而可以将视频授权给广告主账号。
	// REJECTED：创作者已拒绝您的授权申请。
	// ACCEPTED：创作者已批准您的授权申请。
	AuthStatus enum.TTVideoAuthStatus `json:"auth_status,omitempty"`
	// AuthStatusUpdateTime auth_status（授权状态）最后一次更新的时间。例如，auth_status从WAITING变为ACCEPTED的时间。
	AuthStatusUpdateTime model.DateTime `json:"auth_status_update_time,omitzero"`
	// NumRemainingRequest 请求Spark Ads授权的剩余次数。在授权请求被拒绝或批准前，您最多剩余30次请求授权的机会
	NumRemainingRequest int `json:"num_remaining_request,omitempty"`
	// RequestedAuthorizationDays 您申请授权时指定授权的有效期（按天计），或申请更新已批准授权的有效期。枚举值: 7, 30, 60, 365。
	// 默认值：30。
	// 注意：实际有效期(auth_code_end_time - auth_code_start_time)可能不同于您请求的有效期，因为创作者在批准申请时可自行选择任意有效期（7, 30, 60, 365）。
	RequestedAuthorizationDays int `json:"requested_authorization_days,omitempty"`
	// AuthCodeStatus 授权码状态。枚举值：
	// NOT_USED：授权码已生成，但尚未使用。
	// IN_USE：授权码已生成，且已经使用。
	// EXPIRED：授权码已生成，但已经失效。
	// DELETED：授权码已删除。
	AuthCodeStatus enum.TCMAuthCodeStatus `json:"auth_code_status,omitempty"`
	// AuthCodeStartTime 授权码生效的UTC+0日期及时间
	AuthCodeStartTime model.DateTime `json:"auth_code_start_time,omitzero"`
	// AuthCodeEndeTime 授权码失效的UTC+0日期及时间
	AuthCodeEndTime model.DateTime `json:"auth_code_end_time,omitzero"`
	// AuthCode 在订单中推广视频的授权码
	AuthCode string `json:"auth_code,omitempty"`
}
