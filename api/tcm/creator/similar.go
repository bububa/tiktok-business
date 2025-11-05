package creator

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tcm/creator"
)

// Similar 获取相似创作者
// 您可以使用本接口获取 TikTok Creator Marketplace（TTCM）上相似创作者的信息。
func Similar(ctx context.Context, clt *core.SDKClient, req *creator.SimilarRequest, accessToken string) ([]creator.Creator, error) {
	var resp creator.SimilarResponse
	if err := clt.Get(ctx, "v1.3/tcm/similar/creator/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Creators, nil
}
