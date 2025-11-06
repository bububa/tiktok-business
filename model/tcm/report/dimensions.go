package report

// Dimensions 维度数据
type Dimensions struct {
	// TCMAccountID TTCM 账户 ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// OrderID 订单 ID
	OrderID string `json:"order_id,omitempty"`
	// OrderType 订单类型，代表订单的创建方式。
	// 枚举值：API_ORDER_V2 （订单通过 /tcm/order/create/v2/接口创建）
	OrderType string `json:"order_type,omitempty"`
	// VideoID 视频 ID
	VideoID string `json:"video_id,omitempty"`
	// HandleName TikTok创作者的用户名
	HandleName string `json:"handle_name,omitempty"`
	// VideoFirstReleaseTime 视频首次发布的时间，格式：YYYY-MM-DD（UTC+0 时间）。 示例： 2020-06-10
	VideoFirstReleaseTime string `json:"video_first_release_time,omitzero"`
}
