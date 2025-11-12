package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment"
)

// Reference 获取关联评论
// 您可以使用本接口获取关联评论。
// 若评论类型为COMMENT，返回该评论及其所有回复。若评论类型为REPLY，返回该回复及原评论
func Reference(ctx context.Context, clt *core.SDKClient, req *comment.ReferenceRequest, accessToken string) (*comment.ListResult, error) {
	var resp comment.ListResponse
	if err := clt.Get(ctx, "v1.3/comment/reference/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
