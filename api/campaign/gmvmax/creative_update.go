package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

func CreativeUpdate(ctx context.Context, clt *core.SDKClient, req *gmvmax.CreativeUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/campaign/gmv_max/creative/update/", req, nil, accessToken)
}
