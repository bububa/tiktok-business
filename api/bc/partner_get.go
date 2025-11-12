package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PartnerGet 获取合作伙伴
// 商务中心账户管理员可以使用本接口获取合作伙伴列表
func PartnerGet(ctx context.Context, clt *core.SDKClient, req *bc.PartnerGetRequest, accessToken string) (*bc.PartnerGetResult, error) {
	var resp bc.PartnerGetResponse
	if err := clt.Get(ctx, "v1.3/bc/partner/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
