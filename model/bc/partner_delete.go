package bc

import (
	"github.com/bububa/tiktok-business/util"
)

// PartnerDeleteRequest 删除合作伙伴 API Request
type PartnerDeleteRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// PartnerID 要删除的合作伙伴的商务中心的 ID
	PartnerID string `json:"partner_id,omitempty"`
}

// Encode implements PostRequest
func (r *PartnerDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
