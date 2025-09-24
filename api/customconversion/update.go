package customconversion

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/customconversion"
)

// Update 更新自定义转化
func Update(ctx context.Context, clt *core.SDKClient, req *customconversion.UpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/custom_conversion/update/", req, nil, accessToken)
}
