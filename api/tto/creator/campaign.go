package creator

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/creator"
)

func CampaignJoin(ctx context.Context, clt *core.SDKClient, req *creator.CampaignJoinRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tto/creator/campaign/join/", req, nil, accessToken)
}
func CampaignVideoLink(ctx context.Context, clt *core.SDKClient, req *creator.CampaignVideoLinkRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tto/creator/campaign/video/link/", req, nil, accessToken)
}
func LinkRequestGet(ctx context.Context, clt *core.SDKClient, req *creator.LinkRequestGetRequest, accessToken string) (*creator.LinkRequestGetResponse, error) {
	var resp creator.LinkRequestGetResponse
	if err := clt.Get(ctx, "v1.3/tto/creator/link/request/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return &resp, nil
}
func LinkRequestConfirm(ctx context.Context, clt *core.SDKClient, req *creator.LinkRequestConfirmRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tto/creator/link/request/confirm/", req, nil, accessToken)
}
