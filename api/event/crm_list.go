package event

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/event"
)

// CRMList 获取 CRM 事件组列表
func CRMList(ctx context.Context, clt *core.SDKClient, req *event.CRMListRequest, accessToken string) ([]event.CRMEventSet, error) {
	var ret event.CRMListResponse
	if err := clt.Get(ctx, "v1.3/crm/list/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.CRMEventSets, nil
}
