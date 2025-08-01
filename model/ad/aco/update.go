package aco

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新智能创意广告素材 API Request
type UpdateRequest struct {
	Ad
	// PatchUpdate 是否使用增量更新模式。
	// 例如，若想增量更新广告文案，可仅设置以下参数：
	// patch_update 设置为 true 。
	// title_list设置为新的广告文案。
	// 传入其他必填参数，包括 Header 参数， advertiser_id 和 adgroup_id。
	PatchUpdate bool `json:"patch_update,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新智能创意广告素材 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *Ad `json:"data,omitempty"`
}
