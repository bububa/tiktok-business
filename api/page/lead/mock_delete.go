package lead

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/lead"
)

// MockDelete 删除测试线索
func MockDelete(ctx context.Context, clt *core.SDKClient, req *lead.MockDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/page/lead/mock/create/", req, nil, accessToken)
}
