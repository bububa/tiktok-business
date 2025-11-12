package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceUpdate 修改受众
// 您可以使用本接口修改已创建的受众群体。
// 对于客户文件受众，TikTok API for Business支持以下三种修改操作：
// REPLACE （替换）：用参数中的数据文件去替换之前（所有）的数据文件
// APPEND（新增） ：增加新的数据文件
// REMOVE （移除）：移除指定的数据文件
func CustomAudienceUpdate(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceUpdateRequest, accessToken string) error {
	return clt.Post(ctx, "v1.3/dmp/custom_audience/update/", req, nil, accessToken)
}
