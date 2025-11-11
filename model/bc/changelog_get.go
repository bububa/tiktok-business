package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ChangelogGetRequest 获取商务中心的日志 API Request
type ChangelogGetRequest struct {
	// BcID 商务中心ID。不传则默认返回用户所有的商务中心列表，传入时返回指定商务中心账户
	BcID string `json:"bc_id,omitempty"`
	// Filtering 筛选条件
	Filtering *ChangelogGetFilter `json:"filtering,omitempty"`
	// Lang 日志的语言。
	// 默认值：en。
	// 若想获取枚举值，可查看lang枚举值
	Lang string `json:"lang,omitempty"`
	// SortField 排序字段。
	// 枚举值：
	// operation_time：按改动发生的时间排序，对应返回中对象数组 changelog_list 中的 time 参数。
	// 默认值： operation_time
	SortField string `json:"sort_field,omitempty"`
	// SortType 排序顺序。
	// 枚举值：
	// DESC：降序。
	// sort_field 为 operation_time 或未传入时，本值代表按改动发生时间从近到远排序。
	// ASC：升序。
	// sort_field 为 operation_time 或未传入时，本值代表按改动发生时间从远到近排序。
	// 默认值：DESC
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Page 当前页数，默认值: 1，取值范围: ≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，默认值: 10，取值范围: 1-50
	PageSize int `json:"page_size,omitempty"`
}

type ChangelogGetFilter struct {
	// StartDate 查询起始时间，格式为YYYY-MM-DD（UTC+0）。
	// 您需同时传入start_date 和 end_date 或两者均不传入。
	// 若 start_date 和 end_date 均不传入，将默认返回过去 7 天的日志。
	// 推荐查询的时间范围：6 个月。
	// 示例：2025-01-01。
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询终止时间，格式为YYYY-MM-DD（UTC+0）。
	// 您需同时传入start_date 和 end_date 或两者均不传入。
	// 若 start_date 和 end_date 均不传入，将默认返回过去 7 天的日志。
	// 推荐查询的时间范围：6 个月
	EndDate string `json:"end_date,omitempty"`
	// ActivityType 查询的日志类型。
	// 枚举值：
	// ALL: 所有类型。
	// USER: 用户。主要与成员和组织相关的改动，包括但不仅限于以下场景：
	// 邀请用户成为商务中心的成员
	// 接受邀请成为商务中心的成员
	// 将成员从商务中心中移除
	// ACCOUNT：账号。主要与广告账号的创建和权限管理相关的改动，包括但不仅限于以下场景：
	// 创建广告账号
	// 为成员分配广告账号的权限
	// 取消成员的广告账号权限
	// 与合作伙伴分享广告账号的权限
	// 取消合作伙伴的广告账号权限
	// ASSET：资产。主要与资产相关的改动，包括但不仅限于以下场景：
	// 创建或删除商品库
	// 请求添加 TikTok Shop 或同意添加 TikTok Shop 的请求
	// 为成员分配资产的权限
	// 取消成员的资产权限
	// BUSINESS: 商务中心。主要与商务中心相关的改动，包括但不仅限于以下场景：
	// 创建商务中心
	// 修改商务中心的设置，包括名称和公司名称
	// 默认值：ALL。
	ActivityType enum.BcChangelogActivityType `json:"activity_type,omitempty"`
}

// Encode implements GetRequest
func (r *ChangelogGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Lang != "" {
		values.Set("lang", r.Lang)
	}
	if r.SortField != "" {
		values.Set("sort_field", r.SortField)
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
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

// ChangelogGetResponse 获取商务中心的日志 API Response
type ChangelogGetResponse struct {
	model.BaseResponse
	Data *ChangelogGetResult `json:"data,omitempty"`
}

type ChangelogGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ChangelogList 日志相关信息
	ChangelogList []Changelog `json:"changelog_list,omitempty"`
}
