package portfolio

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative/portfolio"
)

// Get 根据 ID 获取素材包
func Get(ctx context.Context, clt *core.SDKClient, req *portfolio.GetRequest, accessToken string) (*portfolio.Portfolio, error) {
	var resp portfolio.GetResponse
	if err := clt.Get(ctx, "v1.3/creative/portfolio/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
