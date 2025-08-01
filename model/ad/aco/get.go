package aco

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取智能创意广告素材 API Request
type GetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupIDs 广告组ID，允许数量范围: 1-100。
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// ExcludeFieldTypesInResponse 想要从返回中移除的字段种类。
	// 允许的枚举值：
	// NULL_FIELD: 值为null的字段。
	ExcludeFieldTypesInResponse []enum.ExcludeFieldType `json:"exclude_field_types_in_response,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("adgroup_ids", string(util.JSONMarshal(r.AdgroupIDs)))
	if len(r.ExcludeFieldTypesInResponse) > 0 {
		values.Set("exclude_field_types_in_response", string(util.JSONMarshal(r.ExcludeFieldTypesInResponse)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取智能创意广告素材 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// List 智能创意广告素材列表
		List []Ad `json:"list,omitempty"`
	} `json:"data"`
}
