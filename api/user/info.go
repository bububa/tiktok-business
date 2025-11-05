package user

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/user"
)

// Info 获取用户信息
// 您可以使用本接口获取已获得访问某一开发者应用的授权的用户信息。
func Info(ctx context.Context, clt *core.SDKClient, accessToken string) (*user.User, error) {
	var ret user.InfoResponse
	if err := clt.Post(ctx, "v1.3/user/info/", nil, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
