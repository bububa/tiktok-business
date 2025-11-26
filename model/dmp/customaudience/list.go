package customaudience

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取受众列表 API Request
type ListRequest struct {
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
func (r *ListRequest) Encode() string {
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

// ListResponse 获取受众列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 受众群体列表
	List []Audience `json:"list,omitempty"`
}

// Audience 受众群体
type Audience struct {
	// AudienceID 受众群体ID
	AudienceID string `json:"audience_id,omitempty"`
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
	// AudienceSubType 受众子类型，表明可以使用的广告类型。枚举值：NORMAL: 常规受众，可用于非覆盖和频次广告。REACH_FREQUENCY: 覆盖和频次广告受众，只可用于覆盖和频次广告。
	AudienceSubType enum.AudienceSubType `json:"audience_sub_type,omitempty"`
	// CalculateType 加密类型。只适用于客户文件受众。枚举值详见枚举值-加密类型。
	CalculateType enum.CalculateType `json:"calculate_type,omitempty"`
	// CoverNum 覆盖人数。即TikTok上匹配的用户人数
	CoverNum int64 `json:"cover_num,omitempty"`
	// Shared 是否是共享的受众。枚举值：True，False
	Shared bool `json:"shared,omitempty"`
	// IsCreator 广告主是否是受众的拥有者。枚举值：true，false。 您可以使用shared 和 is_creator 查看广告主的所属情况和分享状态。规则如下：
	// 1. shared = true 且 is_creator = true：广告主是该受众的拥有者，且该受众由广告主分享给其他账户。
	// 2. shared = true 且 is_creator = false：广告主不是该受众的拥有者，该受众由其他账户分享给广告主。
	// 3. shared = false 且 is_creator = true：广告主是该受众的拥有者，且该受众未与其他人共享。
	IsCreator bool `json:"is_creator,omitempty"`
	// IsValid 受众群体是否可用。计算中或已过期的受众群体不可用
	IsValid bool `json:"is_valid,omitempty"`
	// IsExpiring 受众群体是否即将过期。受众群体会在过期时间的60天内进入即将过期状态
	IsExpiring bool `json:"is_expiring,omitempty"`
	// IsAutoRefresh 如果受众群体属于广告互动、应用活动或网站访客受众，此字段代表是否启用了受众群体每日自动更新功能
	IsAutoRefresh bool `json:"is_auto_refresh,omitempty"`
	// OwnerID 受众拥有者的广告主ID。
	// 若is_creator=true，owner_id为您在请求中指定的advertiser_id。
	// 若is_creator=false，owner_id为创建该受众的广告主ID。
	OwnerID string `json:"owner_id,omitempty"`
	// Rule 如果受众群体属于广告互动、应用活动或网站访客受众，此字段代表受众创建规则
	Rule string `json:"rule,omitempty"`
	// LookalikeSpec 相似受众群体的参数
	LookalikeSpec *LookalikeSpec `json:"lookalike_spec,omitempty"`
	// Msg 受众状态信息
	Msg string `json:"msg,omitempty"`
	// ErrorMsg 创建、修改受众群体时的失败信息
	ErrorMsg string `json:"error_msg,omitempty"`
	// CreateTime 受众创建时间，GMT时间
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ExpiredTime 受众群体过期时间。时间格式为“YYYY-MM-DD HH:MM:SS"
	ExpiredTime model.DateTime `json:"expired_time,omitzero"`
}

// LookalikeSpec 相似受众群体的参数
type LookalikeSpec struct {
	// SourceAudienceID 种子受众ID
	SourceAudienceID string `json:"source_audience_id,omitempty"`
	// IncludeSource 是否包含种子受众群体
	IncludeSource *bool `json:"include_source,omitempty"`
	// MobileOS 操作系统。枚举值：ALL：支持iOS和安卓， ANDROID：仅支持安卓，IOS：仅支持iOS。具体请查看枚举值-相似人群手机系统
	MobileOS enum.OperatingSystem `json:"mobile_os,omitempty"`
	// Placements 投放平台。枚举值：TikTok, TopBuzz & BuzzVideo，Pangle。详见枚举值-相似受众版位
	Placements enum.LookalikePlacement `json:"placements,omitempty"`
	// LocationIDs 投放地区。具体请查看枚举值-相似受众地点
	LocationIDs []string `json:"location_ids,omitempty"`
	// AudienceSize 相似受众的大小。枚举值：NARROW, BALANCED，BROAD。具体请查看枚举值-相似受众大小
	AudienceSize enum.LookalikeAudienceSize `json:"audience_size,omitempty"`
}
