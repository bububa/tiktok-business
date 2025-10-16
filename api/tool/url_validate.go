package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// URLValidate 获取 URL 的验证结果
// 您可以使用本接口验证 URL是否为 iOS 通用链接、Android 应用链接或自定义网址架构，以及该 URL 是否合法。
func URLValidate(ctx context.Context, clt *core.SDKClient, req *tool.URLValidateRequest, accessToken string) (*tool.URLValidateInfo, error) {
	var ret tool.URLValidateResponse
	if err := clt.Get(ctx, "v1.3/tool/url/validate/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.LandingPageURLInfo.ValidateInfo, nil
}
