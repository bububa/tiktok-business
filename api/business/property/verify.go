package property

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/property"
)

// Verify 检查 URL 资源验证结果
// 您可以使用本接口检查 URL 资源（网域或网址前缀）的所有权验证结果。
// 欲了解添加和验证 URL 资源的详细步骤，参阅管理 URL 资源。
func Verify(ctx context.Context, clt *core.SDKClient, req *property.VerifyRequest) (*property.URLProperty, error) {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	var resp property.VerifyResponse
	if err := clt.Post(ctx, "v1.3/business/property/verify/", req, &resp, ""); err != nil {
		return nil, err
	}
	return resp.Data.URLPropertyInfo, nil
}
