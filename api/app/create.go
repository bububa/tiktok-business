package app

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	appmodel "github.com/bububa/tiktok-business/model/app"
)

// Create creates an app and returns its App ID.
func Create(ctx context.Context, clt *core.SDKClient, req *appmodel.CreateRequest, accessToken string) (string, error) {
	var resp appmodel.CreateResponse
	if err := clt.Post(ctx, "v1.3/app/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.AppID, nil
}
