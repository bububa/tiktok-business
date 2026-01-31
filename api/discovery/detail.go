package discovery

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/discovery"
)

// Detail Get details of a popular hashtag
// Use this endpoint to get detailed information about a popular hashtag, such as audience insights and hashtag status.
// If the returned hashtag_status is ONLINE, you can use the /adgroup/create/ endpoint and pass the returned hashtag_id to the action_category_ids field. This will set the hashtag as a targeting action category for the ad group.
// You can also discover similar hashtags by using the /tool/hashtag/recommend/ endpoint. Pass the returned hashtag_name to the keywords field to get similar hashtags.
func Detail(ctx context.Context, clt *core.SDKClient, req *discovery.DetailRequest, accessToken string) (*discovery.Hashtag, error) {
	var resp discovery.DetailResponse
	if err := clt.Get(ctx, "v1.3/discovery/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
