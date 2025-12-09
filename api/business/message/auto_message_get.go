package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// AutoMessageGet Get the automatic messages for a Business Account
func AutoMessageGet(ctx context.Context, clt *core.SDKClient, req *message.AutoMessageGetRequest, accessToken string) (*message.AutoMessageGetResult, error) {
	var resp message.AutoMessageGetResponse
	if err := clt.Get(ctx, "v1.3/business/message/auto_message/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
