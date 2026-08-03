package account

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	accountmodel "github.com/bububa/tiktok-business/model/account"
)

// VerificationSubmit submits a verification request for a Business Center or ad account.
func VerificationSubmit(
	ctx context.Context,
	clt *core.SDKClient,
	req *accountmodel.VerificationSubmitRequest,
	accessToken string,
) error {
	return clt.Post(
		ctx,
		"v1.3/account/verification/submit/",
		req,
		nil,
		accessToken,
	)
}
