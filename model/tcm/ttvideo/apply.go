package ttvideo

import "github.com/bububa/tiktok-business/util"

// ApplyRequest 申请 Spark Ads 授权 API Request
type ApplyRequest struct {
	// VideoID 在 TTCM 账户下创建的订单的 ID。
	// 视频需上传且标记为品牌推广内容（branded content），且创作者需已将通过/tcm/campaign/code/create/创建的推广系列代码与视频绑定。您可通过/tcm/order/get/接口获取 video_id。
	VideoID string `json:"video_id,omitempty"`
	// TCMAccountID TikTok Creator Marketplace 账户 ID。
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// AuthorizedDays 授权的有效期（按天计）。
	// 枚举值：7、30、60、365。默认值：30。
	// 注意：本字段可在申请授权时指定授权的有效期，或用于申请更新已批准授权的有效期。
	// 实际有效期可能不同于您申请的有效期，因为创作者在批准申请时可自行选择任意有效期（7, 30, 60, 365）。
	AuthorizedDays int `json:"authorized_days,omitempty"`
}

// Encode implements PostRequest
func (r *ApplyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
