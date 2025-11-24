package lead

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/lead"
)

// MockGet 获取测试线索
func MockGet(ctx context.Context, clt *core.SDKClient, req *lead.MockGetRequest, accessToken string) (*lead.Lead, error) {
	var resp lead.MockGetResponse
	if err := clt.Get(ctx, "v1.3/page/lead/mock/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
