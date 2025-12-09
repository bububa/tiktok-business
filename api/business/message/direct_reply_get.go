package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// DirectReplyGet Get the Comment-to-Message setting of a Business Account
func DirectReplyGet(ctx context.Context, clt *core.SDKClient, req *message.DirectReplyGetRequest, accessToken string) (*message.DirectReplyGetResult, error) {
	var resp message.DirectReplyGetResponse
	if err := clt.Get(ctx, "v1.3/business/message/direct_reply/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
