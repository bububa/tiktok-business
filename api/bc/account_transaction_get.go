package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AccountTransactionGet Get the transaction records of a BC or ad accounts
func AccountTransactionGet(ctx context.Context, clt *core.SDKClient, req *bc.AccountTransactionGetRequest, accessToken string) (*bc.AccountTransactionGetResult, error) {
	var resp bc.AccountTransactionGetResponse
	if err := clt.Get(ctx, "v1.3/bc/account/transaction/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
