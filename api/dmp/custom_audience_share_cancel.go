package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceShareCancel 取消受众分享
func CustomAudienceShareCancel(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceShareCancelRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/share/cancel/", req, nil, accessToken)
}
