package negativekeyword

import (
	"context"
	"io"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/searchad/negativekeyword"
)

// Download 下载否定关键词
func Download(ctx context.Context, clt *core.SDKClient, req *negativekeyword.DownloadRequest, accessToken string, w io.Writer) error {
	resp := negativekeyword.NewDownloadResponse(w)
	return clt.Get(ctx, "v1.3/search_ad/negative_keyword/download/", req, resp, accessToken)
}
