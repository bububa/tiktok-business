package library

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TransferRequest 迁移线索到商务中心 API Request
type TransferRequest struct {
	// AdvertiserID 拥有这些线索的广告账号ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BcID 你想要迁移到的商务中心ID
	BcID string `json:"bc_id,omitempty"`
}

// Encode implements PostRequest
func (r *TransferRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TransferResponse 迁移线索到商务中心 API Response
type TransferResponse struct {
	model.BaseResponse
	Data struct {
		// LibraryID 生成的表单库ID，用来存储广告账户的线索
		LibraryID string `json:"library_id,omitempty"`
	} `json:"data"`
}
