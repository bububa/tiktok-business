package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// Carrier 获取运营商列表
func Carrier(ctx context.Context, clt *core.SDKClient, req *tool.CarrierRequest, accessToken string) ([]tool.CountryCarriers, error) {
	var ret tool.CarrierResponse
	if err := clt.Get(ctx, "v1.3/tool/carrier/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.Countries, nil
}
