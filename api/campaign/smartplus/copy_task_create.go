package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/smartplus"
)

// CopyTaskCreate Create an asynchronous copy task for an Upgraded Smart+ Campaign
func CopyTaskCreate(ctx context.Context, clt *core.SDKClient, req *smartplus.CopyTaskCreateRequest, accessToken string) (*smartplus.CopyTaskCreateResult, error) {
	var ret smartplus.CopyTaskCreateResponse
	if err := clt.Post(ctx, "v1.3/smart_plus/campaign/copy/task/create/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
