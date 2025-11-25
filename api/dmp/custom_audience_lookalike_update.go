package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceLookalikeUpdate 刷新相似受众
func CustomAudienceLookalikeUpdate(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceLookalikeUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/lookalike/update/", req, nil, accessToken)
}
