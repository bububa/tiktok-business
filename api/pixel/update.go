package pixel

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pixel"
)

// Update 更新pixel
func Update(ctx context.Context, clt *core.SDKClient, req *pixel.UpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/pixel/update/", req, nil, accessToken)
}
