package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CarrierRequest 获取运营商列表 API Request
type CarrierRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest
func (r *CarrierRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CarrierResponse 获取运营商列表 API Response
type CarrierResponse struct {
	model.BaseResponse
	Data struct {
		// Countries 国家或地区列表
		Countries []CountryCarriers `json:"countries,omitempty"`
	} `json:"data"`
}

type CountryCarriers struct {
	// CountryCode 国家或地区代码
	CountryCode string `json:"country_code,omitempty"`
	// Carriers 当前国家或地区内运营商列表
	Carriers []Carrier `json:"carriers,omitempty"`
}

// Carrier 运营商
type Carrier struct {
	// CarrierID 运营商ID。您可在创建广告时使用
	CarrierID string `json:"carrier_id,omitempty"`
	// InUse 运营商是否可用于定向
	InUse bool `json:"in_use,omitempty"`
	// Name 运营商名称
	Name string `json:"name,omitempty"`
	// Value 运营商MCC（移动国家代码）/MNC（移动网络代码）列表
	Value []CarrierValue `json:"value,omitempty"`
}

// CarrierValue 运营商MCC（移动国家代码）/MNC（移动网络代码）
type CarrierValue struct {
	// HniID 运营商HNI ID（家庭网络身份ID）。由移动国家代码和移动网络代码组成
	HniID string `json:"hni_id,omitempty"`
	// InUse HNI ID是否可用于定向
	InUse bool `json:"in_use,omitempty"`
}
