package blockedword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment/blockedword"
)

// List 获取屏蔽词列表
func List(ctx context.Context, clt *core.SDKClient, req *blockedword.ListRequest, accessToken string) (*blockedword.ListResult, error) {
	var resp blockedword.ListResponse
	if err := clt.Get(ctx, "v1.3/blockedword/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
