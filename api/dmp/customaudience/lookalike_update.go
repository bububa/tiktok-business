package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// LookalikeUpdate 刷新相似受众
func LookalikeUpdate(ctx context.Context, clt *core.SDKClient, req *customaudience.LookalikeUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/lookalike/update/", req, nil, accessToken)
}
