package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TargetingListRequest 获取互联网服务提供商列表 API Request
type TargetingListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// LocationIDs 地域ID列表，用于获取对应支持的互联网服务提供商ID。
	// 最大数量：1000。
	// 详细地域ID列表参见地域ID。
	// 注意：您需传入国家或地区层级的地域ID。您可以使用/tool/region/接口，根据您的版位和推广目标来查询您可以投放的地域以及相应ID，国家或地区层级的地域ID对应的level为COUNTRY。
	LocationIDs []string `json:"location_ids,omitempty"`
	// Scene 定向标签适用的定向类型。
	// 允许的枚举值： ISP（互联网服务提供商定向）
	Scene enum.TargetingType `json:"scene,omitempty"`
}

// Encode implements GetRequest
func (r *TargetingListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("location_ids", string(util.JSONMarshal(r.LocationIDs)))
	values.Set("scene", string(r.Scene))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// TargetingListResponse 获取互联网服务提供商列表 API Response
type TargetingListResponse struct {
	model.BaseResponse
	Data *TargetingListResult `json:"data,omitempty"`
}

type TargetingListResult struct {
	// TargetingTagList 所指定的关键词所匹配到的定向标签的信息。
	// 注意：对于同一个关键词（keyword）可能返回多个匹配到的定向标签，对应不同名称（name）。
	TargetingTagList []TargetingTag `json:"targeting_tag_list,omitempty"`
	// ParentTags 所有父层级定向标签的相关信息
	ParentTags []TargetingTag `json:"targeting_tag,omitempty"`
}
