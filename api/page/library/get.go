package library

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/library"
)

// Get 获取表单库
// 您可以使用本接口获取您可以访问的表单库信息。表单库位于商务中心下，被用于存储表单和相关线索。
func Get(ctx context.Context, clt *core.SDKClient, req *library.GetRequest, accessToken string) ([]library.Library, error) {
	var ret library.GetResponse
	if err := clt.Get(ctx, "v1.3/page/library/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
