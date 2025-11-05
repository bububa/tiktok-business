package ttuser

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
	"github.com/bububa/tiktok-business/model/ttuser"
)

// RefreshToken 刷新创作者访问令牌
func RefreshToken(ctx context.Context, clt *core.SDKClient, refreshToken string) (*oauth.AccessToken, error) {
	req := ttuser.RefreshTokenRequest{
		ClientID:     clt.AppID(),
		ClientSecret: clt.Secret(),
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}
	var ret oauth.AccessTokenResponse
	if err := clt.Post(ctx, "v1.3/tt_user/oauth2/refresh_token/", &req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
