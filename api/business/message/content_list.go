package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// ContentList Get a list of messages
func ContentList(ctx context.Context, clt *core.SDKClient, req *message.ContentListRequest, accessToken string) (*message.ContentListResult, error) {
	var resp message.ContentListResponse
	if err := clt.Get(ctx, "v1.3/business/message/content/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
