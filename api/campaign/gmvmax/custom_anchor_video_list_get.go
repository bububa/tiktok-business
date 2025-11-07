package gmvmax

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/campaign/gmvmax"
)

// CustomAnchorVideoListGet 获取自定义作品中视频的详情
// 您可以使用本接口获取商品 GMV Max 推广系列中的自定义作品合集中使用的视频详情。
func CustomAnchorVideoListGet(ctx context.Context, clt *core.SDKClient, req *gmvmax.CustomAnchorVideoListGetRequest, accessToken string) ([]gmvmax.AnchorVideo, error) {
	var resp gmvmax.CustomAnchorVideoListGetResponse
	if err := clt.Get(ctx, "v1.3/gmv_max/custom_anchor_video_list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CustomAnchorVideoList, nil
}
