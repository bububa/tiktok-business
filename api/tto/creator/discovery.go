package creator

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/creator"
)

func LabelGet(ctx context.Context, clt *core.SDKClient, req *creator.LabelGetRequest, accessToken string) (*creator.LabelGetResult, error) {
	var resp creator.LabelGetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/category/label/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func RankGet(ctx context.Context, clt *core.SDKClient, req *creator.RankGetRequest, accessToken string) (*creator.RankGetResult, error) {
	var resp creator.RankGetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/rank/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func Discover(ctx context.Context, clt *core.SDKClient, req *creator.DiscoverRequest, accessToken string) (*creator.DiscoverResult, error) {
	var resp creator.DiscoverResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/creator/discover/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
