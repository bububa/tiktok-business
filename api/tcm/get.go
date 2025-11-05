package tcm

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm"
)

// Get 通过访问令牌获取 TTCM 账号
func Get(ctx context.Context, clt *core.SDKClient, req *tcm.GetRequest) (*tcm.GetResult, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp tcm.GetResponse
	if err := clt.Get(ctx, "v1.3/oauth2/tcm/get/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
