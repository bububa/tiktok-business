package video

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PackageGetRequest 获取视频包 API Request
type PackageGetRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// ShoppingAdsVideoPackageID 商品库视频包ID
	ShoppingAdsVideoPackageID string `json:"shopping_ads_video_package_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *PackageGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if r.ShoppingAdsVideoPackageID != "" {
		values.Set("shopping_ads_video_package_id", r.ShoppingAdsVideoPackageID)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// PackageGetResponse 获取视频包 API Response
type PackageGetResponse struct {
	model.BaseResponse
	Data *Package `json:"data,omitempty"`
}

type Package struct {
	ShoppingAdsVideoPackageID string `json:"shopping_ads_video_package_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// Name 视频模板名称
	Name string `json:"name,omitempty"`
	// CreateTime 创建时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// UpdateTime 更新时间
	UpdateTime model.DateTime `json:"update_time,omitzero"`
	// Audit 视频包审核信息
	Audit *PackageAudit `json:"audit,omitempty"`
}

// PackageAudit 视频包审核信息
type PackageAudit struct {
	// AuditStatus 审核状态。枚举值: PROCESSING(处理中), APPROVED(已过审), REJECTED(已被拒审)
	AuditStatus string `json:"audit_status,omitempty"`
	// RejectInfo 拒审信息。只有审核状态为拒审时返回
	RejectInfo []PackageRejectInfo `json:"reject_info,omitempty"`
}

// PackageRejectInfo 拒审信息
type PackageRejectInfo struct {
	// AuditObject 拒审对象
	AuditObject string `json:"audit_object,omitempty"`
	// Reason 拒审原因
	Reason string `json:"reason,omitempty"`
}
