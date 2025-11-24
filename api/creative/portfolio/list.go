package portfolio

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative/portfolio"
)

// List 获取广告账户下的素材包列表
func List(ctx context.Context, clt *core.SDKClient, req *portfolio.ListRequest, accessToken string) (*portfolio.ListResult, error) {
	var resp portfolio.ListResponse
	if err := clt.Get(ctx, "v1.3/creative/portfolio/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
