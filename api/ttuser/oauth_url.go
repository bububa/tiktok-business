package ttuser

import (
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/util"
)

// OAuthURL TikTok user OAuth URL
func OAuthURL(clt *core.SDKClient, redirectURI string, state string) string {
	values := util.NewURLValues()
	defer util.ReleaseURLValues(values)
	values.Set("client_key", clt.AppID())
	values.Set("response_type", "code")
	values.Set("redirect_uri", redirectURI)
	if state != "" {
		values.Set("state", state)
	}
	return util.StringsJoin(core.TTUserOAuthURL, "?", values.Encode())
}
