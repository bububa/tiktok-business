package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// AutoMessageDelete Delete the automatic message for a Business Account
func AutoMessageDelete(ctx context.Context, clt *core.SDKClient, req *message.AutoMessageDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/business/message/auto_message/delete/", req, nil, accessToken)
}
