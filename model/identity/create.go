package identity

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建认证身份 API Request
type CreateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// DisplayName 显示名称。
	// 最大长度为100个字符
	DisplayName string `json:"display_name,omitempty"`
	// ImageURI 认证身份的头像图片的ID。
	// 头像图片的长宽比必须为1:1。
	// 若想获取头像图片的ID，您可使用/file/image/ad/upload/上传该图片。
	// 若未传入本字段，将使用默认图片作为头像
	ImageURI string `json:"image_uri,omitempty"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建认证身份 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// IdentityID 认证身份 ID
		IdentityID string `json:"identity_id,omitempty"`
	} `json:"data"`
}
