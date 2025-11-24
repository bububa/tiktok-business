package field

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/field"
)

// Get 获取即时表单字段值
// 您可以使用本接口获取即时表单（线索表单）的字段值。
// 我们推荐您通过本接口直接获取即时表单值。您也可以使用/page/lead/mock/create/接口创建测试线索，获取即时表单字段值。
func Get(ctx context.Context, clt *core.SDKClient, req *field.GetRequest, accessToken string) (*field.Field, error) {
	var ret field.GetResponse
	if err := clt.Get(ctx, "v1.3/page/field/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
