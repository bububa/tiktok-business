package ad

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad"
)

// Get 获取广告列表
func Get(ctx context.Context, clt *core.SDKClient, req *ad.GetRequest, accessToken string) (*ad.GetResult, error) {
	var ret ad.GetResponse
	if err := clt.Get(ctx, "v1.3/ad/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
