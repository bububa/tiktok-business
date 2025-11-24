package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetGroupCreateRequest 创建资产组 API Request
type AssetGroupCreateRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetGroupName 资产组名称
	AssetGroupName string `json:"asset_group_name,omitempty"`
	// Assets 您想要添加到资产组的资产。最多可传入200个资产ID
	Assets []Asset `json:"assets,omitempty"`
	// Members 您希望添加到资产组的商务中心成员。最多可传入50个ID
	Members []AssetGroupMember `json:"members,omitempty"`
}

// Encode implements PostRequest
func (r *AssetGroupCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AssetGroupCreateResponse 创建资产组 API Response
type AssetGroupCreateResponse struct {
	model.BaseResponse
	Data struct {
		// AssetGroupID 您创建的资产组的ID
		AssetGroupID string `json:"asset_group_id,omitempty"`
	} `json:"data"`
}
