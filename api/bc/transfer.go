package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// Transfer Process a payment
// Use this endpoint to recharge money to or deduct money from an ad account in a Business Center, or to increase or decrease the credit balance of a Business Center.
// Both Agency and Direct Business Centers are able to perform payment-related operations.
func Transfer(ctx context.Context, clt *core.SDKClient, req *bc.TransferRequest, accessToken string) (*bc.TransferResult, error) {
	var resp bc.TransferResponse
	if err := clt.Post(ctx, "v1.3/bc/transfer/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
