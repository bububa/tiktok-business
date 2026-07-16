package anchor

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/anchor"
)

func Create(ctx context.Context, clt *core.SDKClient, req *anchor.CreateRequest, accessToken string) (*anchor.Anchor, error) {
	var resp anchor.CreateResponse
	if err := clt.Post(ctx, "v1.3/tto/tcm/anchor/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func CreateByFile(ctx context.Context, clt *core.SDKClient, req *anchor.CreateByFileRequest, accessToken string) (*anchor.Anchor, error) {
	var resp anchor.CreateResponse
	if err := clt.Upload(ctx, "v1.3/tto/tcm/anchor/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func Get(ctx context.Context, clt *core.SDKClient, req *anchor.GetRequest, accessToken string) (*anchor.GetResult, error) {
	var resp anchor.GetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/anchor/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func Delete(ctx context.Context, clt *core.SDKClient, req *anchor.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/tto/tcm/anchor/delete/", req, nil, accessToken)
}
