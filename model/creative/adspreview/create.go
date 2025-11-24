package adspreview

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// ICreateRequest create ad preview request interface
type ICreateRequest interface {
	model.PostRequest
	GetPreviewType() enum.AdPreviewType
}

type CreateRequest struct {
	// AdvertiserID 广告主 ID
	// 注意：您需对广告账户有操作员或管理员权限。若想查看您对某个广告账户的权限，可使用/bc/asset/get/，返回的 advertiser_role 需为 ADMIN 或 OPERATOR
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PreviewType 预览类型。
	PreviewType enum.AdPreviewType `json:"preview_type,omitempty"`
}

func (r CreateRequest) GetPreviewType() enum.AdPreviewType {
	return r.PreviewType
}

// CreateResponse 广告预览 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *CreateResult `json:"data,omitempty"`
}

type CreateResult struct {
	// PreviewLink 预览链接。
	// 有效期：
	// preview_type 设置为 ADS_CREATION 或 AD时：30 天。
	// preview_type 设置为 CARD，PAGE，SINGLE_VIDEO 或 SINGLE_IMAGE时：24 小时。
	PreviewLink string `json:"preview_link,omitempty"`
	// Iframe 包含预览链接的iframe代码。您可以将iframe代码嵌入HTML。
	Iframe string `json:"iframe,omitempty"`
	// Tips 修改建议。根据提示修改使用的素材，获取更佳显示效果。
	Tips []PreviewTip `json:"tips,omitempty"`
}

type PreviewTip struct {
	// Placement 错误版位。枚举值见枚举值 - 版位。如果该字段缺失，那么表示通用提示信息（即适用于所有版位）。
	Placement enum.Placement `json:"placement,omitempty"`
	// Messages 提示信息
	Messages []string `json:"messages,omitempty"`
}
