package blockedword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment/blockedword"
)

// Create 创建屏蔽词
func Create(ctx context.Context, clt *core.SDKClient, req *blockedword.CreateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/blockedword/create/", req, nil, accessToken)
}
