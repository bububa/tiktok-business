package portfolio

import "github.com/bububa/tiktok-business/util"

// DeleteRequest 删除素材包 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CreativePortfolioIDs 要删除的创意素材包 ID 列表。
	// 最大数量: 100。
	// 若想获取广告账户下的创意素材包，可使用/creative/portfolio/list/。
	// 注意：如果存在任一无效的创意素材包 ID，删除操作将全部失败。无效的创意素材包 ID 包括：
	// 不存在的创意素材包 ID。
	// 已删除的创意素材包的 ID。
	CreativePortfolioIDs []string `json:"creative_portfolio_ids,omitempty"`
}

// Encode implements PostRequest
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
