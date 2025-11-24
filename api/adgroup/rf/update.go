package rf

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup/rf"
)

// Update 更新覆盖和频次广告组
func Update(ctx context.Context, clt *core.SDKClient, req *rf.UpdateRequest, accessToken string) (*rf.Adgroup, error) {
	var resp rf.UpdateResponse
	if err := clt.Post(ctx, "v1.3/adgroup/rf/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
