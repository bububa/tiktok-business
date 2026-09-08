package report

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/report"
)

// BidProtectionDetailGet Get bid protection history
func BidProtectionDetailGet(ctx context.Context, clt *core.SDKClient, req *report.BidProtectionDetailGetRequest, accessToken string) ([]report.BidProtectionRecord, error) {
	var ret report.BidProtectionDetailGetResponse
	if err := clt.Get(ctx, "v1.3/report/bid_protection/detail/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.BidProtectionRecords, nil
}
