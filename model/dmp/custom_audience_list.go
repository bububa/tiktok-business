package dmp

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CustomAudienceListRequest 获取受众列表 API Request
type CustomAudienceListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 受众ID列表。长度范围：[1, 100]
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
	// Page 当前页数。默认值: 1。取值范围: ≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。默认值: 10。取值范围：1-100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *CustomAudienceListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.CustomAudienceIDs) > 0 {
		values.Set("custom_audience_ids", string(util.JSONMarshal(r.CustomAudienceIDs)))
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

// CustomAudienceListResponse 获取受众列表 API Response
type CustomAudienceListResponse struct {
	model.BaseResponse
	Data *CustomAudienceListResult `json:"data,omitempty"`
}

type CustomAudienceListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 受众群体列表
	List []CustomAudience `json:"list,omitempty"`
}

// CustomAudience 受众群体
type CustomAudience struct {
	// Shared 是否是共享的受众。枚举值：True，False
	Shared bool `json:"shared,omitempty"`
	// IsCreator 广告主是否是受众的拥有者。枚举值：true，false。 您可以使用shared 和 is_creator 查看广告主的所属情况和分享状态。规则如下：
	// 1. shared = true 且 is_creator = true：广告主是该受众的拥有者，且该受众由广告主分享给其他账户。
	// 2. shared = true 且 is_creator = false：广告主不是该受众的拥有者，该受众由其他账户分享给广告主。
	// 3. shared = false 且 is_creator = true：广告主是该受众的拥有者，且该受众未与其他人共享。
	IsCreator bool `json:"is_creator,omitempty"`
	// AudienceID 受众群体ID
	AudienceID string `json:"audience_id,omitempty"`
	// CoverNum 覆盖人数。即TikTok上匹配的用户人数
	CoverNum int64 `json:"cover_num,omitempty"`
	// CreateTime 受众创建时间，GMT时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// IsValid 受众群体是否可用。计算中或已过期的受众群体不可用
	IsValid bool `json:"is_valid,omitempty"`
	// IsExpiring 受众群体是否即将过期。受众群体会在过期时间的60天内进入即将过期状态
	IsExpiring bool `json:"is_expiring,omitempty"`
	// ExpiredTime 受众群体过期时间。时间格式为“YYYY-MM-DD HH:MM:SS"
	ExpiredTime model.DateTime `json:"expired_time,omitzero"`
	// Name 受众群体名称
	Name string `json:"name,omitempty"`
	// AudienceType 受众群体类型。 枚举值：
	// Customer File Audience (客户文件受众)。
	// Engagement Audience (广告互动受众)。
	// App Activity Audience (应用活动受众)。
	// Website Audience (网站访客受众)。
	// Lookalike Audience (相似受众)。
	// Lead Generation (线索收集受众)
	// Rule Business Account (企业号受众)
	// Rule ad shop （商店活动受众）
	// Partner （合作伙伴受众）
	// Challenge (非标互动受众)
	// Preferred Population (优选人群受众)
	// 注意：当前，无法通过API创建合作伙伴受众，非标互动受众和优选人群受众。
	AudienceType string `json:"audience_type,omitempty"`
	// CalculateType 加密类型。只适用于客户文件受众。枚举值详见枚举值-加密类型。
	CalculateType enum.CalculateType `json:"calculate_type,omitempty"`
}
