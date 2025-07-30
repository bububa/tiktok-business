package advertiser

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InfoRequest 获取广告账号信息 API Request
type InfoRequest struct {
	// AdvertiserIDs 需要查看详细信息的广告主 ID 列表。
	// 可从获取广告主账号列表接口获取 ID。
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// Fields 需要获取的字段。
	// 如果您没有传入该字段，系统将默认返回除company_name_editable外的所有数据。
	//
	// 可选值：
	// telephone_number, contacter, currency, cellphone_number, timezone, advertiser_id, role, company, status, description, rejection_reason, address, name, language, industry, license_no, email, license_url, country, balance, create_time, display_timezone, owner_bc_id, company_name_editable。
	//
	// 含义详见返回参数。
	//
	// 注意：若您在本字段的值中包含了company_name_editable，则仅返回company_name_editable和 advertiser_id ，同时忽略通过本字段指定的其他值。
	Fields []string `json:"fields,omitempty"`
}

// Encode implements GetRequest interface
func (r *InfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_ids", string(util.JSONMarshal(r.AdvertiserIDs)))
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InfoResponse 获取广告账号信息 API Response
type InfoResponse struct {
	model.BaseResponse
	Data struct {
		// List 广告账号（广告主账号）信息列表
		List []Advertiser `json:"list,omitempty"`
	} `json:"data"`
}
