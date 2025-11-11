package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetMemberGet 根据资产获取成员
// 商务中心的管理员用户可以使用本接口获取对某资产有访问权限的成员列表
func AssetMemberGet(ctx context.Context, clt *core.SDKClient, req *bc.AssetMemberGetRequest, accessToken string) (*bc.AssetMemberGetResult, error) {
	var resp bc.AssetMemberGetResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/member/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
