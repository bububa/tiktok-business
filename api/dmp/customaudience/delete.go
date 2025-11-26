package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// Delete 删除受众
func Delete(ctx context.Context, clt *core.SDKClient, req *customaudience.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/delete/", req, nil, accessToken)
}
