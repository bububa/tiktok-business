package smartplus

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest Get Upgraded Smart+ Ads API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Fields 返回字段。查看返回中list下的字段，获取您可以指定的字段。如果您不传入该字段，系统将默认返回所有字段。
	Fields []string `json:"fields,omitempty"`
	// Filtering 筛选条件
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围：≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 10。
	// 取值范围：1-1,000。
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 筛选条件
type GetFilter struct {
	// CampaignIDs 推广系列 ID，允许数量范围: 1-100
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// AdgroupIDs 广告组id列表，支持筛选指定的广告组，允许数量范围：1-100
	AdgroupIDs []string `json:"adgroup_ids,omitempty"`
	// SmartPlusAdIDs A list of Ad IDs.
	// Max size: 100.
	SmartPlusAdIDs []string `json:"smart_plus_ad_ids,omitempty"`
	// PrimaryStatus 一级状态。枚举值详见 枚举值 - 一级状态
	// 默认值：STATUS_NOT_DELETE，返回除STATUS_DELETE外所有状态的广告组。如果您想获得包括STATUS_DELETE在内的所有状态的广告组，请使用STATUS_ALL。
	PrimaryStatus enum.PrimaryStatus `json:"primary_status,omitempty"`
	// SecondaryStatus 二级状态。枚举值详见枚举值 - 二级状态 - 推广系列状态。
	// 注意：
	// 沙箱环境下本字段不返回，因为推广系列未实际投放。
	// 请查看一级状态下所支持的二级状态 ，了解对于所指定的一级状态，您可以传入的secondary_status 有效值。
	SecondaryStatus enum.SecondaryStatus `json:"secondary_status,omitempty"`
	// ObjectiveType 推广目标，见 枚举值-推广目标
	// 注意：若您选择WEB_CONVERSIONS或CONVERSIONS的推广目标作为筛选条件，目前会同时返回以WEB_CONVERSIONS或CONVERSIONS为推广目标的数据。
	ObjectiveType enum.ObjectiveType `json:"objective_type,omitempty"`
	// SalesDestination Sales destination, the destination where you want to drive your sales.
	// Enum values:
	// WEBSITE: Website. Drive sales on your website.
	// APP: App. Drive sales on your app (product catalog required).
	// WEB_AND_APP: Website and app. Drive sales on both your website and your app.
	SalesDestination enum.SalesDestination `json:"sales_destination,omitempty"`
	// OptimizationGoal 优化目标。枚举值见枚举值-优化目标。
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// CreationFilterStartTime 筛选创建时间晚于某一时间的广告，格式为 YYYY-MM-DD HH:MM:SS（UTC 时区）。
	// 建议：为确保任务的成功执行及速度，创建时间的筛选范围建议不超过 6 个月。
	CreationFilterStartTime model.DateTime `json:"creation_filter_start_time,omitzero"`
	// CreationFilterEndTime 筛选创建时间早于某一时间的广告，格式为 YYYY-MM-DD HH:MM:SS（UTC 时区）。
	// 建议：为确保任务的成功执行及速度，创建时间的筛选范围建议不超过 6 个月。
	CreationFilterEndTime model.DateTime `json:"creation_filter_end_time,omitzero"`
	// ModifiedAfter 筛选修改时间晚于某一时间的广告，格式为 YYYY-MM-DD HH:MM:SS（UTC 时区）。
	// 建议：为确保任务的成功执行及速度，修改时间的筛选范围建议不超过 6 个月。
	ModifiedAfter model.DateTime `json:"modified_after,omitzero"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
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

// GetResponse Get Upgraded Smart+ Ads API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// List 符合条件的广告列表，请求参数未设置fields字段或者fields中包含以下子字段时会返回相关字段
	// 注意：返回中不包含通过 TikTok Shop 创建的商品总交易额最大化广告（Product GMV max ads）和直播总交易额最大化广告（LIVE GMV max ads）。通过 API 或 TikTok 广告管理平台创建的购物广告则不受影响。
	List []Ad `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
