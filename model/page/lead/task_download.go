package lead

import (
	"io"
	"net/http"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskDownloadRequest 下载线索 API Request
type TaskDownloadRequest struct {
	// AdvertiserID 如果相应的表单和线索在广告账号下（即未迁移到商务中心），你必须传入广告账号ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// LibraryID 如果相应的表单和线索已经迁移到商务中心下，你必须传入相应的表单库ID
	LibraryID string `json:"library_id,omitempty"`
	// TaskID 任务 ID
	TaskID string `json:"task_id,omitempty"`
	// XLeadRegion 对于欧洲经济区/瑞士/英国市场的线索数据，使用URL https://business-api.tiktok.com/open_api/并指定 Header x-lead-region: eu
	// 对于美国市场的线索数据，选择以下 URL 和 Header 设置方式之一
	// 对于其他市场的线索数据，基础 URL 为https://business-api.tiktok.com/open_api/，无需指定 Header x-lead-region
	XLeadRegion string `json:"-"`
}

// Encode implements GetRequest
func (r *TaskDownloadRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("task_id", r.TaskID)
	if r.AdvertiserID != "" {
		values.Set("advertiser_id", r.AdvertiserID)
	}
	if r.LibraryID != "" {
		values.Set("library_id", r.LibraryID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// PreRequest implements PreRequest
func (r *TaskDownloadRequest) PreRequest(httpReq *http.Request) error {
	if r.XLeadRegion == "" {
		return nil
	}
	httpReq.Header.Set("x-lead-region", r.XLeadRegion)
	return nil
}

// TaskDownloadResponse 下载线索 API Response
type TaskDownloadResponse struct {
	model.BaseResponse
	w io.Writer
}

func NewTaskDownloadResponse(w io.Writer) *TaskDownloadResponse {
	return &TaskDownloadResponse{w: w}
}

// Read implements DownloadResponse
func (resp *TaskDownloadResponse) Read(r io.Reader) error {
	_, err := io.Copy(resp.w, r)
	return err
}
