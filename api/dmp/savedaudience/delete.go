package savedaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/savedaudience"
)

// Delete 删除已保存受众
func Delete(ctx context.Context, clt *core.SDKClient, req *savedaudience.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/saved_audience/delete/", req, nil, accessToken)
}
