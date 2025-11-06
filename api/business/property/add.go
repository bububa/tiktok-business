package property

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/property"
)

// Add 将 URL 资源添加至广告账户
// 您可以使用本接口向广告账户添加要验证所有权的 URL 资源（网域或网址前缀）。
// 欲了解添加和验证 URL 资源的详细步骤，参阅管理 URL 资源。
func Add(ctx context.Context, clt *core.SDKClient, req *property.AddRequest) (*property.URLProperty, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp property.AddResponse
	if err := clt.Post(ctx, "v1.3/business/property/add/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data.URLPropertyInfo, nil
}
