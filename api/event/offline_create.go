package event

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/event"
)

// OfflineCreate 创建线下事件组
func OfflineCreate(ctx context.Context, clt *core.SDKClient, req *event.OfflineCreateRequest, accessToken string) (string, error) {
	var resp event.OfflineCreateResponse
	if err := clt.Post(ctx, "v1.3/offline/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.EventSetID, nil
}
