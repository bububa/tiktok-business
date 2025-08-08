package oauth

import (
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/util"
)

func URL(clt *core.SDKClient, redirectURI string, state string) string {
	values := util.NewURLValues()
	defer util.ReleaseURLValues(values)
	values.Set("app_id", clt.AppID())
	values.Set("redirect_uri", redirectURI)
	if state != "" {
		values.Set("state", state)
	}
	return util.StringsJoin(core.BASE_OAUTH_URL, "?", values.Encode())
}
