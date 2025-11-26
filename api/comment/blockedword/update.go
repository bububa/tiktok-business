package blockedword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment/blockedword"
)

// Update 更新屏蔽词
func Update(ctx context.Context, clt *core.SDKClient, req *blockedword.UpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/blockedword/update/", req, nil, accessToken)
}
