package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/oauth"
)

// AccessToken 获得长期有效的访问令牌
func AccessToken(ctx context.Context, clt *core.SDKClient, req *oauth.AccessTokenRequest) (*oauth.AccessToken, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var ret oauth.AccessTokenResponse
	if err := clt.Post(ctx, "v1.3/oauth2/access_token/", req, &ret, ""); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
