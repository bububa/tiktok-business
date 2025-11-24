package portfolio

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 根据 ID 获取素材包 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CreativePortfolioID 创意素材包ID
	CreativePortfolioID string `json:"creative_portfolio_id,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("creative_portfolio_id", r.CreativePortfolioID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 根据 ID 获取素材包 API Response
type GetResponse struct {
	model.BaseResponse
	Data *Portfolio `json:"data,omitempty"`
}
