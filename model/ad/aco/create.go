package aco

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建智能创意广告 API Request
type CreateRequest Ad

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建智能创意广告 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Ad `json:"data,omitempty"`
}
