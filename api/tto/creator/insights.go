package creator

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/creator"
)

func StatusGet(ctx context.Context, clt *core.SDKClient, req *creator.StatusGetRequest, accessToken string) ([]creator.OnboardingStatus, error) {
	var resp creator.StatusGetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/creator/status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.OnboardingStatus, nil
}
func PublicGet(ctx context.Context, clt *core.SDKClient, req *creator.PublicGetRequest, accessToken string) (*creator.Creator, error) {
	var resp creator.CreatorResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/creator/public/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func PublicVideoList(ctx context.Context, clt *core.SDKClient, req *creator.PublicVideoListRequest, accessToken string) (*creator.PublicVideoListResult, error) {
	var resp creator.PublicVideoListResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/creator/public/video/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func AuthorizedGet(ctx context.Context, clt *core.SDKClient, req *creator.AuthorizedGetRequest, accessToken string) (*creator.Creator, error) {
	var resp creator.CreatorResponse
	if err := clt.Get(ctx, "v1.3/tto/creator/authorized/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func AuthorizedVideoList(ctx context.Context, clt *core.SDKClient, req *creator.AuthorizedVideoListRequest, accessToken string) (*creator.AuthorizedVideoListResult, error) {
	var resp creator.AuthorizedVideoListResponse
	if err := clt.Get(ctx, "v1.3/tto/creator/authorized/video/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
