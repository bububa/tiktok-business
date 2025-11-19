package portfolio

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative/portfolio"
)

// Delete 删除素材包
func Delete(ctx context.Context, clt *core.SDKClient, req *portfolio.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/creative/portfolio/delete/", req, nil, accessToken)
}
