package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// OACreate Create an Organization Account in a Business Center
// Use this endpoint to create an Organization Account in a Business Center.
// Note
//   - Creating Organization Accounts in a Business Center through API is currently an allowlist-only feature. If you would like to access it, please contact your TikTok representative.
func OACreate(ctx context.Context, clt *core.SDKClient, req *bc.OACreateRequest, accessToken string) (*bc.Asset, error) {
	var resp bc.OACreateResponse
	if err := clt.Post(ctx, "v1.3/bc/oa/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
