package pangle

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pangle"
)

// BlockListUpdate 修改Pangle屏蔽列表
func BlockListUpdate(ctx context.Context, clt *core.SDKClient, req *pangle.BlockListUpdateRequest, accessToken string) (int, error) {
	var resp pangle.BlockListUpdateResponse
	if err := clt.Post(ctx, "v1.3/pangle_block_list/update/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.SuccessCount, nil
}
