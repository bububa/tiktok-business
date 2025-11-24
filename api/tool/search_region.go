package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// SearchRegion 基于广告主 ID 获取可投放地域
// 您可以使用本接口获取基于广告主ID的可以投放广告的地域。
// 若您希望基于推广目标、版位、定向受众设备操作系统以及品牌安全设置来获取可以推广的国家或地区代码，请使用/tool/region/接口。
func SearchRegion(ctx context.Context, clt *core.SDKClient, req *tool.SearchRegionRequest, accessToken string) ([]tool.Region, error) {
	var ret tool.SearchRegionResponse
	if err := clt.Get(ctx, "v1.3/tool/search/region/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.RegionList, nil
}
