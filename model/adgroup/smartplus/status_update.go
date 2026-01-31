package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/adgroup"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest Update the operation statuses of Upgraded Smart+ Ad Groups API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupIDs The IDs of the ad groups to operate on.
	// Size range: 1-20.
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// OperationStatus The operation to perform on the ad groups.
	// Enum values: DISABLE (pause), ENABLE (enable), DELETE (delete).
	// Note: The status of deleted ad groups or ad groups in deleted campaigns cannot be modified
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
}

// Encode implements PostRequest interface
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse Update the operation statuses of Upgraded Smart+ Ad Groups API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *StatusUpdateResult `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// AdgroupList A list of ad groups for which the operation was successful.
	AdgroupList []UpdatedAdgroup `json:"adgroup_list,omitempty"`
	// ErrorList A list of errors and the corresponding failed ad groups.
	ErrorList []adgroup.ErrUpdate `json:"error_list,omitempty"`
}

type UpdatedAdgroup struct {
	// AdgroupID The ID of the ad group for which the operation was successful
	AdgroupID string `json:"adgroup_id,omitempty"`
	// Status The current status of the ad group for which the operation was successful.
	// Enum values: DISABLE (paused), ENABLE (enabled), DELETE (deleted).
	Status enum.OperationStatus `json:"status,omitempty"`
}
