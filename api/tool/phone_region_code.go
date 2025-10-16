package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// PhoneRegionCode 获取电话号码的地区代码和区号
// 您可以使用本接口获取电话号码可用的地区名称，地区代码和区号。
func PhoneRegionCode(ctx context.Context, clt *core.SDKClient, req *tool.PhoneRegionCodeRequest, accessToken string) ([]tool.PhoneRegionCodeInfo, error) {
	var ret tool.PhoneRegionCodeResponse
	if err := clt.Get(ctx, "v1.3/tool/phone_region_code/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.PhoneRegionCodeInfos, nil
}
