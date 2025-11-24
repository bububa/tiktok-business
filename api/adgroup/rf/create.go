package rf

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/rf"
)

// Create 创建覆盖和频次广告组
func Create(ctx context.Context, clt *core.SDKClient, req *rf.CreateRequest, accessToken string) (*rf.Adgroup, error) {
	var resp rf.CreateResponse
	if err := clt.Post(ctx, "v1.3/adgroup/rf/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
