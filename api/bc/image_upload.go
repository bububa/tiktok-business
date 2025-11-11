package bc

import (
	"context"

	"github.com/bububa/tiktok-business/core"
	"github.com/bububa/tiktok-business/model/bc"
)

// ImageUpload 上传资质证书
// 您可以使用本接口上传资质证书。您必须是商务中心管理员用户。
// 当您在商务中心为注册地为中国大陆、中国香港地区、法国、巴西或墨西哥的广告主创建广告账户时，您需要先上传营业执照或其他资质证明。
// 本接口将返回资质证书图片 ID（image_id）。您可以将该字段的值传入/bc/advertiser/create/接口的license_image_id字段（适用于注册地为中国大陆或中国香港地区的广告账户）或qualification_image_ids字段（适用于注册地为法国、巴西或墨西哥的广告账户）
func ImageUpload(ctx context.Context, clt *core.SDKClient, req *bc.ImageUploadRequest, accessToken string) (*bc.ImageUploadResult, error) {
	var resp bc.ImageUploadResponse
	if err := clt.Upload(ctx, "v1.3/bc/image/upload/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
