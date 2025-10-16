package feed

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/catalog/feed"
)

// Delete 修改更新源
func Delete(ctx context.Context, clt *core.SDKClient, req *feed.DeleteRequest, accessToken string) (string, error) {
	var ret feed.DeleteResponse
	if err := clt.Post(ctx, "v1.3/catalog/feed/delete/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.FeedID.String(), nil
}
