package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserQualificationGetRequest 获取商务中心中已有的资质信息 API Request
type AdvertiserQualificationGetRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// Filtering 筛选条件
	Filtering *AdvertiserQualificationGetFilter `json:"filtering,omitempty"`
	// Page 当前页数，默认值: 1，取值范围: ≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，默认值: 10，取值范围: 1-100
	PageSize int `json:"page_size,omitempty"`
}

// AdvertiserQualificationGetFilter 筛选条件
type AdvertiserQualificationGetFilter struct {
	// Verified 资质是否经审核
	Verified bool `json:"verified"`
}

// Encode implements GetRequest
func (r *AdvertiserQualificationGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AdvertiserQualificationGetResponse 获取商务中心中已有的资质信息 API Response
type AdvertiserQualificationGetResponse struct {
	model.BaseResponse
	Data *AdvertiserQualificationGetResult `json:"data,omitempty"`
}

type AdvertiserQualificationGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Qualifications 资质信息
	Qualificiations []Qualification `json:"qualifications,omitempty"`
}

// Qualification 资质信息。
type Qualification struct {
	// QualificationID 资质ID
	QualificationID string `json:"qualification_id,omitempty"`
	// CompanyName 广告账户公司名称
	CompanyName string `json:"company_name,omitempty"`
	// Status 资质审核状态。枚举值：VERIFIED , UNVERIFIED
	Status enum.QualificationStatus `json:"status,omitempty"`
	// OwnerAdvertiserID 持有资质的广告账户ID
	OwnerAdvertiserID string `json:"owner_advertiser_id,omitempty"`
	// LinkedAdvertiserCount 关联的广告账户的数量
	LinkedAdvertiserCount int `json:"linked_advertiser_count,omitempty"`
	// RegionCode 国家或地区代码。例如，“DE”
	RegionCode string `json:"region_code,omitempty"`
}
