package property

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/property"
)

// List 获取添加至广告账户的 URL 资源列表
// 您可以使用本接口获取已添加至广告账户的 URL 资源列表。
// 欲了解添加和验证 URL 资源的详细步骤，参阅管理 URL 资源。
func List(ctx context.Context, clt *core.SDKClient) ([]property.URLProperty, error) {
	req := property.ListRequest{
		AppID:  clt.AppID(),
		Secret: clt.Secret(),
	}
	var resp property.ListResponse
	if err := clt.Get(ctx, "v1.3/business/property/list/", &req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data.URLPropertyInfoList, nil
}
