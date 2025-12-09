package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// Send Send a message to a conversation
func Send(ctx context.Context, clt *core.SDKClient, req *message.SendRequest, accessToken string) (string, error) {
	var resp message.SendResponse
	if err := clt.Post(ctx, "v1.3/business/message/send/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.Message.MessaegID, nil
}
