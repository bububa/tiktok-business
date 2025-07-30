package event

import (
	"time"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CRMListRequest 获取 CRM 事件组列表 API Request
type CRMListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Name 用于筛选结果的 CRM 事件组名称。
	// 长度限制：1-40 字符。
	// 仅支持精确匹配。
	// 若未传入本字段，将返回广告账户中的所有 CRM 事件组。
	// 若您传入的名称没有匹配结果，返回结果将为空。
	Name string `json:"name,omitempty"`
	// EventSetIDs 用于筛选结果的 CRM 事件组 ID 列表。
	// 最大数量：50。
	// 仅支持精确匹配。
	// 若未传入本字段，将返回广告账户中的所有 CRM 事件组。
	// 仅返回与您在该字段指定的 ID 匹配的 CRM 事件组，无匹配结果的 ID 将不会返回对应的 CRM 事件组。
	// 若您同时传入event_set_ids和name，则将在匹配到使用所指定的名称，且该事件组的 ID 在指定的 CRM 事件组 ID 列表中时，获取一条结果，或未匹配到 CRM 事件组时，获取0条结果。
	// 注意：单个广告账户中允许的 CRM 事件组最大数量为 50。
	EventSetIDs []string `json:"event_set_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r *CRMListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.Name != "" {
		values.Set("name", r.Name)
	}
	if len(r.EventSetIDs) > 0 {
		values.Set("event_set_ids", string(util.JSONMarshal(r.EventSetIDs)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

type CRMListResponse struct {
	model.BaseResponse
	Data struct {
		// CRMEventSets 广告账户中匹配到的 CRM 事件组列表
		CRMEventSets []CRMEventSet `json:"crm_event_sets,omitempty"`
	} `json:"data"`
}

// CRMEventSet 广告账户中匹配到的 CRM 事件组
type CRMEventSet struct {
	// EventSetID CRM 事件组 ID
	EventSetID string `json:"event_set_id,omitempty"`
	// Name CRM 事件组名称
	Name string `json:"name,omitempty"`
	// CreateTime 该 CRM 事件组的创建时间，以ISO 8601格式表示。
	// 示例：”2023-04-15T00:23:38Z”。
	CreateTime time.Time `json:"create_time,omitzero"`
}
