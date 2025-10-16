package feed

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/feed"
)

// Switch 修改更新源定时更新状态
func Switch(ctx context.Context, clt *core.SDKClient, req *feed.SwitchRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/catalog/feed/switch/", req, nil, accessToken)
}
