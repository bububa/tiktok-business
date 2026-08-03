package account

import (
	"context"
	"errors"

	"github.com/bububa/tiktok-business/core"
	accountmodel "github.com/bububa/tiktok-business/model/account"
)

// VerificationUpload uploads one or two account verification documents.
func VerificationUpload(
	ctx context.Context,
	clt *core.SDKClient,
	req *accountmodel.VerificationUploadRequest,
	accessToken string,
) ([]string, error) {
	if req == nil || req.ImageFile1 == nil {
		return nil, errors.New("missing image_file1")
	}
	var resp accountmodel.VerificationUploadResponse
	if err := clt.Upload(
		ctx,
		"v1.3/account/verification/upload/",
		req,
		&resp,
		accessToken,
	); err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.QualificationImageIDs, nil
}
