package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeviceModelRequest 获取设备机型列表 API Request
type DeviceModelRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest
func (r *DeviceModelRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// DeviceModelResponse 获取设备机型列表 API Response
type DeviceModelResponse struct {
	model.BaseResponse
	Data struct {
		// DeviceModels 设备机型列表
		DeviceModels []DeviceModel `json:"device_models,omitempty"`
	} `json:"data"`
}

// DeviceModel 设备机型
type DeviceModel struct {
	// DeviceModelID 设备机型ID
	DeviceModelID string `json:"device_model_id,omitempty"`
	// DeviceModelName 设备机型名称
	DeviceModelName string `json:"device_model_name,omitempty"`
	// ChildDeviceIDs 下级设备机型 ID 列表
	ChildDeviceIDs []string `json:"child_device_ids,omitempty"`
	// IsActive 设备机型是否活跃。只有活跃的设备机型才可以用于定向
	IsActive bool `json:"is_active,omitempty"`
	// Level 设备机型层级。枚举值: BRAND,SERIES,MODEL
	Level enum.DeviceModelLevel `json:"level,omitempty"`
	// OSType 操作系统类型，枚举值: ANDROID,IOS
	OSType enum.OperatingSystem `json:"os_type,omitempty"`
}
