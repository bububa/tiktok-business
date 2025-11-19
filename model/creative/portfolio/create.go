package portfolio

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建素材包 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	Portfolio
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建素材包 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// CreativePortfolioID 创意素材包ID
		CreativePortfolioID string `json:"creative_portfolio_id,omitempty"`
	} `json:"data"`
}
