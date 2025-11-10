package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// SessionDelete 删除最大投放量或创意作品加热时段
func SessionDelete(ctx context.Context, clt *core.SDKClient, req *gmvmax.SessionDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/gmv_max/session/delete/", req, nil, accessToken)
}
