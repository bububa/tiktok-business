package identity

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/identity"
)

// MusicAuthorization 获取音乐版权信息
func MusicAuthorization(ctx context.Context, clt *core.SDKClient, req *identity.MusicAuthorizationRequest, accessToken string) ([]identity.MusicAuthorization, error) {
	var resp identity.MusicAuthorizationResponse
	if err := clt.Get(ctx, "v1.3/identity/music/authorization/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.MusicAuthorization, nil
}
