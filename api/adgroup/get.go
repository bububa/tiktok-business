package adgroup

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup"
)

// Get 获取广告组列表
func Get(ctx context.Context, clt *core.SDKClient, req *adgroup.GetRequest, accessToken string) (*adgroup.GetResult, error) {
	var ret adgroup.GetResponse
	if err := clt.Get(ctx, "v1.3/adgroup/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
