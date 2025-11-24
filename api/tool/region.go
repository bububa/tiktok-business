package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// Region 基于具体设置获取可投放地域
// 您可以使用本接口获取可以推广的国家或地区代码，列表结果将基于您的推广目标，版位，定向受众设备操作系统，以及品牌安全设置。
// 若您希望获取基于广告主ID的可以投放广告的地域，请使用/search/region/接口。
func Region(ctx context.Context, clt *core.SDKClient, req *tool.RegionRequest, accessToken string) (*tool.RegionResult, error) {
	var ret tool.RegionResponse
	if err := clt.Get(ctx, "v1.3/tool/region/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
