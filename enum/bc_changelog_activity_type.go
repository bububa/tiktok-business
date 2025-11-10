package enum

// BcChangelogActivityType 日志类型。
type BcChangelogActivityType string

const (
	// BcChangelogActivityType_ALL: 所有类型。
	BcChangelogActivityType_ALL BcChangelogActivityType = "ALL"
	// BcChangelogActivityType_USER: 用户。主要与成员和组织相关的改动，包括但不仅限于以下场景：
	// 邀请用户成为商务中心的成员
	// 接受邀请成为商务中心的成员
	// 将成员从商务中心中移除
	BcChangelogActivityType_USER BcChangelogActivityType = "USER"
	// BcChangelogActivityType_ACCOUNT：账号。主要与广告账号的创建和权限管理相关的改动，包括但不仅限于以下场景：
	// 创建广告账号
	// 为成员分配广告账号的权限
	// 取消成员的广告账号权限
	// 与合作伙伴分享广告账号的权限
	// 取消合作伙伴的广告账号权限
	BcChangelogActivityType_ACCOUNT BcChangelogActivityType = "ACCOUNT"
	// BcChangelogActivityType_ASSET：资产。主要与资产相关的改动，包括但不仅限于以下场景：
	// 创建或删除商品库
	// 请求添加 TikTok Shop 或同意添加 TikTok Shop 的请求
	// 为成员分配资产的权限
	// 取消成员的资产权限
	BcChangelogActivityType_ASSET BcChangelogActivityType = "ASSET"
	// BcChangelogActivityType_BUSINESS: 商务中心。主要与商务中心相关的改动，包括但不仅限于以下场景：
	// 创建商务中心
	// 修改商务中心的设置，包括名称和公司名称
	BcChangelogActivityType_BUSINESS BcChangelogActivityType = "BUSINESS"
)
