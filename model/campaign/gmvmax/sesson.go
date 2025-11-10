package gmvmax

import "github.com/bububa/tiktok-business/enum"

// Session 最大投放量时段或创意作品加热时段
type Session struct {
	// CampaignID 商品 GMV Max 推广系列的 ID
	CampaignID string `json:"campaign_id,omitempty"`
	// BidType 时段类型。
	// 枚举值:
	// NO_BID: 最大投放量时段。
	// CREATIVE_NO_BID: 创意作品加热时段
	BidType enum.GMVMaxBidType `json:"bid_type,omitempty"`
	// SessionID 商品 GMV Max 推广系列中商品的最大投放量时段 ID 或视频的创意作品加热时段 ID。
	// 当 bid_type 为 NO_BID，此字段代表最大投放量时段。
	// 当 bid_type 为 CREATIVE_NO_BID，此字段代表创意作品加热时段。
	SessionID string `json:"session_id,omitempty"`
	// Budget 最大投放量日预算或创意作品加热日预算。
	// 最大投放量使用额外的预算，与目标 ROI 优化的预算分开，并按照更可预测的广告支出节奏进行投放。您的目标 ROI 预算不会用于采用最大投放量的商品。
	// 商品 GMV Max 推广系列中的创意作品加热功能允许商家手动为特定视频配置额外日预算，用于视频推广
	Budget float64 `json:"budget,omitempty"`
	// ProductList 有关特定商品的详情
	ProductList []SessionProduct `json:"product_list,omitempty"`
	// ScheduleType 最大投放量或创意作品加热模式的排期类型。
	// 枚举值:
	// SCHEDULE_FROM_NOW: 从 schedule_start_time 开始，为该商品持续启用最大投放量模式或创意作品加热模式，直到推广系列的排期结束时间。
	// SCHEDULE_START_END: 在 schedule_start_time 和 schedule_end_time 排期期间为该商品启用最大投放量模式或创意作品加热模式。
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// ScheduleStartTime 最大投放量模式或创意作品加热模式的开始时间，格式为 YYYY-MM-DD HH:MM:SS (UTC+0)
	ScheduleStartTime string `json:"schedule_start_time,omitempty"`
	// ScheduleEndTime 最大投放量模式或创意作品加热模式的结束时间，格式为 YYYY-MM-DD HH:MM:SS (UTC+0)
	ScheduleEndTime string `json:"schedule_end_type,omitempty"`
	// ItemID 仅当 bid_type 为 CREATIVE_NO_BID 时返回。
	// 启用创意作品加热的视频的 ID
	ItemID string `json:"item_id,omitempty"`
}

// SessionProduct 有关特定商品的详情
type SessionProduct struct {
	// SpuID 商品 SPU（标准化产品单元） ID
	SpuID string `json:"spu_id,omitempty"`
}
