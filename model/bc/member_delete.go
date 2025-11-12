package bc

import (
	"github.com/bububa/tiktok-business/util"
)

// MemberDeleteRequest 删除成员或撤销成员邀请 API Request
type MemberDeleteRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// UserID 移除现有成员时必填。
	// 用户 ID。
	// 若想获取商务中心中现有成员的用户 ID，可使用 /bc/member/get/ 并查看返回的 user_id。
	UserID string `json:"user_id,omitempty"`
	// UserEmail 撤销成员邀请时必填。
	// 用户邮件地址。
	// 若想获取商务中心中受邀成员的用户邮件地址，可使用 /bc/member/get/ 并查看返回的 user_email。
	UserEmail string `json:"user_email,omitempty"`
}

// Encode implements PostRequest
func (r *MemberDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
