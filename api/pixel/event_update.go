package pixel

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pixel"
)

// EventUpdate 更新Pixel事件
func EventUpdate(ctx context.Context, clt *core.SDKClient, req *pixel.EventUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/pixel/event/update/", req, nil, accessToken)
}
