package tool

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tool"
)

// OSVersion 获取操作系统列表
func OSVersion(ctx context.Context, clt *core.SDKClient, req *tool.OSVersionRequest, accessToken string) ([]tool.OSVersion, error) {
	var ret tool.OSVersionResponse
	if err := clt.Get(ctx, "v1.3/tool/os_version/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.OSVersions, nil
}
