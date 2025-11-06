package ttvideo

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/ttvideo"
)

// AuditStatusGet 获取视频审核状态
func AuditStatusGet(ctx context.Context, clt *core.SDKClient, req *ttvideo.AuditStatusGetRequest, accessToken string) ([]ttvideo.AuditStatus, error) {
	var resp ttvideo.AuditStatusGetResponse
	if err := clt.Get(ctx, "v1.3/tcm/video/audit_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AuditStatusList, nil
}
