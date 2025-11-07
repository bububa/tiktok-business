package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// Get 获取 GMV Max 推广系列列表
func Get(ctx context.Context, clt *core.SDKClient, req *gmvmax.GetRequest, accessToken string) (*gmvmax.GetResult, error) {
	var resp gmvmax.GetResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/campaign/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
