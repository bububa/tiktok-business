package customconversion

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/customconversion"
)

// Delete 删除自定义转化
func Delete(ctx context.Context, clt *core.SDKClient, req *customconversion.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/custom_conversion/delete/", req, nil, accessToken)
}
