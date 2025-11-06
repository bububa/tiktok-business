package order

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取 TTCM 订单 API Request
type GetRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// QueryType 查询类型。您需为所选查询类型传入对应的查询参数。否则，将出现报错。
	// 枚举值：
	// ORDER_ID ：通过工作流程 2.0 订单的 ID 进行查询。您需同时传入 order_ids 。
	// CAMPAIGN_CODE ：通过与工作流程 2.0 订单绑定的推广系列代码进行查询。您需同时传入 campaign_code 。
	// INVITE_LINK ：通过与工作流程 2.0 订单绑定的邀请链接进行查询。您需同时传入 invite_link 。
	QueryType enum.TCMOrderQueryType `json:"query_type,omitempty"`
	// OrderIDs 当query_type为ORDER_ID时必填。
	// 您想要筛选的工作流程 2.0 订单的 ID 列表。
	// 最大数量：200。
	// 您可通过/tcm/order/create/v2/获取工作流程 2.0 订单的ID。
	OrderIDs []string `json:"order_ids,omitempty"`
	// CampaignCode 当 query_type 为 CAMPAIGN_CODE 时必填。
	// 您想要筛选的工作流程 2.0 推广系列代码。
	// 您可通过/tcm/order/create/v2/ 获取工作流程 2.0 推广系列代码。
	CampaignCode string `json:"campaign_code,omitempty"`
	// InviteLink 当 query_type 为 INVITE_LINK 时必填。
	// 您想要筛选的工作流程 2.0 邀请链接。
	// 您可通过/tcm/order/create/v2/获取工作流程 2.0 邀请链接。
	InviteLink string `json:"invite_link,omitempty"`
	// Fields 您想要返回中显示的订单信息。
	// 枚举值：order_id，handle_name，campaign_code，invite_link，video_ids，current_video_count，total_video_limit，per_creator_video_limit 。
	// 默认值：["order_id", "handle_name", "campaign_code", "invite_link", "video_ids", "current_video_count", "total_video_limit", "per_creator_video_limit"]。
	Fields []string `json:"fields,omitempty"`
	// Page 当前页数。
	// 默认值：1。
	Page int `json:"page,omitempty"`
	// PageSize  分页大小。
	// 默认值：10。
	// 取值范围：1-20。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("tcm_account_id", r.TCMAccountID)
	values.Set("query_type", string(r.QueryType))
	if len(r.OrderIDs) > 0 {
		values.Set("order_ids", string(util.JSONMarshal(r.OrderIDs)))
	}
	if r.CampaignCode != "" {
		values.Set("campaign_code", r.CampaignCode)
	}
	if r.InviteLink != "" {
		values.Set("invite_link", r.InviteLink)
	}
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
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

// GetResponse 获取 TTCM 订单 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Orders 订单的数据列表
	Orders []Order `json:"orders,omitempty"`
}
