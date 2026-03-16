package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AccountCostGet Get the cost records of a BC and ad accounts
func AccountCostGet(ctx context.Context, clt *core.SDKClient, req *bc.AccountCostGetRequest, accessToken string) (*bc.AccountCostGetResult, error) {
	var resp bc.AccountCostGetResponse
	if err := clt.Get(ctx, "v1.3/bc/account/cost/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
