package lead

import (
	"net/http"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskRequest 创建线索下载任务 API Request
type TaskRequest struct {
	// AdvertiserID 需传入 advertiser_id 和 library_id 其中之一。
	// 若您尚未使用/page/library/transfer/将广告账户的线索迁移至商务中心，需指定 advertiser_id。
	// 若您已使用/page/library/transfer/将广告账户的线索迁移至商务中心，需指定 library_id。
	// 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// LibraryID 需传入 advertiser_id 和 library_id 其中之一。
	// 若您尚未使用/page/library/transfer/将广告账户的线索迁移至商务中心，需指定 advertiser_id。
	// 若您已使用/page/library/transfer/将广告账户的线索迁移至商务中心，需指定 library_id。
	// 表单库 ID。
	// 若想获取您有权访问的表单库 ID 列表，可使用/page/library/get/
	LibraryID string `json:"library_id,omitempty"`
	// AdID 需传入 ad_id 和 page_id 其中之一。
	// 若指定 ad_id，将为某个广告创建线索下载任务。
	// 若指定 page_id，将为某个即时表单（线索表单）创建线索下载任务。
	// 广告 ID。
	// 您需指定使用即时表单收集线索的广告。若想了解如何创建此类广告，可查看创建优化位置为即时表单的线索广告。
	// 若想获取广告账户中的广告 ID 列表，可使用 /ad/get/。
	AdID string `json:"ad_id,omitempty"`
	// PageID 需传入 ad_id 和 page_id 其中之一。
	// 若指定 ad_id，将为某个广告创建线索下载任务。
	// 若指定 page_id，将为某个即时表单（线索表单）创建线索下载任务。
	// 即时表单 ID。
	// 若想获取广告账户中的即时表单 ID 列表，可使用 /page/get/ 并将 business_type 设置为 LEAD_GEN
	PageID string `json:"page_id,omitempty"`
	// 任务 ID。
	// 若不传入该字段，该请求将创建一个线索下载任务。您会在返回中得到task_id。
	// 若传入该字段，该请求会进行轮询操作，检查任务是否完成 (status 为 SUCCEED)
	TaskID string `json:"task_id,omitempty"`
	// XLeadRegion 对于欧洲经济区/瑞士/英国市场的线索数据，使用URL https://business-api.tiktok.com/open_api/并指定 Header x-lead-region: eu
	// 对于美国市场的线索数据，选择以下 URL 和 Header 设置方式之一
	// 对于其他市场的线索数据，基础 URL 为https://business-api.tiktok.com/open_api/，无需指定 Header x-lead-region
	XLeadRegion string `json:"-"`
}

// Encode implements PostRequest
func (r *TaskRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PreRequest implements PreRequest
func (r *TaskRequest) PreRequest(httpReq *http.Request) error {
	if r.XLeadRegion == "" {
		return nil
	}
	httpReq.Header.Set("x-lead-region", r.XLeadRegion)
	return nil
}

// TaskResponse 创建线索下载任务 API Response
type TaskResponse struct {
	model.BaseResponse
	Data *Task `json:"data,omitempty"`
}

// Task 线索下载任务
type Task struct {
	// Status 任务状态。
	// 枚举值: CREATED,FAILED,RUNNING,SUCCEED
	Status enum.LeadTaskStatus `json:"status,omitempty"`
	// TaskID 任务 ID。
	// 您可以在请求中传入task_id进行轮询操作，检查任务是否完成 (status = SUCCEED)
	TaskID string `json:"task_id,omitempty"`
	//  FileName 仅当 status 为 SUCCEED 时返回。
	// 线索文件的文件名
	FileName string `json:"file_name,omitempty"`
	// FileType 仅当 status 为 SUCCEED 时返回。
	// 线索文件的文件类型。
	// 枚举值： csv, zip
	FileType string `json:"file_type,omitempty"`
}
