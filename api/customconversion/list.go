package customconversion

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/customconversion"
)

// List 获取与事件源关联的自定义转化
func List(ctx context.Context, clt *core.SDKClient, req *customconversion.ListRequest, accessToken string) (*customconversion.ListResult, error) {
	var ret customconversion.ListResponse
	if err := clt.Get(ctx, "v1.3/custom_conversion/list/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
