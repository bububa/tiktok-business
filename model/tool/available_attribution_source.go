package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AvailableAttributionSourceRequest 获取应用可用归因来源和数据源 API Request。
type AvailableAttributionSourceRequest struct {
	// AdvertiserID 广告主 ID。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AppID 应用 ID。
	AppID string `json:"app_id,omitempty"`
}

// Encode implements GetRequest.
func (r *AvailableAttributionSourceRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("app_id", r.AppID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AvailableAttributionSourceResponse 获取应用可用归因来源和数据源 API Response。
type AvailableAttributionSourceResponse struct {
	model.BaseResponse
	Data *AvailableAttributionSourceResult `json:"data,omitempty"`
}

// AvailableAttributionSourceResult 应用可用归因来源和数据源。
type AvailableAttributionSourceResult struct {
	// AppAttributionSources 可用于应用的归因来源。
	AppAttributionSources []enum.AppAttributionSource `json:"app_attribution_source,omitempty"`
	// AppDataSources 可用于应用的数据源。
	AppDataSources []enum.AppDataSource `json:"app_data_source,omitempty"`
}
