package blockedword

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取屏蔽词列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.PageInfo != nil {
		values.Set("page_info", string(util.JSONMarshal(r.PageInfo)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse 获取屏蔽词列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 屏蔽词列表
	List []struct {
		// BwContent 屏蔽词内容
		BwContent string `json:"bw_content,omitempty"`
	} `json:"list,omitempty"`
}
