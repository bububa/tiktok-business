package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// BalanceGet Get the balance of a BC
// Use this endpoint to obtain the balance of a Business Center.
// Both Agency and Direct Business Centers are able to perform payment-related operations.
func BalanceGet(ctx context.Context, clt *core.SDKClient, req *bc.BalanceGetRequest, accessToken string) (*bc.BalanceGetResult, error) {
	var resp bc.BalanceGetResponse
	if err := clt.Get(ctx, "v1.3/bc/balance/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
