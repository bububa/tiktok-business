package app

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	appmodel "github.com/bububa/tiktok-business/model/app"
)

// OptimizationEvent gets app conversion events available for optimization.
func OptimizationEvent(
	ctx context.Context,
	clt *core.SDKClient,
	req *appmodel.OptimizationEventRequest,
	accessToken string,
) ([]appmodel.ConversionEvent, error) {
	var resp appmodel.OptimizationEventResponse
	if err := clt.Get(ctx, "v1.3/app/optimization_event/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.OptimizationEvents, nil
}

// OptimizationEventRetargeting gets app retargeting events.
func OptimizationEventRetargeting(
	ctx context.Context,
	clt *core.SDKClient,
	req *appmodel.OptimizationEventRetargetingRequest,
	accessToken string,
) ([]appmodel.ConversionEvent, error) {
	var resp appmodel.OptimizationEventRetargetingResponse
	if err := clt.Get(ctx, "v1.3/app/optimization_event/retargeting/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.OptimizationEvents, nil
}
