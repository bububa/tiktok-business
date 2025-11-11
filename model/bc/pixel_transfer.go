package bc

import "github.com/bububa/tiktok-business/util"

// PixelTransferRequest 将pixel转移至商务中心 API Request
type PixelTransferRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// PixelCode 待转移的Pixel代码
	PixelCode string `json:"pixel_code,omitempty"`
}

// Encode implements PostRequest
func (r *PixelTransferRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
