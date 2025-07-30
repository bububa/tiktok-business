package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
)

// Token 获得长期有效的访问令牌
func Token(ctx context.Context, clt *core.SDKClient, code string) (*oauth.AccessToken, error) {
	req := oauth.TokenRequest{
		ClientID:     clt.AppID(),
		ClientSecret: clt.Secret(),
		Code:         code,
		GrantType:    "authorization_code",
	}
	var ret oauth.TokenResponse
	if err := clt.Post(ctx, "v1.3/oauth2/token/", &req, &ret, ""); err != nil {
		return nil, err
	}
	return &ret.AccessToken, nil
}
