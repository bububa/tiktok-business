package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// TargetingList 获取互联网服务提供商列表
// 您可以使用本接口获取定向标签。
// 仅支持获取一种用于定向受众的定向标签：互联网服务提供商ID。若要将定向标签应用于广告中，需在调用/adgroup/create/，/adgroup/update/，或/ad/audience_size/estimate/时，将本接口返回的互联网服务提供商ID传给isp_ids字段。
func TargetingList(ctx context.Context, clt *core.SDKClient, req *tool.TargetingListRequest, accessToken string) (*tool.TargetingListResult, error) {
	var ret tool.TargetingListResponse
	if err := clt.Get(ctx, "v1.3/tool/targeting/list/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
