package campaign

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新推广系列 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignID 推广系列 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：
	// 在推广系列商品来源为商品库和推广系列商品来源非商品库的购物广告中使用特殊广告分类目前为单独的白名单功能。如需使用这些功能，请联系您的TikTok销售代表。
	// 本字段对于注册地为美国或加拿大的广告主已全量。若您是注册地非美国或加拿大的广告主且想在定向美国或加拿大的广告中使用特殊广告分类，您需申请额外的白名单。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// Budget 推广系列预算
	Budget float64 `json:"budget,omitempty"`
	// PONumber PO（采购订单）号。
	// PO 号可用于结算单追踪和对账。
	// 了解月度结算单上的 PO 号。
	PONumber string `json:"po_number,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新推广系列 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Campaign `json:"data,omitempty"`
}
