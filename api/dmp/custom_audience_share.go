package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceShare 分享受众
func CustomAudienceShare(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceShareRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/share/", req, nil, accessToken)
}
