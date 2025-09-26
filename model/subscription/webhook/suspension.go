package webhook

import "github.com/bububa/tiktok-business/model"

// Suspension 不同商务中心的广告账户的暂时停用状态
type Suspension struct {
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告账户名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// OwnerBcID 拥有该广告账户的商务中心的 ID
	OwnerBcID string `json:"owner_bc_id,omitempty"`
	// Status 发送 Webhook 通知时该广告账户的状态。
	// 枚举值：
	// STATUS_ENABLE: 解除暂时停用状态。
	// STATUS_LIMIT: 暂时停用。
	// 您接收到 Webhook 通知时，广告账户状态可能已发生变更。若想获取广告账户的最新状态，可使用/advertiser/info/并检查返回的 status。
	Status string `json:"status,omitempty"`
	// StatusChangeTime 广告账户暂时停用或解除暂时停用状态的时间，格式为 YYY-MM-DD HH:MM:SS（UTC+0）。
	// 示例：2025-01-01 00:00:01。
	StatusChangeTime model.DateTime `json:"status_change_time,omitzero"`
	// SuspensionReason 仅当 status 为 STATUS_LIMIT 时返回本字段。
	// 广告账户暂时停用的原因。
	SuspensionReason string `json:"suspension_reasion,omitempty"`
}
