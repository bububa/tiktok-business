package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// LookalikeDelete 删除受众
func LookalikeDelete(ctx context.Context, clt *core.SDKClient, req *customaudience.LookalikeDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/lookalike/delete/", req, nil, accessToken)
}
