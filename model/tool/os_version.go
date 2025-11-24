package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// OSVersionRequest 获取操作系统列表 API Request
type OSVersionRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// OSType 枚举值: ANDROID,IOS
	OSType enum.OperatingSystem `json:"os_type,omitempty"`
}

// Encode implements GetRequest
func (r *OSVersionRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("os_type", string(r.OSType))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// OSVersionResponse 获取操作系统列表 API Response
type OSVersionResponse struct {
	model.BaseResponse
	Data struct {
		// OSVersions 受众操作系统版本列表
		OSVersions []OSVersion `json:"os_versions,omitempty"`
	} `json:"data"`
}

// OSVersion 受众操作系统版本
type OSVersion struct {
	// OSID 受众操作系统版本 ID
	OSID string `json:"os_id,omitempty"`
	// OSType 操作系统类型
	OSType enum.OperatingSystem `json:"os_type,omitempty"`
	// Version 受众操作系统版本
	Version string `json:"version,omitempty"`
	// Name 受众操作系统版本名称
	Name string `json:"name,omitempty"`
}
