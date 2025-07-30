package event

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/event"
)

// Track 报告应用、网站、线下或 CRM 事件
func Track(ctx context.Context, clt *core.SDKClient, req *event.TrackRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/event/track/", req, nil, accessToken)
}
