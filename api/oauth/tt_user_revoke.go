package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
)

// TTUserRevoke 撤销短期访问令牌
func TTUserRevoke(ctx context.Context, clt *core.SDKClient, accessToken string) error {
	req := oauth.TTUserRevokeRequest{
		ClientID:     clt.AppID(),
		ClientSecret: clt.Secret(),
		AccessToken:  accessToken,
	}
	return clt.Post(ctx, "v1.3/tt_user/oauth2/revoke/", &req, nil, "")
}
