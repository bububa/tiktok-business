package lead

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/lead"
)

// Get 获取即时表单线索或私信线索
func Get(ctx context.Context, clt *core.SDKClient, req *lead.GetRequest, accessToken string) (*lead.Lead, error) {
	var resp lead.GetResponse
	if err := clt.Get(ctx, "v1.3/lead/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
