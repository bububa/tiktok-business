package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// VideoGet 获取认证身份下帖子
func VideoGet(ctx context.Context, clt *core.SDKClient, req *identity.VideoGetRequest, accessToken string) (*identity.VideoGetResult, error) {
	var resp identity.VideoGetResponse
	if err := clt.Get(ctx, "v1.3/identity/video/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
