package ad

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ad"
)

// StatusUpdate 更新广告操作状态 API Request
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *ad.StatusUpdateRequest, accessToken string) (*ad.StatusUpdateResult, error) {
	var ret ad.StatusUpdateResponse
	if err := clt.Post(ctx, "v1.3/ad/status/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
