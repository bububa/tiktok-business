package diagnostic

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/diagnostic"
)

// EventSourceIssue 获取商品库事件源诊断信息
// 您可以使用本接口获取与商品库绑定事件源的诊断信息。
func EventSourceIssue(ctx context.Context, clt *core.SDKClient, req *diagnostic.EventSourceIssueRequest, accessToken string) ([]diagnostic.EventSourceIssue, error) {
	var resp diagnostic.EventSourceIssueResponse
	if err := clt.Get(ctx, "v1.3/diagnostic/catalog/eventsource/issue", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
