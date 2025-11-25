package blockedword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment/blockedword"
)

// Delete 删除屏蔽词
func Delete(ctx context.Context, clt *core.SDKClient, req *blockedword.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/blockedword/delete/", req, nil, accessToken)
}
