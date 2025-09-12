package changelog

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TaskCheckRequest Check the status of a download task API Request
type TaskCheckRequest struct {
	// AdvertiserID Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// TaskID Task ID
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

// TaskCheckResponse Check the status of a download task
type TaskCheckResponse struct {
	model.BaseResponse
	Data struct {
		// Status Task status
		Status enum.ChangelogTaskStatus `json:"status,omitempty"`
	} `json:"data"`
}
