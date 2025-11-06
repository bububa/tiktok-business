package enum

// TCMOrderQueryType 查询类型
type TCMOrderQueryType string

const (
	// TCMOrderQueryType_ORDER_ID ：通过工作流程 2.0 订单的 ID 进行查询。您需同时传入 order_ids 。
	TCMOrderQueryType_ORDER_ID TCMOrderQueryType = "ORDER_ID"
	// TCMOrderQueryType_CAMPAIGN_CODE ：通过与工作流程 2.0 订单绑定的推广系列代码进行查询。您需同时传入 campaign_code 。
	TCMOrderQueryType_CAMPAIGN_CODE TCMOrderQueryType = "CAMPAIGN_CODE"
	// TCMOrderQueryType_INVITE_LINK ：通过与工作流程 2.0 订单绑定的邀请链接进行查询。您需同时传入 invite_link 。
	TCMOrderQueryType_INVITE_LINK TCMOrderQueryType = "INVITE_LINK"
)
