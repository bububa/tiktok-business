package image

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdInfoRequest 获取图片信息 API Request
type AdInfoRequest struct {
	// AdvertiserID 广告账号
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ImageIDs 图片 ID 列表。单次最多请求 100 个。
	// 注意：若您在TikTok 广告管理平台创建广告后，使用/ad/get/获取对应广告缩略图的图片ID，则这些缩略图的图片ID不支持传入本字段，因为这些缩略图未上传到素材库，也未和广告账户绑定。
	ImageIDs []string `json:"image_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r AdInfoRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("image_ids", string(util.JSONMarshal(r.ImageIDs)))
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AdInfoResponse 获取图片信息 API Response
type AdInfoResponse struct {
	model.BaseResponse
	Data struct {
		// List 图片数据
		List []Image `json:"list,omitempty"`
	} `json:"data"`
}
