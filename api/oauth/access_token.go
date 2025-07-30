package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
)

// AccessToken 获得长期有效的访问令牌
func AccessToken(ctx context.Context, clt *core.SDKClient, code string) (*oauth.AccessToken, error) {
	req := oauth.AccessTokenRequest{
		AppID:    clt.AppID(),
		Secret:   clt.Secret(),
		AuthCode: code,
	}
	var ret oauth.AccessTokenResponse
	if err := clt.Post(ctx, "v1.3/oauth2/access_token/", &req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
