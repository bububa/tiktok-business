package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// Share 分享受众
func Share(ctx context.Context, clt *core.SDKClient, req *customaudience.ShareRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/share/", req, nil, accessToken)
}
