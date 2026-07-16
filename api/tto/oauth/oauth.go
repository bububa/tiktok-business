package oauth

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/oauth"
)

func Token(ctx context.Context, clt *core.SDKClient, req *oauth.TokenRequest) (*oauth.Token, error) {
	var resp oauth.TokenResponse
	if err := clt.Post(ctx, "v1.3/tt_user/oauth2/token/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func RefreshToken(ctx context.Context, clt *core.SDKClient, req *oauth.RefreshTokenRequest) (*oauth.Token, error) {
	var resp oauth.TokenResponse
	if err := clt.Post(ctx, "v1.3/tt_user/oauth2/refresh_token/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func RevokeToken(ctx context.Context, clt *core.SDKClient, req *oauth.RevokeTokenRequest) error {
	return clt.Post(ctx, "v1.3/tt_user/oauth2/revoke/", req, nil, "")
}
func TokenInfo(ctx context.Context, clt *core.SDKClient, req *oauth.TokenInfoRequest) (*oauth.TokenInfo, error) {
	var resp oauth.TokenInfoResponse
	if err := clt.Post(ctx, "v1.3/tt_user/token_info/get/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func AccountList(ctx context.Context, clt *core.SDKClient, req *oauth.AccountListRequest, accessToken string) ([]string, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp oauth.AccountListResponse
	if err := clt.Get(ctx, "v1.3/tto/oauth2/tcm/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AccountIDs, nil
}
func AccountGet(ctx context.Context, clt *core.SDKClient, req *oauth.AccountGetRequest, accessToken string) (*oauth.Account, error) {
	var resp oauth.AccountGetResponse
	if err := clt.Get(ctx, "v1.3/tto/oauth2/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
