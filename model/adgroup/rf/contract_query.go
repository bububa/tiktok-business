package rf

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ContractQueryRequest 查询合同有效期 API Request
type ContractQueryRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// IncludedDate 合同有效期中的某一日期，可以是广告组排期的开始时间/结束时间。格式："2020-09-18"。
	IncludedDate string `json:"included_date,omitempty"`
	// RfCampaignType 使用本字段设置要查询的合同是否是TikTok Pulse合同。枚举值： STANDARD （查询覆盖和频次合同）， PULSE（查询TikTok Pulse合同）。
	// 注意：本字段目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	RfCampaignType enum.RfCampaignType `json:"rf_campaign_type,omitempty"`
}

// Encode implements GetRequest
func (r *ContractQueryRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("included_date", r.IncludedDate)
	if r.RfCampaignType != "" {
		values.Set("rf_campaign_type", string(r.RfCampaignType))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ContractQueryResponse 查询合同有效期 API Response
type ContractQueryResponse struct {
	model.BaseResponse
	Data *ContractQueryResult `json:"data,omitempty"`
}

type ContractQueryResult struct {
	// HasValidContract 是否有有效合同。枚举值：True，False。
	// 仅当has_valid_contract 为True时才会返回正确的合同起止时间，此时代表传入的日期included_date在覆盖和频次合同或TikTok Pulse的有效期内。
	// 若请求中rf_campaign_type 未设置为PULSE ，则has_valid_contract的值代表传入的included_date是否在覆盖和频次合同的有效期（ [start_date, end_date]）内。included_date在覆盖和频次合同的有效期内时本字段返回True。
	// 若请求中rf_campaign_type 设置为PULSE ，则has_valid_contract的值代表传入的included_date是否在TikTok Pulse合同的有效期（ [pulse_start_date, pulse_end_date]）内。included_date在TikTok Pulse合同的有效期内时本字段返回True。
	HasValidContract bool `json:"has_valid_contract,omitempty"`
	// PulseStartDate TikTok Pulse合同的起始时间。当请求中rf_campaign_type 设置为PULSE时返回。
	// 注意：若您使用/campaign/create/创建覆盖和频次推广系列时，将rf_campaign_type 设置为PULSE，则对应广告组的投放时间([schedule_start_time, schedule_end_time]) 应处于本接口返回的两个时间段[start_date, end_date]和[pulse_start_date, pulse_end_date]之间重合的时间段。schedule_start_time和schedule_end_time需在使用/adgroup/rf/create/创建覆盖和频次广告组时传入。
	PulseStartDate string `json:"pulse_start_date,omitempty"`
	// PulseEndDate TikTok Pulse合同的终止时间。当请求中rf_campaign_type 设置为PULSE时返回。
	// 注意：若您使用/campaign/create/创建覆盖和频次推广系列时，将rf_campaign_type 设置为PULSE，则对应广告组的投放时间([schedule_start_time, schedule_end_time])应处于本接口返回的两个时间段[start_date, end_date]和[pulse_start_date, pulse_end_date]之间重合的时间段。schedule_start_time和schedule_end_time需在使用/adgroup/rf/create/创建覆盖和频次广告组时传入。
	PulseEndDate string `json:"pulse_end_date,omitempty"`
	// StartDate 覆盖和频次合同开始时间
	StartDate string `json:"start_date,omitempty"`
	// EndDate 覆盖和频次合同结束时间
	EndDate string `json:"end_date,omitempty"`
}
