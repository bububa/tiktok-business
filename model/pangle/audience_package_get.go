package pangle

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AudiencePackageGetRequest 获取Pangle人群包 API Request
type AudiencePackageGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest
func (r *AudiencePackageGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AudiencePackageGetResponse 获取Pangle人群包 API Response
type AudiencePackageGetResponse struct {
	model.BaseResponse
	Data struct {
		// Packages Pangle 流量包列表
		Packages []AudiencePackage `json:"packages,omitempty"`
	} `json:"data"`
}

// AudiencePackage Pangle 流量包
type AudiencePackage struct {
	// BindType 流量包类型。枚举值: EXCLUDE（排除流量）,INCLUDE（包含流量）
	BindType string `json:"bind_type,omitempty"`
	// PackageID 流量包 ID
	PackageID string `json:"package_id,omitempty"`
	// PackageName 流量包名称
	PackageName string `json:"package_name,omitempty"`
}
