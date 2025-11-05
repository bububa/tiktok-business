package creative

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/creative"
)

// ReportGet 创意基础/视频洞察报表
func ReportGet(ctx context.Context, clt *core.SDKClient, req *creative.ReportGetRequest, accessToken string) (*creative.ReportGetResult, error) {
	var ret creative.ReportGetResponse
	if err := clt.Get(ctx, "v1.3/creative/report/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
