package report

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/report"
)

func Get(ctx context.Context, clt *core.SDKClient, req *report.GetRequest, accessToken string) (*report.GetResult, error) {
	var resp report.GetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/report/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
