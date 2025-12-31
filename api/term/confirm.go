package term

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/term"
)

// Confirm 签订协议
// 您可以使用本接口签订协议，目前仅支持签订线索广告的协议。
// 签订完成后，您可以使用/term/check/查看协议状态。
func Confirm(ctx context.Context, clt *core.SDKClient, req *term.ConfirmRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/term/confirm/", req, nil, accessToken)
}
