package term

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/term"
)

// Get 查看协议
// 您可以使用本接口查看协议内容。您需要使用term_type字段指定希望查看的协议类型，目前仅支持查看线索广告协议。
func Get(ctx context.Context, clt *core.SDKClient, req *term.GetRequest, accessToken string) ([]string, error) {
	var ret term.GetResponse
	if err := clt.Get(ctx, "v1.3/term/get/", req, &ret, accessToken); err != nil {
		return nil, err
	}
	return ret.Data.Terms, nil
}
