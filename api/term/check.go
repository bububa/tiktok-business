package term

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/term"
)

// Check 检查协议状态
// 您可以使用本接口检查协议的签订情况，确认签订完成后才可进行协议相关内容的操作。
// 若协议已签订完成，您可以通过/term/get/获取协议。
func Check(ctx context.Context, clt *core.SDKClient, req *term.CheckRequest, accessToken string) (bool, error) {
	var ret term.CheckResponse
	if err := clt.Get(ctx, "v1.3/term/check/", req, &ret, accessToken); err != nil {
		return false, err
	}
	return ret.Data.Confirmed, nil
}
