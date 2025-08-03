package spc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建 Smart+ 推广系列 API Request
type CreateRequest Campaign

// Encode implements PostRequest interface
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建 Smart+ 推广系列 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
