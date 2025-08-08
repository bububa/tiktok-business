package aco

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest 更新智能创意广告素材状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdGroupID 广告组 ID。
	// 注意：传入已开启智能创意的广告组的 ID。该广告组的 creative_material_mode 需为 SMART_CREATIVE。
	AdGroupID string `json:"ad_group_id,omitempty"`
	// MaterialIDs 素材 ID。您只能传入一个素材 ID
	MaterialIDs []string `json:"material_ids,omitempty"`
	// MaterialStatus 素材状态。枚举值： ENABLE （素材可使用）， DISABLE（素材不可使用）
	MaterialStatus enum.EnableDisable `json:"MaterialStatus,omitempty"`
}

// Encode implements PostRequest interface
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 更新智能创意广告素材状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *StatusUpdateResult `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// AdGroupID 广告组 ID。
	AdGroupID string `json:"ad_group_id,omitempty"`
	// MaterialIDs 素材ID
	MaterialIDs []string `json:"material_ids,omitempty"`
	// MaterialStatus 素材状态。枚举值： ENABLE (素材可使用)， DISABLE(素材不可使用)
	MaterialStatus []enum.EnableDisable `json:"material_status,omitempty"`
}
