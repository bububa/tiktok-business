package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// SessionGet 获取最大投放量或创意作品加热时段的详情
func SessionGet(ctx context.Context, clt *core.SDKClient, req *gmvmax.SessionGetRequest, accessToken string) ([]gmvmax.Session, error) {
	var resp gmvmax.SessionGetResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/session/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.SessionList, nil
}
