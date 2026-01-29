package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad/smartplus"
)

// MaterialStatusUpdate Disable or enable creatives in an Upgraded Smart+ Ad
func MaterialStatusUpdate(ctx context.Context, clt *core.SDKClient, req *smartplus.MaterialStatusUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/smart_plus/ad/status/update/", req, nil, accessToken)
}
