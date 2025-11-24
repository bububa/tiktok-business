package pangle

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BlockListGetRequest 获取Pangle屏蔽列表 API Request
type BlockListGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest
func (r *BlockListGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// BlockListGetResponse 获取Pangle屏蔽列表 API Response
type BlockListGetResponse struct {
	model.BaseResponse
	Data *BlockListGetResult `json:"data,omitempty"`
}

type BlockListGetResult struct {
	// AppList 屏蔽 App 列表。iOS 应用为 iTunes ID，Android 应用为包名
	AppList []string `json:"app_list,omitempty"`
	// AppPackageID 屏蔽 App 列表 ID
	AppPackageID string `json:"app_package_id,omitempty"`
	// ModifyTime 屏蔽列表的最后更新时间，采用 ISO 8601 格式。
	// 示例："2024-01-01T00:00:01+00:00"。
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
}
