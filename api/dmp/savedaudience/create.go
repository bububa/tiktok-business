package savedaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/savedaudience"
)

// Create 创建已保存受众
// 您可以使用本接口创建已保存受众。单个广告账户下最多创建 400 个已保存受众。
func Create(ctx context.Context, clt *core.SDKClient, req *savedaudience.CreateRequest, accessToken string) (string, error) {
	var resp savedaudience.CreateResponse
	if err := clt.Post(ctx, "v1.3/dmp/saved_audience/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.SavedAudienceID, nil
}
