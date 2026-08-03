package account

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	accountmodel "github.com/bububa/tiktok-business/model/account"
)

// VerificationStatus returns the current verification status of a Business Center or ad account.
func VerificationStatus(
	ctx context.Context,
	clt *core.SDKClient,
	req *accountmodel.VerificationStatusRequest,
	accessToken string,
) (*accountmodel.VerificationStatusResult, error) {
	var resp accountmodel.VerificationStatusResponse
	if err := clt.Get(
		ctx,
		"v1.3/account/verification/status/",
		req,
		&resp,
		accessToken,
	); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
