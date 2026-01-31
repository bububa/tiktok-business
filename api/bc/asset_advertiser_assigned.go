package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetAdvertiserAssigned Get ad accounts linked to a TikTok account in Business Center
// Use this endpoint to retrieve the list of ad accounts linked to a TikTok account in a Business Center.
// Once an ad account is linked to a TikTok account in a Business Center, Business Center members with access to the linked ad account can use the TikTok account for ad delivery from the ad account.
func AssetAdvertiserAssigned(ctx context.Context, clt *core.SDKClient, req *bc.AssetAdvertiserAssignedRequest, accessToken string) (*bc.AssetAdvertiserAssignedResult, error) {
	var resp bc.AssetAdvertiserAssignedResponse
	if err := clt.Get(ctx, "v1.3/bc/asset/advertiser/assigned/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
