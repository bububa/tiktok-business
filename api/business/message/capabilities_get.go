package message

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/business/message"
)

// CapabilitiesGet Check the capability of a Business Account for a conversation
func CapabilitiesGet(ctx context.Context, clt *core.SDKClient, req *message.CapabilitiesGetRequest, accessToken string) ([]message.CapabilityInfo, error) {
	var resp message.CapabilitiesGetResponse
	if err := clt.Get(ctx, "v1.3/business/message/capabilities/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.CapabilityInfos, nil
}
