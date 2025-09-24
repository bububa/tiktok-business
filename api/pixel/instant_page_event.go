package pixel

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pixel"
)

// InstantPageEvent 获取即时体验页面转化事件
func InstantPageEvent(ctx context.Context, clt *core.SDKClient, req *pixel.InstantPageEventRequest, accessToken string) ([]pixel.InstantPageEvent, error) {
	var ret pixel.InstantPageEventResponse
	if err := clt.Get(ctx, "v1.3/pixel/instant_page/event/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
