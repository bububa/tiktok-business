package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// AutoMessageUpdate Update an automatic message for a Business Account
func AutoMessageUpdate(ctx context.Context, clt *core.SDKClient, req *message.AutoMessageUpdateRequest, accessToken string) (string, error) {
	var resp message.AutoMessageUpdateResponse
	if err := clt.Post(ctx, "v1.3/business/message/auto_message/update/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.AutoMessage.AutoMessageID, nil
}
