package adgroup

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup"
)

// BudgetUpdate 更新广告组预算 API Request
func BudgetUpdate(ctx context.Context, clt *core.SDKClient, req *adgroup.BudgetUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/adgroup/budget/update/", req, nil, accessToken)
}
