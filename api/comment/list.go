package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment"
)

// List 获取评论列表
// 您可以使用本接口获取广告账户所发视频广告下的评论。
// 您需设置搜索值及想要获取评论的时间区间。本接口可返回导流广告以及Spark Ads的评论
func List(ctx context.Context, clt *core.SDKClient, req *comment.ListRequest, accessToken string) (*comment.ListResult, error) {
	var resp comment.ListResponse
	if err := clt.Get(ctx, "v1.3/comment/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
