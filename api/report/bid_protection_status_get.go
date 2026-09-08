package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// BidProtectionStatusGet Get bid protection statuses
func BidProtectionStatusGet(ctx context.Context, clt *core.SDKClient, req *report.BidProtectionStatusGetRequest, accessToken string) ([]report.BidProtectionStatus, error) {
	var ret report.BidProtectionStatusGetResponse
	if err := clt.Get(ctx, "v1.3/report/bid_protection/status/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.List, nil
}
