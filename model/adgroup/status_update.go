package adgroup

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest 更新广告组状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupIDs 广告组ID
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// OperationStatus 要进行的操作。
	// 枚举值：DISABLE（暂停），ENABLE（启用），DELETE（删除）。
	// 注意：
	// 您无法修改已删除的广告组，或已删除的推广系列中的广告组的状态。
	// 不支持对覆盖和频次广告组执行 ENABLE（启用）和DISABLE（暂停）操作。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// AllowPartialSuccess 是否支持操作部分成功。
	// 枚举值：true，false。 默认值：false。
	// 若本字段设置为 true 且您指定了多个广告组，则支持仅对部分广告组操作成功，此时将筛除不支持修改状态的广告组。
	AllowPartialSuccess bool `json:"allow_partial_success,omitempty"`
}

// Encode implements PostRequest interface
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 更新广告组状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *StatusUpdateResult `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// AdgroupIDs 操作成功的广告组 ID 列表
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// Status 操作成功的广告组当前的状态。
	// 枚举值：DISABLE（已暂停），ENABLE（已启用），DELETE（已删除）。
	Status enum.OperationStatus `json:"status,omitempty"`
	// ErrorList 仅当 allow_partial_success 设置为 true 时返回本字段。
	// 报错列表及对应的操作失败的广告组。
	ErrorList []ErrUpdate `json:"error_list,omitempty"`
}

type ErrUpdate struct {
	// AdgroupID 操作失败的广告组的 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// ErrorMessage 与该操作失败的广告组关联的报错信息
	ErrorMessage string `json:"error_message,omitempty"`
}

func (e ErrUpdate) Error() string {
	return util.StringsJoin(e.ErrorMessage, "[adgroup_id:", e.AdgroupID, "]")
}
