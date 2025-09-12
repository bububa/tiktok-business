package changelog

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskCreateRequest 创建更改日志下载任务 API Request
type TaskCreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StartTime 查询起始时间，格式为 %Y-%m-%d %H:%M:%S
	StartTime string `json:"start_time,omitempty"`
	// EndTime 查询结束时间，格式为 %Y-%m-%d %H:%M:%S
	// 注意: 允许的最长时间间隔为30天，请确保end_time在Tstart_time后的30天内。
	EndTime string `json:"end_time,omitempty"`
	// Timezone 返回结果时区，枚举值见附录-时区信息。
	Timezone string `json:"timezone,omitempty"`
	// Module 查询模块，枚举值: BIDDING_AND_OPTIMIZATION,BUDGET,STATUS,TARGETING。
	Module string `json:"module,omitempty"`
	// ObjectIDs 查询对象ID
	ObjectIDs []string `json:"object_ids,omitempty"`
	// ObjectType 查询对象类型，枚举值: AD,ADGROUP,ADVERTISER,CAMPAIGN,INSTANT_FORM。
	ObjectType string `json:"object_type,omitempty"`
	// OperationTypes 操作类型。枚举值：CREATE, AUDIT, STATUS, UPDATE, DELETE, DOWNLOAD_LEADS, SUBSCRIBE_FORM, UNSUBSCRIBE_FORM
	OperationTypes []string `json:"operation_types,omitempty"`
	// OrderFields 排序字段。
	OrderFields []string `json:"order_fields,omitempty"`
}

// Encode implements PostRequest interface
func (r TaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TaskCreateResponse 创建更改日志下载任务 API Response
type TaskCreateResponse struct {
	model.BaseResponse
	Data struct {
		// TaskID 任务ID
		TaskID string `json:"task_id,omitempty"`
	} `json:"data"`
}
