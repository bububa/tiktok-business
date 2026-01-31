package smartplus

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report/smartplus"
)

// MaterialReportOverview Run an Upgraded Smart+ Creative Overview Report
// Use this endpoint to retrieve an Upgraded Smart+ Creative Overview Report, an overview of the reporting data on creatives within Upgraded Smart+ Ads.
// Upgraded Smart+ Creative Overview Reports doesn't support breakdown by time range (hour, day, week, or month).
func MaterialReportOverview(ctx context.Context, clt *core.SDKClient, req *smartplus.MaterialReportRequest, accessToken string) (*smartplus.MaterialReportResult, error) {
	var ret smartplus.MaterialReportResponse
	if err := clt.Get(ctx, "v1.3/smart_plus/material_report/overview/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
