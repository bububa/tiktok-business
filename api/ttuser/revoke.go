package ttuser

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/ttuser"
)

// Revoke 撤销短期访问令牌
func Revoke(ctx context.Context, clt *core.SDKClient, accessToken string) error {
	req := ttuser.RevokeRequest{
		ClientID:     clt.AppID(),
		ClientSecret: clt.Secret(),
		AccessToken:  accessToken,
	}
	return clt.Post(ctx, "v1.3/tt_user/oauth2/revoke/", &req, nil, "")
}
