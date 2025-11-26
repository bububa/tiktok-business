package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// Apply 将受众应用到多个广告组
func Apply(ctx context.Context, clt *core.SDKClient, req *customaudience.ApplyRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/apply/", req, nil, accessToken)
}
