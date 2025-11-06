package order

// Order TTCM订单详情
type Order struct {
	// OrderID 订单ID
	OrderID string `json:"order_id,omitempty"`
	// HandleName 创作者用户名
	HandleName string `json:"handle_name,omitempty"`
	// VideoIDs 该订单下已上传的视频 ID 列表
	VideoIDs []string `json:"video_ids,omitempty"`
	// CampaignCode 与该订单绑定的推广系列代码
	CampaignCode string `json:"campaign_code,omitempty"`
	// InviteLink 与该订单绑定的邀请链接
	InviteLink string `json:"invite_link,omitempty"`
	// Quota 与该订单的视频配额有关的信息
	Quota *Quota `json:"quota,omitempty"`
}

// Quota 与该订单的视频配额有关的信息
type Quota struct {
	// CurrentVideoCount 该订单下已上传的视频数量
	CurrentVideoCount int `json:"current_video_count,omitempty"`
	// PerCreatorVideoLimit 该订单下允许单个创作者上传的最大视频数量。
	// 例如，若您使用 /tcm/order/create/v2/创建工作流程 2.0 订单时，将 video_number 设置 为 4 或传入了大小为 4 的 deliverables 字段，则本字段的值将为 4 。
	PerCreatorVideoLimit int `json:"per_creator_video_limit,omitempty"`
	// TotalVideoLimit 与该订单绑定的推广系列下允许上传的最大视频数量。
	// 例如，若您使用/tcm/order/create/v2/创建工作流程 2.0 订单时，将 video_number 设置为 4 或传入了大小为 4 的 deliverables 字段，则本字段的值将为 800 。
	TotalVideoLimit int `json:"total_video_limit,omitempty"`
}
