package diagnostic

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ProductTaskGetRequest 异步下载商品库商品诊断信息 API Request
type ProductTaskGetRequest struct {
	// CatalogID 商品库ID。您需有访问该商品库的权限。
	CatalogID string `json:"catalog_id,omitempty"`
	// BcID 商务中心ID。该商务中心需为所指定商品库的所属商务中心，或所指定商品库已作为资产分享给该商务中心
	BcID string `json:"bc_id,omitempty"`
	// TaskID 商品库诊断信息下载任务的ID。
	// 若想获取该ID，可使用/diagnostic/catalog/product/task/create/。
	TaskID string `json:"task_id,omitempty"`
}

func (r *ProductTaskGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("catalog_id", r.CatalogID)
	values.Set("bc_id", r.BcID)
	values.Set("task_id", r.TaskID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ProductTaskGetResponse 异步下载商品库商品诊断信息 API Response
type ProductTaskGetResponse struct {
	model.BaseResponse
	Data *ProductTaskGetResult `json:"data,omitempty"`
}

type ProductTaskGetResult struct {
	// Status 商品库诊断信息下载任务的状态。
	// 枚举值：
	// SUCCEED ：任务已成功。您可以通过 diagnostic_file_url 字段返回的 URL 下载商品库诊断信息文件。
	// PROCESSING ：任务在处理中。
	// FAILED ：任务失败。
	Status string `json:"status,omitempty"`
	// DiagnosticFileURL 用于下载包含商品库诊断信息的 CSV 文件的 URL。
	// 注意：此 URL 不会过期。
	DiagnosticFileURL string `json:"diagnostic_file_url,omitempty"`
}
