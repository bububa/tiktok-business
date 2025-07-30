package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
)

// TTUserToken 获取短期访问令牌
func TTUserToken(ctx context.Context, clt *core.SDKClient, code string, redirectURI string) (*oauth.AccessToken, error) {
	req := oauth.TTUserTokenRequest{
		ClientID:     clt.AppID(),
		ClientSecret: clt.Secret(),
		AuthCode:     code,
		GrantType:    "authorization_code",
		RedirectURI:  redirectURI,
	}
	var ret oauth.AccessTokenResponse
	if err := clt.Post(ctx, "v1.3/tt_user/oauth2/token/", &req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
