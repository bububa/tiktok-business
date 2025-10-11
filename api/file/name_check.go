package file

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/file"
)

// NameCheck 查验文件名是否重名
func NameCheck(ctx context.Context, clt *core.SDKClient, req *file.NameCheckRequest, accessToken string) (*file.NameCheckResult, error) {
	var ret file.NameCheckResponse
	if err := clt.Get(ctx, "v1.3/file/name/check/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
