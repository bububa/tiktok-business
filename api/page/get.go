package page

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page"
)

// Get 获取页面 ID
func Get(ctx context.Context, clt *core.SDKClient, req *page.GetRequest, accessToken string) (*page.GetResult, error) {
	var ret page.GetResponse
	if err := clt.Get(ctx, "v1.3/page/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
