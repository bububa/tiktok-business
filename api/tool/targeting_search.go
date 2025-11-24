package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// TargetingSearch 搜索地域定向标签
// 您可以使用本接口通过关键词搜索定向标签。
// 本接口可搜索两种用于定向受众的定向标签：邮政编码 ID（目前仅支持美国、加拿大、巴西、印度尼西亚、泰国、越南）和地域 ID。若要将定向标签应用于广告中，需在调用/adgroup/create/，/adgroup/update/，或/ad/audience_size/estimate/时，将本接口返回的邮政编码 ID 传给zipcode_ids字段，或将地域 ID 传给location_ids字段。
// 请注意，/tool/region/和/tool/targeting/search/接口均可用于获取地域 ID。
func TargetingSearch(ctx context.Context, clt *core.SDKClient, req *tool.TargetingSearchRequest, accessToken string) (*tool.TargetingSearchResult, error) {
	var ret tool.TargetingSearchResponse
	if err := clt.Get(ctx, "v1.3/tool/targeting/search/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
