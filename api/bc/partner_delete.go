package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PartnerDelete 删除合作伙伴
// 商务中心的管理员用户可以使用本接口删除合作伙伴
func PartnerDelete(ctx context.Context, clt *core.SDKClient, req *bc.PartnerDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/partner/delete/", req, nil, accessToken)
}
