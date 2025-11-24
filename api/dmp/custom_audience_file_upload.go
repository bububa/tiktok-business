package dmp

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/dmp"
)

// CustomAudienceFileUpload 上传受众文件
// 您可以使用本接口上传受众数据文件。文件上传后，您将得到一个全局唯一的file_path，可用于后续创建或更新受众。
// 请查看流式API - 客户文件API相关提示，了解电子邮件地址标准化的相关提示。
func CustomAudienceFileUpload(ctx context.Context, clt *core.SDKClient, req *dmp.CustomAudienceFileUploadRequest, accessToken string) (string, error) {
	var resp dmp.CustomAudienceFileUploadResponse
	if err := clt.Upload(ctx, "v1.3/dmp/custom_audience/file/upload/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.FilePath, nil
}
