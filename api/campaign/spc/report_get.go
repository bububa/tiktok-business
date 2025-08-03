package spc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/spc"
)

// ReportGet 生成 Smart+ 推广系列报表 API Response
func ReportGet(ctx context.Context, clt *core.SDKClient, req *spc.ReportGetRequest, accessToken string) (*spc.ReportGetResult, error) {
	var ret spc.ReportGetResponse
	if err := clt.Get(ctx, "v1.3/campaign/spc/report/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
