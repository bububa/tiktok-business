package comment

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/comment"
)

// Delete 删除评论
// 您可以使用本接口删除评论。
// 绑定以下两类认证身份，并且拥有评论管理权限的广告主可以删除自己的评论：
// 自定义用户 （CUSTOMIZED_USER）
// TikTok用户（TT_USER）
func Delete(ctx context.Context, clt *core.SDKClient, req *comment.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/comment/delete/", req, nil, accessToken)
}
