package tool

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CampaignLabelGetRequest 获取广告账号的推广系列标签 API Request
type CampaignLabelGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignLabelIDs 推广系列标签 ID 列表。
	// 最大数量：50。
	// 每个标签 ID 必须为 19 位数字字符串。
	// 注意：本筛选条件仅支持在同步基础报表中使用。
	CampaignLabelIDs []string `json:"campaign_label_ids,omitempty"`
	// CampaignLabelNames 推广系列标签名称列表。
	// 支持模糊匹配。
	// 最大数量：10。
	CampaignLabelNames []string `json:"campaign_label_names,omitempty"`
	// CampaignLabelTypes 推广系列标签类型列表。
	// 枚举值：
	// GENERAL：通用标签。
	// MARKETING_EVENT：营销事件标签。
	CampaignLabelTypes []enum.CampaignLabelType `json:"campaign_label_types,omitempty"`
	// Page 当前页数。
	// 默认值：1。
	// 取值范围：≥1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10。
	// 取值范围：1-1,000。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *CampaignLabelGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.CampaignLabelIDs) > 0 {
		values.Set("campaign_label_ids", string(util.JSONMarshal(r.CampaignLabelIDs)))
	}
	if len(r.CampaignLabelNames) > 0 {
		values.Set("campaign_label_names", string(util.JSONMarshal(r.CampaignLabelNames)))
	}
	if len(r.CampaignLabelTypes) > 0 {
		values.Set("campaign_label_tppes", string(util.JSONMarshal(r.CampaignLabelTypes)))
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

// CampaignLabelGetResponse 获取广告账号的推广系列标签 API Response
type CampaignLabelGetResponse struct {
	model.BaseResponse
	Data *CampaignLabelGetResult `json:"data,omitempty"`
}

type CampaignLabelGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 广告账号内推广系列标签的详细信息
	List []CampaignLabel `json:"list,omitempty"`
}

// CampaignLabel 广告账号内推广系列标签的详细信息
type CampaignLabel struct {
	// CampaignLabelID 推广系列标签 ID
	CampaignLabelID string `json:"campaign_label_id,omitempty"`
	// CampaignLabelName 推广系列标签名称
	CampaignLabelName string `json:"campaign_label_name,omitempty"`
	// CampaignLabelType 推广系列标签类型。
	// 枚举值：
	// GENERAL：通用标签。
	// MARKETING_EVENT：营销事件标签
	CampaignLabelType enum.CampaignLabelType `json:"campaign_label_type,omitempty"`
	// CampaignLabelColor 推广系列标签颜色
	CampaignLabelColor string `json:"campaign_label_color,omitempty"`
	// CreateTime 推广系列标签创建时间，格式为 YYYY-MM-DD HH:MM:SS (UTC+0)。
	// 示例：2025-01-01 00:00:01。
	CreateTime model.DateTime `json:"create_time,omitzero"`
}
