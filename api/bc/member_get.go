package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// MemberGet 获取成员列表
// 商务中心的管理员用户可以使用本接口获取商务中心下的成员列表及信息。
// 如果您需要获取某个成员的资产信息，可以使用/bc/asset/get/接口。您需要传入user_id或user_email。
func MemberGet(ctx context.Context, clt *core.SDKClient, req *bc.MemberGetRequest, accessToken string) (*bc.MemberGetResult, error) {
	var resp bc.MemberGetResponse
	if err := clt.Get(ctx, "v1.3/bc/member/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
