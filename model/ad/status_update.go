package ad

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest 更新广告状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdIDs 广告 ID 列表。允许数量范围： 1-20。您需传入ad_ids 或 aco_ad_ids。
	AdIDs []string `json:"ad_ids,omitempty"`
	// AcoAdIDs 程序化创意广告 ID 列表，此类型广告仅支持 ENABLE 和 DISABLE，允许数量范围： 1-20。您需传入ad_ids 或 aco_ad_ids。
	AcoAdIDs []string `json:"aco_ad_ids,omitempty"`
	// OperationStatus 要进行的操作。
	// 枚举值：DELETE（删除）,DISABLE（暂停）,ENABLE（启用）。
	// 注意：您无法修改已删除的广告，或已删除的推广系列或广告组中的广告的状态。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
}

// Encode implements PostRequest interface
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 更新广告状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *StatusUpdateResult `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// AdIDs 更新的广告 ID 列表
	AdIDs []string `json:"ad_ids,omitempty"`
	// AcoAdIDs 更新的程序化创意广告 ID 列表
	AcoAdIDs []string `json:"aco_ad_ids,omitempty"`
	// Status 执行状态，枚举值: DELETE(删除),DISABLE(关停),ENABLE(开启)。
	Status enum.OperationStatus `json:"status,omitempty"`
}
