package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// AssetAdvertiserUnassign Unlink a TikTok account from an ad account in Business Center
// Use this endpoint to unlink a TikTok account from an ad account in a Business Center.
// Once the linkage is removed, Business Center members with access to the ad account that is previously linked to the TikTok account can no longer use the TikTok account for ad delivery from the ad account.
func AssetAdvertiserUnassign(ctx context.Context, clt *core.SDKClient, req *bc.AssetAdvertiserUnassignRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/bc/asset/advertiser/unassign/", req, nil, accessToken)
}
