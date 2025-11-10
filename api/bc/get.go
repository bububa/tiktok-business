package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// Get 获取商务中心列表
func Get(ctx context.Context, clt *core.SDKClient, req *bc.GetRequest, accessToken string) (*bc.GetResult, error) {
	var resp bc.GetResponse
	if err := clt.Get(ctx, "v1.3/bc/get", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
