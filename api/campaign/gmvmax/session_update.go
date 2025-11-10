package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// SessionUpdate 更新最大投放量或创意作品加热时段
func SessionUpdate(ctx context.Context, clt *core.SDKClient, req *gmvmax.SessionUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/gmv_max/session/update/", req, nil, accessToken)
}
