package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AdvertiserBalanceGet Get the balance and budget of ad accounts
// Use this endpoint to obtain the balance of ad accounts in the Business Center. You can also use this endpoint to obtain the budget of the ad accounts owned by the Business Center in auto-allocation mode.
// This endpoint only returns the ad accounts that the Business Center has administrator permissions over.
// Both Agency and Direct Business Centers are able to perform payment-related operations.
func AdvertiserBalanceGet(ctx context.Context, clt *core.SDKClient, req *bc.AdvertiserBalanceGetRequest, accessToken string) (*bc.AdvertiserBalanceGetResult, error) {
	var resp bc.AdvertiserBalanceGetResponse
	if err := clt.Get(ctx, "v1.3/bc/advertiser/balance/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
