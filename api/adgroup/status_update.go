package adgroup

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/adgroup"
)

// StatusUpdate 更新广告组操作状态 API Request
func StatusUpdate(ctx context.Context, clt *core.SDKClient, req *adgroup.StatusUpdateRequest, accessToken string) (*adgroup.StatusUpdateResult, error) {
	var ret adgroup.StatusUpdateResponse
	if err := clt.Post(ctx, "v1.3/adgroup/status/update/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
