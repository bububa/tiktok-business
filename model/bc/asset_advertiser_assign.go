package bc

import "github.com/bububa/tiktok-business/util"

// AssetAdvertiserAssignRequest 在商务中心内将 TikTok 账号和广告账号绑定 API Request
type AssetAdvertiserAssignRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetID 商务中心内某个 TikTok 账号的 ID。
	// 如需获取商务中心内的 TikTok 账号列表，使用 /bc/asset/get/ 接口，并将 asset_type 设为 TT_ACCOUNT
	AssetID string `json:"asset_id,omitempty"`
	// AdvertiserID 要与该 TikTok 账号绑定的广告账号的 ID，广告账号需在同一商务中心中。
	// 如需获取商务中心内的广告账号列表，使用 /bc/asset/get/ 接口，并将 asset_type 设为 ADVERTISER
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest
func (r *AssetAdvertiserAssignRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
