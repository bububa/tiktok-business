package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// PartnerAdd  添加合作伙伴
// 商务中心的管理员用户可以使用本接口添加合作伙伴或合作伙伴资产至商务中心。商务中心只能分享其拥有的资产。
func PartnerAdd(ctx context.Context, clt *core.SDKClient, req *bc.PartnerAddRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/partner/add/", req, nil, accessToken)
}
