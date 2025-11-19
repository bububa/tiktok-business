package portfolio

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative/portfolio"
)

// Create 创建素材包
func Create(ctx context.Context, clt *core.SDKClient, req *portfolio.CreateRequest, accessToken string) (string, error) {
	var resp portfolio.CreateResponse
	if err := clt.Post(ctx, "v1.3/creative/portfolio/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.CreativePortfolioID, nil
}
