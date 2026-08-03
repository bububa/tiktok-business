package app

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	appmodel "github.com/bububa/tiktok-business/model/app"
)

// Update updates an app.
func Update(ctx context.Context, clt *core.SDKClient, req *appmodel.UpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/app/update/", req, nil, accessToken)
}
