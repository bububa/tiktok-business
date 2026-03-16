package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AccountBudgetChangelogGet Get the budget change history of an ad account
// Use this endpoint to retrieve the budget change history of an ad account within a Business Center.
func AccountBudgetChangelogGet(ctx context.Context, clt *core.SDKClient, req *bc.AccountBudgetChangelogGetRequest, accessToken string) (*bc.AccountBudgetChangelogGetResult, error) {
	var resp bc.AccountBudgetChangelogGetResponse
	if err := clt.Get(ctx, "v1.3/bc/account/budget/changelog/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
