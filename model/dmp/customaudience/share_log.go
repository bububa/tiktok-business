package customaudience

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ShareLogRequest 获取受众分享记录 API Request
type ShareLogRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceID 您想要获取分享记录的受众ID
	CustomAudienceID string `json:"custom_audience_id,omitempty"`
}

// Encode implements GetRequest
func (r *ShareLogRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("custom_audience_id", r.CustomAudienceID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ShareLogResponse 获取受众分享记录 API Response
type ShareLogResponse struct {
	model.BaseResponse
	Data struct {
		// List 受众分享记录列表
		List []ShareLog `json:"list,omitempty"`
	} `json:"data"`
}

// ShareLog 受众分享记录
type ShareLog struct {
	// SharedAdvertiserID 分享给的广告主ID
	SharedAdvertiserID string `json:"shared_advertiser_id,omitempty"`
	// SharedAdvertiserName 分享给的广告主名称
	SharedAdvertiserName string `json:"shared_advertiser_name,omitempty"`
	// CustomAudienceID 您想要获取分享记录的受众的ID
	CustomAudienceID string `json:"custom_audience_id,omitempty"`
	// Status 受众状态
	Status string `json:"status,omitempty"`
}
