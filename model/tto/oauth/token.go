package oauth

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type TokenRequest struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
	AuthCode     string `json:"auth_code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

func (r *TokenRequest) Encode() []byte { return util.JSONMarshal(r) }

type RefreshTokenRequest struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (r *RefreshTokenRequest) Encode() []byte { return util.JSONMarshal(r) }

type RevokeTokenRequest struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
}

func (r *RevokeTokenRequest) Encode() []byte { return util.JSONMarshal(r) }

type TokenResponse struct {
	model.BaseResponse
	Data *Token `json:"data,omitempty"`
}

type Token struct {
	AccessToken           string      `json:"access_token,omitempty"`
	TokenType             string      `json:"token_type,omitempty"`
	Scope                 string      `json:"scope,omitempty"`
	ExpiresIn             model.Int64 `json:"expires_in,omitempty"`
	RefreshToken          string      `json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn model.Int64 `json:"refresh_token_expires_in,omitempty"`
	OpenID                string      `json:"open_id,omitempty"`
}

type TokenInfoRequest struct {
	AppID       string `json:"app_id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

func (r *TokenInfoRequest) Encode() []byte { return util.JSONMarshal(r) }

type TokenInfoResponse struct {
	model.BaseResponse
	Data *TokenInfo `json:"data,omitempty"`
}

type TokenInfo struct {
	AppID     string `json:"app_id,omitempty"`
	Scope     string `json:"scope,omitempty"`
	CreatorID string `json:"creator_id,omitempty"`
}
