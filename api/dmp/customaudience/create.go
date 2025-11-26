package customaudience

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp/customaudience"
)

// Create 通过文件创建受众
// 您可以使用本接口通过上传的文件创建受众。您需要先上传受众数据文件，获得一个全局唯一的file_path，并使用该file_path创建一个对应的受众群体。
// 请查看流式API - 客户文件API相关提示，了解电子邮件地址标准化的相关提示
func Create(ctx context.Context, clt *core.SDKClient, req *customaudience.CreateRequest, accessToken string) (string, error) {
	var resp customaudience.CreateResponse
	if err := clt.Post(ctx, "v1.3/dmp/custom_audience/create/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.CustomAudienceID, nil
}
