package library

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/page/library"
)

// Transfer 迁移线索到商务中心
// 您可以使用本接口将线索从广告账户迁移至商务中心。迁移时将自动创建一个表单库，用来存储迁移的线索
func Transfer(ctx context.Context, clt *core.SDKClient, req *library.TransferRequest, accessToken string) (string, error) {
	var ret library.TransferResponse
	if err := clt.Post(ctx, "v1.3/page/library/transfer/", req, &ret, accessToken); err != nil {
		return "", err
	}
	return ret.Data.LibraryID, nil
}
