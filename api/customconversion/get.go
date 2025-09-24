package customconversion

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/customconversion"
)

// Get 获取自定义转化
func Get(ctx context.Context, clt *core.SDKClient, req *customconversion.GetRequest, accessToken string) (*customconversion.CustomConversion, error) {
	var ret customconversion.GetResponse
	if err := clt.Get(ctx, "v1.3/custom_conversion/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
