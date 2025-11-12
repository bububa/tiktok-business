package bc

import "github.com/bububa/tiktok-business/util"

// AssetGroupDeleteRequest 删除资产组 API Request
type AssetGroupDeleteRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetGroupIDs 您想要删除的资产组的ID。目前，您每次只能传入一个ID
	AssetGroupIDs []string `json:"asset_group_ids,omitempty"`
}

// Encode implements PostRequest
func (r *AssetGroupDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
