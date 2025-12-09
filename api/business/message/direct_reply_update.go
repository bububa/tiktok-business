package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// DirectReplyUpdate Enable or disable Comment-to-Message for a Business Account
// Use this endpoint to enable or disable the Comment-to-Message feature for a Business Account.
// The Comment-to-Message feature enables businesses to proactively engage with users by sending direct messages in response to high-intent comments on their video posts. A high-intent comment is one that expresses a strong intention to purchase or seek further information. When you enable this feature, the TikTok system will identify those high-intent comments, allowing you to send customized replies using /business/message/send/ and assist these users more effectively.
func DirectReplyUpdate(ctx context.Context, clt *core.SDKClient, req *message.DirectReplyUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/business/message/direct_reply/update/", req, nil, accessToken)
}
