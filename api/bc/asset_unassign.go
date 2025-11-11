package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetUnassign 撤销用户对资产的权限
func AssetUnassign(ctx context.Context, clt *core.SDKClient, req *bc.AssetUnassignRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/asset/unassign/", req, nil, accessToken)
}
