package set

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/set"
)

// Update 更新商品系列
// 您可以使用本接口更新已有商品系列的筛选条件和名称。您只需传入您希望修改的字段。
func Update(ctx context.Context, clt *core.SDKClient, req *set.UpdateRequest, accessToken string) (*set.UpdateResult, error) {
	var resp set.UpdateResponse
	if err := clt.Post(ctx, "v1.3/catalog/set/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
