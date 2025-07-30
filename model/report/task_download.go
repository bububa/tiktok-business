package report

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskDownloadRequest 下载异步报表任务结果 API Request
type TaskDownloadRequest struct {
	// AdvertiserID 广告主 ID。
	// 注意：若您调用/report/task/create/ （POST方法）时传入了advertiser_ids，您需将本字段设置为advertiser_ids中指定的广告主ID其中之一。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID 异步报表任务ID。您可在/report/task/create/ 接口中的返回数据中获得该字段。
	TaskID string `json:"task_id,omitempty"`
}

// Encode implements GetRequest interface
func (r TaskDownloadRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("task_id", r.TaskID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// TaskDownloadResponse 下载异步报表任务结果 API Response
type TaskDownloadResponse struct {
	model.BaseResponse
	Data *TaskDownloadResult `json:"data,omitempty"`
}

type TaskDownloadResult struct {
	// DownloadURL 文件下载URL。有效期1小时。若URL失效，您可以使用接口/report/task/download/获取新的URL。
	DownloadURL string `json:"download_url,omitempty"`
	// FileName 文件名称。
	FileName string `json:"file_name,omitempty"`
	// OutputFormat 输出格式。 枚举值：CSV_DOWNLOAD，XLSX_DOWNLOAD
	OutputFormat enum.OutputFormat `json:"output_format,omitempty"`
}
