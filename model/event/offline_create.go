package event

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OfflineCreateRequest 创建线下事件组 API Request
type OfflineCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Name 线下事件组名称
	Name string `json:"name,omitempty"`
	// Description 线下事件组描述
	Description string `json:"description,omitempty"`
	// AutoTracking 是否开启自动监测。
	// 如果本字段设置为true，今后在本广告主账户下新建的所有推广系列都会将该事件组用于广告监测和归因。
	AutoTracking bool `json:"auto_tracking,omitempty"`
}

// Encode implements PostRequest interface
func (r OfflineCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// OfflineCreateResponse 创建线下事件组 API Response
type OfflineCreateResponse struct {
	model.BaseResponse
	Data struct {
		// EventSetID 您刚创建的线下事件组的ID。
		EventSetID string `json:"event_set_id,omitempty"`
	} `json:"data"`
}
