package pangle

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/pangle"
)

// AudiencePackageGet 获取Pangle人群包
// 您可以使用本接口获取广告主可用 Pangle 人群包（受众）。
// 人群包仅作用于Pangle版位，与其他定向条件共同生效，不影响其他版位的投放。人群包能够会缩小投放范围，帮助您的广告投放至更合适的人群。
func AudiencePackageGet(ctx context.Context, clt *core.SDKClient, req *pangle.AudiencePackageGetRequest, accessToken string) ([]pangle.AudiencePackage, error) {
	var resp pangle.AudiencePackageGetResponse
	if err := clt.Get(ctx, "v1.3/pangle_audience_package/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Packages, nil
}
