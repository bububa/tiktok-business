package gmvmax

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取 GMV Max 推广系列列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Fields 希望返回的字段集合。
	// 不传入本字段时，默认返回所有字段。
	// 枚举值参见返回中list下包含的字段。
	Fields []string `json:"fields,omitempty"`
	// Filtering 筛选条件
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 取值范围：≥ 1。
	// 默认值: 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-100。
	// 默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 筛选条件
type GetFilter struct {
	// GMVMaxPromotionTypes GMV Max 推广系列类型。
	// 枚举值：
	// PRODUCT_GMV_MAX: 商品 GMV Max 推广系列。
	// LIVE_GMV_MAX: 直播 GMV Max 推广系列。
	GMVMaxPromotionTypes []enum.GMVMaxPromotionType `json:"gmv_max_promotion_types,omitempty"`
	// StoreIDs TikTok Shop ID 列表。
	// 最大数量：10。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreIDs []string `json:"store_ids,omitempty"`
	// CampaignIDs 推广系列 ID 列表。
	// 最大数量：100
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// CampaignName 推广系列名称
	CampaignName string `json:"campaign_name,omitempty"`
	// PrimaryStatus 一级状态。
	// 枚举值：
	// STATUS_DELIVERY_OK: 已生效。
	// STATUS_DISABLE: 已暂停。
	// STATUS_DELETE: 已删除。
	PrimaryStatus enum.PrimaryStatus `json:"primary_status,omitempty"`
	// CreationFilterStartTime 推广系列最早创建时间，格式为 YYYY-MM-DD HH:MM:SS (UTC 时区)。将返回创建时间晚于此时间的推广系列。
	// 示例： 2025-01-01 00:00:01。
	// 若 creation_filter_start_time 和 creation_filter_end_time 不传入，将返回任意时间创建的所有 GMV Max 推广系列。
	// 注意：为确保任务成功执行及任务速度不受影响，创建时间的范围建议在6个月以内。
	CreationFilterStartTime string `json:"creation_filter_start_time,omitempty"`
	// CreationFilterEndTime 推广系列最晚创建时间，格式为 YYYY-MM-DD HH:MM:SS (UTC 时区)。将返回创建时间先于此时间的推广系列。
	// 示例： 2025-02-01 00:00:01。
	// 若 creation_filter_start_time 和 creation_filter_end_time 不传入，将返回任意时间创建的所有 GMV Max 推广系列。
	// 注意：为确保任务成功执行及任务速度不受影响，创建时间的筛选范围建议在6个月以内
	CreationFilterEndTime string `json:"creation_filter_end_time,omitempty"`
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

// GetResponse 获取 GMV Max 推广系列列表 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 推广系列列表
	List []Campaign `json:"list,omitempty"`
}
