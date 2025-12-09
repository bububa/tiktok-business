package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// AutoMessageSort Sort the automatic message for a Business Account
func AutoMessageSort(ctx context.Context, clt *core.SDKClient, req *message.AutoMessageSortRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/business/message/auto_message/sort/", req, nil, accessToken)
}
