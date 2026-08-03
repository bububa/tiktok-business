package account

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	accountmodel "github.com/bububa/tiktok-business/model/account"
)

// VerificationFileType returns supported verification document types for a country or region.
func VerificationFileType(
	ctx context.Context,
	clt *core.SDKClient,
	req *accountmodel.VerificationFileTypeRequest,
	accessToken string,
) ([]accountmodel.VerificationFileType, error) {
	var resp accountmodel.VerificationFileTypeResponse
	if err := clt.Get(
		ctx,
		"v1.3/account/verification/filetype/",
		req,
		&resp,
		accessToken,
	); err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.FileTypes, nil
}
