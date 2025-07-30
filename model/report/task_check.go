package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskCheckRequest 获取异步报表任务状态 API Request
type TaskCheckRequest struct {
	// AdvertiserID 广告主 ID。
	// 注意：若您调用/report/task/create/ （POST方法）时传入了advertiser_ids，您需将本字段设置为advertiser_ids中指定的广告主ID其中之一。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID 异步报表任务ID。您可在/report/task/create/ 接口中的返回数据中获得该字段。
	TaskID string `json:"task_id,omitempty"`
}

// Encode implements GetRequest interface
func (r TaskCheckRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("task_id", r.TaskID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// TaskCheckResponse 获取异步报表任务状态 API Response
type TaskCheckResponse struct {
	model.BaseResponse
	Data *TaskCheckResult `json:"data,omitempty"`
}

type TaskCheckResult struct {
	// Status 任务状态。
	// 枚举值：QUEUING（队列中），PROCESSING（ 处理中），SUCCESS（ 任务成功），FAILED（任务失败），CANCELED（任务已取消）。
	Status enum.ReportTaskStatus `json:"status,omitempty"`
	// Message 任务失败原因。当status为 FAILED时，该字段返回具体原因。
	Message string `json:"message,omitempty"`
}

func (r TaskCheckResult) IsError() bool {
	return r.Status == enum.ReportTaskStatus_FAILED
}

func (r TaskCheckResult) Error() string {
	return r.Message
}
