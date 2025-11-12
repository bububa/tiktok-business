package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// MemberUpdateRequest 修改成员信息 API Request
type MemberUpdateRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// UserID 用户ID
	UserID string `json:"user_id,omitempty"`
	// UserRole 用户在商务中心的基础角色。
	// 枚举值：
	// ADMIN：管理员可以访问商务中心内的所有功能。
	// STANDARD：标准成员只能对分配给他们的账户进行操作。
	UserRole enum.UserRole `json:"user_role,omitempty"`
	// UserName 新用户名
	UserName string `json:"user_name,omitempty"`
	// ExtUserRole 用户在商务中心除基本角色外的扩展角色
	ExtUserRole *ExtUserRole `json:"ext_user_role,omitempty"`
}

// Encode implements PostRequest
func (r *MemberUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
