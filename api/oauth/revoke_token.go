package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
)

// RevokeToken 撤销长期访问令牌
func RevokeToken(ctx context.Context, clt *core.SDKClient, accessToken string) (*oauth.RevokeTokenResult, error) {
	req := oauth.RevokeTokenRequest{
		AppID:       clt.AppID(),
		Secret:      clt.Secret(),
		AccessToken: accessToken,
	}
	var ret oauth.RevokeTokenResponse
	if err := clt.Post(ctx, "v1.3/oauth2/revoke_token/", &req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
