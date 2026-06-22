package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/smartplus"
)

// CopyTaskCheck Check an asynchronous copy task for an Upgraded Smart+ Campaign
func CopyTaskCheck(ctx context.Context, clt *core.SDKClient, req *smartplus.CopyTaskCheckRequest, accessToken string) (*smartplus.CopyTask, error) {
	var ret smartplus.CopyTaskCheckResponse
	if err := clt.Get(ctx, "v1.3/smart_plus/campaign/copy/task/check/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
