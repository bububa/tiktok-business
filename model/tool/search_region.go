package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// SearchRegionRequest 基于广告主 ID 获取可投放地域 API Request
type SearchRegionRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Language 您希望返回的地域名称所显示的语言。 默认值: en
	// 目前，我们仅支持 zh, fr, es, ko, vi, en, hi, it, tr, ru, ja, id, de, ms, ar, th
	// 注意：如果您没有传入以上值之一，系统将默认处理为默认的en。
	Language string `json:"language,omitempty"`
}

// Encode implements GetRequest interface
func (r *SearchRegionRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// SearchRegionResponse 基于广告主 ID 获取可投放地域 API Response
type SearchRegionResponse struct {
	model.BaseResponse
	Data struct {
		// RegionList 可投放地域列表
		RegionList []Region `json:"region_list,omitempty"`
	} `json:"data"`
}
