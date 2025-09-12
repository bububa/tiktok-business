package changelog

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskDownloadRequest Get the downloaded file API Request
type TaskDownloadRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID Task ID
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

// TaskDownloadResponse Get the downloaded file API Response
type TaskDownloadResponse struct {
	model.BaseResponse
	Data *TaskDownloadResult `json:"data,omitempty"`
}

type TaskDownloadResult struct {
	// Status Task status
	Status enum.ChangelogTaskStatus `json:"status,omitempty"`
	// Changelog CSV file data of the log
	Changelog string `json:"changelog,omitempty"`
}

type Changelog struct {
	FileData string `json:"file_data,omitempty"`
	FileName string `json:"file_name,omitempty"`
}
