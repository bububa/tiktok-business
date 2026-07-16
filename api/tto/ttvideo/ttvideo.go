package ttvideo

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/ttvideo"
)

func Apply(ctx context.Context, clt *core.SDKClient, req *ttvideo.ApplyRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tcm/tt_video/apply/", req, nil, accessToken)
}
func Status(ctx context.Context, clt *core.SDKClient, req *ttvideo.StatusRequest, accessToken string) (*ttvideo.Video, error) {
	var resp ttvideo.StatusResponse
	if err := clt.Get(ctx, "v1.3/tcm/tt_video/status/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
