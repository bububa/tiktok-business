package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Changelog 商务中心的日志
type Changelog struct {
	// Time 改动发生的时间，格式为 YYYY-MM-DD HH:MM:SS（UTC+0）
	Time model.DateTime `json:"time,omitzero"`
	// ActivityType 日志类型。
	// 枚举值：
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
	ActivityType enum.BcChangelogActivityType `json:"activity_type,omitempty"`
	// OperatorID 操作人的用户 ID
	OperatorID string `json:"operator_id,omitempty"`
	// ActivityLog 日志详情
	ActivityLog string `json:"activity_log,omitempty"`
}
