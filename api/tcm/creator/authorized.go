package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// Authorized 授权后获取账户洞察
func Authorized(ctx context.Context, clt *core.SDKClient, req *creator.AuthorizedRequest, accessToken string) (*creator.Creator, error) {
	var resp creator.AuthorizedResponse
	if err := clt.Get(ctx, "v1.3/tcm/creator/authorized/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
