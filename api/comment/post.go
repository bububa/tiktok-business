package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment"
)

// Post 回复评论
// 您可以使用本接口回复评论。
// 绑定以下两类认证身份，并且拥有评论管理权限的广告主可以回复已有的一级评论：
// 自定义用户 （CUSTOMIZED_USER）
// TikTok 用户（TT_USER）
func Post(ctx context.Context, clt *core.SDKClient, req *comment.PostRequest, accessToken string) (*comment.Comment, error) {
	var resp comment.PostResponse
	if err := clt.Post(ctx, "v1.3/comment/post/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
