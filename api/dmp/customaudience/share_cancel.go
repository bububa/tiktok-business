package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// ShareCancel 取消受众分享
func ShareCancel(ctx context.Context, clt *core.SDKClient, req *customaudience.ShareCancelRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/share/cancel/", req, nil, accessToken)
}
