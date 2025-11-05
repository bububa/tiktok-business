package showcase

import "github.com/bububa/tiktok-business/enum"

// Identity 橱窗权限的认证身份
type Identity struct {
	// IdentityID 与该橱窗绑定的认证身份ID
	IdentityID string `json:"identity_id,omitempty"`
	// IdentityType 认证身份类型。
	// 枚举值：TT_USER （TikTok用户）， BC_AUTH_TT（添加到商务中心的TikTok用户）。
	// 您可以查看认证身份了解不同认证身份类型。
	IdentityType enum.IdentityType `json:"identity_type,omitempty"`
	// IdentityAuthorizedBcID identity_type 为 BC_AUTH_TT时返回。
	// 与该橱窗绑定的商务中心的ID。该橱窗同时与添加到商务中心的TikTok用户这一认证身份绑定。
	IdentityAuthorizedBcID string `json:"identity_authorized_bc_id,omitempty"`
}
