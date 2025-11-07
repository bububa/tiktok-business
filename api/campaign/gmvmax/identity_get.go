package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// IdentityGet 获取 GMV Max 推广系列可用的认证身份
func IdentityGet(ctx context.Context, clt *core.SDKClient, req *gmvmax.IdentityGetRequest, accessToken string) ([]gmvmax.Identity, error) {
	var resp gmvmax.IdentityGetResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/identity/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.IdentityList, nil
}
