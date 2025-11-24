package asset

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative/asset"
)

// Delete 删除创意资产
func Delete(ctx context.Context, clt *core.SDKClient, req *asset.DeleteRequest, accessToken string) (*asset.DeleteResult, error) {
	var resp asset.DeleteResponse
	if err := clt.Post(ctx, "v1.3/creative/asset/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
