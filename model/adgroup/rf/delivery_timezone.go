package rf

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeliveryTimezoneRequest 获取覆盖和频次广告投放地区的时区信息 API Request
type DeliveryTimezoneRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// RegionCodes 要查询的地区代码。详见 附录-地区代码
	RegionCodes []string `json:"region_codes,omitempty"`
}

// Encode implements GetRequest
func (r *DeliveryTimezoneRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("region_codes", string(util.JSONMarshal(r.RegionCodes)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// DeliveryTimezoneResponse 获取覆盖和频次广告投放地区的时区信息 API Response
type DeliveryTimezoneResponse struct {
	model.BaseResponse
	Data struct {
		// TimezoneInfo 时区信息
		TimezoneInfo []TimezoneInfo `json:"timezone_info,omitempty"`
	} `json:"data"`
}

// TimezoneInfo 时区信息
type TimezoneInfo struct {
	// RegionCode 地区代码
	RegionCode string `json:"region_code,omitempty"`
	// DeliberyTimezone 对应地区的覆盖和频次广告投放时区
	DeliveryTimezone []string `json:"delivery_timezone,omitempty"`
}
