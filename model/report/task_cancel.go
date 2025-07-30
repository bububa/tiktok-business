package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskCancelRequest 取消异步报表任务 API Request
type TaskCancelRequest struct {
	// AdvertiserID 广告主 ID。
	// 注意：若您调用/report/task/create/ （POST方法）时传入了advertiser_ids，您需将本字段设置为advertiser_ids中指定的广告主ID其中之一。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID 队列中或处理中的异步报表任务的 ID。
	// 若想获取满足条件的任务 ID，需先使用/report/task/create/创建异步报表任务，随后使用/report/task/check/查询任务状态，确保任务的status为 QUEUING或 PROCESSING。
	TaskID string `json:"task_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *TaskCancelRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TaskCancelResponse 取消异步报表任务 API Response
type TaskCancelResponse struct {
	model.BaseResponse
	Data struct {
		// Status 任务状态。
		// 枚举值：CANCELED（任务已取消）。
		// 若任务成功取消，本字段值将为CANCELED。
		// 若任务取消失败，将出现报错。
		Status enum.ReportTaskStatus `json:"status,omitempty"`
	} `json:"data"`
}
