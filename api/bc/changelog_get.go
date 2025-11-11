package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// ChangelogGet 获取商务中心的日志
func ChangelogGet(ctx context.Context, clt *core.SDKClient, req *bc.ChangelogGetRequest, accessToken string) (*bc.ChangelogGetResult, error) {
	var resp bc.ChangelogGetResponse
	if err := clt.Get(ctx, "v1.3/changelog/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
