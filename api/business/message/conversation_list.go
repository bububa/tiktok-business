package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// ConversationList Get a list of conversations
func ConversationList(ctx context.Context, clt *core.SDKClient, req *message.ConversationListRequest, accessToken string) (*message.ConversationListResult, error) {
	var resp message.ConversationListResponse
	if err := clt.Get(ctx, "v1.3/business/message/conversation/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
