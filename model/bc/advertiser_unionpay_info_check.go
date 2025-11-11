package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserUnionpayInfoCheckRequest 查询营业执照的银联验证要求 API Request
type AdvertiserUnionpayInfoCheckRequest struct {
	// LicenseNo 营业执照编号
	LicenseNo string `json:"license_no,omitempty"`
}

// Encode implements GetRequest
func (r *AdvertiserUnionpayInfoCheckRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("license_no", r.LicenseNo)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AdvertiserUnionpayInfoCheckResponse 查询营业执照的银联验证要求 API Response
type AdvertiserUnionpayInfoCheckResponse struct {
	model.BaseResponse
	Data struct {
		// UnionpayVerificationRequired 该营业执照是否需要进行银联验证。
		// 支持的值：
		// true：需要。
		// false：不需要。
		// 注意：若本字段为 true，您需使用 /bc/advertiser/unionpay_info/submit/ 为您的营业执照提交银联验证。
		UnionpayVerificationRequired bool `json:"unionpay_verification_required,omitempty"`
	} `json:"data"`
}
