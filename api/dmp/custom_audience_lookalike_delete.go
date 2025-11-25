package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceLookalikeDelete 删除受众
func CustomAudienceLookalikeDelete(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceLookalikeDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/lookalike/delete/", req, nil, accessToken)
}
