package pangle

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pangle"
)

// BlockListGet 获取Pangle屏蔽列表
func BlockListGet(ctx context.Context, clt *core.SDKClient, req *pangle.BlockListGetRequest, accessToken string) (*pangle.BlockListGetResult, error) {
	var resp pangle.BlockListGetResponse
	if err := clt.Get(ctx, "v1.3/pangle_block_list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
