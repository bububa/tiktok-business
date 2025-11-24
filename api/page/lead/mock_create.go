package lead

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/lead"
)

// MockCreate 创建测试线索
// 您可以使用本接口在订阅线索后，进行线索收集的端到端测试。
// 同一个线索表单只能创建一个测试线索。如需创建新的测试线索，您必须使用/page/lead/mock/delete/删除之前的测试线索。若您想获取即时表单（线索表单）字段，可以使用/page/field/get/接口直接获取字段信息。
func MockCreate(ctx context.Context, clt *core.SDKClient, req *lead.MockCreateRequest, accessToken string) (*lead.Lead, error) {
	var resp lead.MockCreateResponse
	if err := clt.Post(ctx, "v1.3/page/lead/mock/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
