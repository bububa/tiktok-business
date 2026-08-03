package app

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	appmodel "github.com/bububa/tiktok-business/model/app"
)

// Track reports a single app event.
func Track(ctx context.Context, clt *core.SDKClient, req *appmodel.TrackRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/app/track/", req, nil, accessToken)
}

// Batch reports app events in bulk.
func Batch(ctx context.Context, clt *core.SDKClient, req *appmodel.BatchRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/app/batch/", req, nil, accessToken)
}
