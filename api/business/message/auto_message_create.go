package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// AutoMessageCreate Create an automatic message for a Business Account
func AutoMessageCreate(ctx context.Context, clt *core.SDKClient, req *message.AutoMessageCreateRequest, accessToken string) (string, error) {
	var resp message.AutoMessageCreateResponse
	if err := clt.Post(ctx, "v1.3/business/message/auto_message/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.AutoMessage.AutoMessageID, nil
}
