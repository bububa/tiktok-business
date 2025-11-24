package lead

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/field"
	"github.com/bububa/tiktok-business/model/page/lead"
)

// FieldGet 获取即时表单或私信线索的字段值
func FieldGet(ctx context.Context, clt *core.SDKClient, req *lead.FieldGetRequest, accessToken string) (*field.Field, error) {
	var resp lead.FieldGetResponse
	if err := clt.Get(ctx, "v1.3/lead/field/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
