package negativekeyword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/searchad/negativekeyword"
)

// Delete 删除否定关键词
func Delete(ctx context.Context, clt *core.SDKClient, req *negativekeyword.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/search_ad/negative_keyword/delete/", req, nil, accessToken)
}
