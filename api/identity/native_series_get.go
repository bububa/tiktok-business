package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// NativeSeriesGet 获取广告账号内可用的 TikTok Series。
func NativeSeriesGet(ctx context.Context, clt *core.SDKClient, req *identity.NativeSeriesGetRequest, accessToken string) (*identity.NativeSeriesGetResult, error) {
	var resp identity.NativeSeriesGetResponse
	if err := clt.Get(ctx, "v1.3/identity/native_series/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
