package customconversion

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/customconversion"
)

// Create 创建自定义转化
func Create(ctx context.Context, clt *core.SDKClient, req *customconversion.CreateRequest, accessToken string) (string, error) {
	var ret customconversion.CreateResponse
	if err := clt.Post(ctx, "v1.3/custom_conversion/create/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.CustomConversionID, nil
}
