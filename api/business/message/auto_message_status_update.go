package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// AutoMessageStatusUpdate Turn on or turn off an automatic message for a Business Account
func AutoMessageStatusUpdate(ctx context.Context, clt *core.SDKClient, req *message.AutoMessageStatusUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/business/message/auto_message/status/update/", req, nil, accessToken)
}
