package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// AssetGroupUpdateRequest 更新资产组 API Request
type AssetGroupUpdateRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetGroupID 您想要更新的资产组的ID
	AssetGroupID string `json:"asset_group_id,omitempty"`
	// UpdateEntity 您想要更新的对象。枚举值：ASSET，MEMBER，NAME
	UpdateEntity enum.AssetGroupEntity `json:"update_entity,omitempty"`
	// Action 进行的操作。枚举值： ADD, DELETE。当update_entity为MEMBER时必填
	Action enum.AssetGroupUpdateAction `json:"action,omitempty"`
	// AssetGroupName 资产组名称
	AssetGroupName string `json:"asset_group_name,omitempty"`
	// Assets 您想要添加到资产组的资产。最多可传入200个资产ID
	Assets []Asset `json:"assets,omitempty"`
	// Members 您希望添加到资产组的商务中心成员。最多可传入50个ID
	Members []AssetGroupMember `json:"members,omitempty"`
}

// Encode implements PostRequest
func (r *AssetGroupUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
