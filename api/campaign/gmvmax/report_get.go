package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// ReportGet 生成 GMV Max 推广系列报表
func ReportGet(ctx context.Context, clt *core.SDKClient, req *gmvmax.ReportGetRequest, accessToken string) (*gmvmax.ReportGetResult, error) {
	var resp gmvmax.ReportGetResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/report/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
