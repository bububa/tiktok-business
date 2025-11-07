package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// Info 获取 GMV Max 推广系列的详情
func Info(ctx context.Context, clt *core.SDKClient, req *gmvmax.InfoRequest, accessToken string) (*gmvmax.Campaign, error) {
	var resp gmvmax.InfoResponse
	if err := clt.Get(ctx, "v1.3/campaign/gmv_max/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
