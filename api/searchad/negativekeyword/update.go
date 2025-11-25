package negativekeyword

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/searchad/negativekeyword"
)

// Update 更新否定关键词
func Update(ctx context.Context, clt *core.SDKClient, req *negativekeyword.UpdateRequest, accessToken string) (string, error) {
	var resp negativekeyword.UpdateResponse
	if err := clt.Post(ctx, "v1.3/search_ad/negative_keyword/update/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.NewKeywordID, nil
}
