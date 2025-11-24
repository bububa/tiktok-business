package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ContentExclusionInfoRequest 获取可用的内容排除类别 API Request
type ContentExclusionInfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CategoryIDs 内容排除类别或行业敏感内容控制类别ID。该字段可以通过/tool/content_exclusion/get/接口的返回中获得。
	CategoryIDs []string `json:"category_ids,omitempty"`
}

// Encode implements InfoRequest interface
func (r *ContentExclusionInfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("category_ids", string(util.JSONMarshal(r.CategoryIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ContentExclusionInfoResponse 获取可用的内容排除类别 API Response
type ContentExclusionInfoResponse struct {
	model.BaseResponse
	Data struct {
		// ContentExclusionList 内容排除类别或行业敏感内容控制类别列表
		ContentExclusionList []Category `json:"content_exclusion_list,omitempty"`
	} `json:"data"`
}
