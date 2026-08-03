package app

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	appmodel "github.com/bububa/tiktok-business/model/app"
)

// List gets apps under an advertiser account.
func List(ctx context.Context, clt *core.SDKClient, req *appmodel.ListRequest, accessToken string) ([]appmodel.App, error) {
	var resp appmodel.ListResponse
	if err := clt.Get(ctx, "v1.3/app/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Apps, nil
}
