package smartplus

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest Update the operation statuses of Upgraded Smart+ Ads API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SmartPlusAdIDs A list of ad IDs.
	// Max size: 20.
	SmartPlusAdIDs []string `json:"smart_plus_ad_ids,omitempty"`
	// OperationStatus The operation to perform on the ads.
	// Enum values: DISABLE (pause), ENABLE (enable), DELETE (delete).
	// Note: The status of deleted ads or ads in deleted campaigns or ad groups cannot be modified.
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
}

// Encode implements PostRequest interface
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse Update the operation statuses of Upgraded Smart+ Ads API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *StatusUpdateResult `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// SmartPlusAdIDs 更新的广告 ID 列表
	SmartPlusAdIDs []string `json:"smart_plus_ad_ids,omitempty"`
	// Status The current statuses of the ads.
	// Enum values: DISABLE (paused), ENABLE (enabled), DELETE (deleted).
	Status enum.OperationStatus `json:"status,omitempty"`
}
