package property

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/property"
)

func Delete(ctx context.Context, clt *core.SDKClient, req *property.DeleteRequest) error {
	req.AppID = clt.AppID()
	req.Secret = clt.Secret()
	return clt.Post(ctx, "v1.3/business/property/delete/", req, nil, "")
}
