package pixel

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pixel"
)

// EventCreate 创建Pixel事件
func EventCreate(ctx context.Context, clt *core.SDKClient, req *pixel.EventCreateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/pixel/event/create/", req, nil, accessToken)
}
