package app

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	appmodel "github.com/bububa/tiktok-business/model/app"
)

// Info gets information about an app.
func Info(ctx context.Context, clt *core.SDKClient, req *appmodel.InfoRequest, accessToken string) (*appmodel.App, error) {
	var resp appmodel.InfoResponse
	if err := clt.Get(ctx, "v1.3/app/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.App, nil
}
