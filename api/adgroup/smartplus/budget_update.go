package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/smartplus"
)

// BudgetUpdate Update the budgets of Upgraded Smart+ Ad Groups
func BudgetUpdate(ctx context.Context, clt *core.SDKClient, req *smartplus.BudgetUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/smart_plus/adgroup/budget/update/", req, nil, accessToken)
}
