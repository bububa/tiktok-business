package advertiser

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/advertiser"
)

// TransactionGet Get the transaction records of ad accounts
// Use this endpoint to get the transaction records of ad accounts in the Business Center. This endpoint only returns the transaction records of ad accounts that the Business Center has administrator rights over.
// Both Agency and Direct Business Centers are able to perform payment-related operations.
func TransactionGet(ctx context.Context, clt *core.SDKClient, req *advertiser.TransactionGetRequest, accessToken string) (*advertiser.TransactionGetResult, error) {
	var ret advertiser.TransactionGetResponse
	if err := clt.Get(ctx, "v1.3/advertiser/transaction/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
