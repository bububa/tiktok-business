package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ContextualTagInfoRequest 获取内容定向标签详情 API Request
type ContextualTagInfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ContextualTagIDs 内容相关定向标签ID列表
	ContextualTagIDs []string `json:"contextual_tag_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r *ContextualTagInfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("contextual_tag_ids", string(util.JSONMarshal(r.ContextualTagIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ContextualTagInfoResponse 获取内容定向标签详情 API Response
type ContextualTagInfoResponse struct {
	model.BaseResponse
	Data struct {
		// ContextualTagList 内容相关定向标签列表。
		// 注意：内容相关定向目前为白名单功能。若您未申请进入白名单，返回值将为空列表（[]）。
		ContextualTagList []ContextualTag `json:"contextual_tag_list,omitempty"`
	} `json:"data"`
}
