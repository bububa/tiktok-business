package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PhoneRegionCodeRequest 获取电话号码的地区代码和区号 API Request
type PhoneRegionCodeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *PhoneRegionCodeRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// PhoneRegionCodeResponse 获取电话号码的地区代码和区号 API Response
type PhoneRegionCodeResponse struct {
	model.BaseResponse
	Data struct {
		// PhoneRegionCodeInfos 电话号码对应的地区名称，地区代码和区号信息
		PhoneRegionCodeInfos []PhoneRegionCodeInfo `json:"phone_region_code_infos,omitempty"`
	} `json:"data"`
}

// PhoneRegionCodeInfo 电话号码对应的地区名称，地区代码和区号信息
type PhoneRegionCodeInfo struct {
	// PhoneRegionName 电话号码对应的地区名称。
	// 例如： "Austria"。
	PhoneRegionName string `json:"phone_region_name,omitempty"`
	// PhoneRegionCode 电话号码对应的地区代码。
	// 例如： "AT"。
	PhoneRegionCode string `json:"phone_region_code,omitempty"`
	// PhoneRegionCallingCode 电话号码对应的区号。
	// 例如： "+43"。
	PhoneRegionCallingCode string `json:"phone_region_calling_code,omitempty"`
}
