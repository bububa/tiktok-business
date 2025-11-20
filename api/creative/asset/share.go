package asset

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative/asset"
)

// Share 分享创意资产
// 您可以使用本接口与其他广告主分享创意资产。
// 您必须是广告账户和您分享创意资产的广告账户的管理员或操作员用户。同时，您需要为创意资产的所有者。
func Share(ctx context.Context, clt *core.SDKClient, req *asset.ShareRequest, accessToken string) (map[string][]string, error) {
	var resp asset.ShareResponse
	if err := clt.Post(ctx, "v1.3/creative/asset/share/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.FailedInfos, nil
}
