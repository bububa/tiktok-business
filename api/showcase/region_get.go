package showcase

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/showcase"
)

// RegionGet 通过认证身份获取橱窗可定向的地区
// 您可以使用本接口通过指定认证身份获取橱窗可定向的国家或地区。指定的认证身份需与橱窗绑定，且有橱窗的权限。
func RegionGet(ctx context.Context, clt *core.SDKClient, req *showcase.RegionGetRequest, accessToken string) ([]string, error) {
	var ret showcase.RegionGetResponse
	if err := clt.Get(ctx, "v1.3/showcase/region/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.RegionCodes, nil
}
