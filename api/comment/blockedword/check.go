package blockedword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment/blockedword"
)

// Check 检查屏蔽词是否存在
func Check(ctx context.Context, clt *core.SDKClient, req *blockedword.CheckRequest, accessToken string) ([]blockedword.CheckWord, error) {
	var resp blockedword.CheckResponse
	if err := clt.Get(ctx, "v1.3/blockedword/check/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Results, nil
}
