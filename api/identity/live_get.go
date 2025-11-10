package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// LiveGet 获取认证身份下直播视频
func LiveGet(ctx context.Context, clt *core.SDKClient, req *identity.LiveGetRequest, accessToken string) (*identity.LiveGetResult, error) {
	var resp identity.LiveGetResponse
	if err := clt.Get(ctx, "v1.3/identity/live/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
