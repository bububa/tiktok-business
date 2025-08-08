package adgroup

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup"
)

// Quota 获取投放中广告组的动态配额
func Quota(ctx context.Context, clt *core.SDKClient, req *adgroup.QuotaRequest, accessToken string) (*adgroup.Quota, error) {
	var ret adgroup.QuotaResponse
	if err := clt.Get(ctx, "v1.3/adgroup/quota/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
