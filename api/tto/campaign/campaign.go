package campaign

import (
	"context"
	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/tto/campaign"
)

func Create(ctx context.Context, clt *core.SDKClient, req *campaign.CreateRequest, accessToken string) (*campaign.Campaign, error) {
	var resp campaign.CampaignResponse
	if err := clt.Post(ctx, "v1.3/tto/tcm/campaign/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func Update(ctx context.Context, clt *core.SDKClient, req *campaign.UpdateRequest, accessToken string) (*campaign.Campaign, error) {
	var resp campaign.CampaignResponse
	if err := clt.Post(ctx, "v1.3/tto/tcm/campaign/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func Get(ctx context.Context, clt *core.SDKClient, req *campaign.GetRequest, accessToken string) (*campaign.GetResult, error) {
	var resp campaign.GetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/campaign/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func Link(ctx context.Context, clt *core.SDKClient, req *campaign.LinkRequest, accessToken string) (*campaign.LinkInfo, error) {
	var resp campaign.LinkResponse
	if err := clt.Post(ctx, "v1.3/tto/tcm/campaign/link/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
func LinkStatusGet(ctx context.Context, clt *core.SDKClient, req *campaign.LinkStatusGetRequest, accessToken string) (*campaign.LinkStatusGetResult, error) {
	var resp campaign.LinkStatusGetResponse
	if err := clt.Get(ctx, "v1.3/tto/tcm/campaign/link/status/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
