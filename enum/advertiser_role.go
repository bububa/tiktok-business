package enum

// AdvertiserRole 商务中心用户对于广告账户的权限。
type AdvertiserRole string

const (
	// AdvertiserRole_ADMIN：管理员。管理广告账号财务、设置和权限。 创建和编辑广告推广系列并查看投放效果数据。
	AdvertiserRole_ADMIN AdvertiserRole = "ADMIN"
	// AdvertiserRole_OPERATOR：操作员。创建和编辑广告推广系列并查看投放效果数据。
	AdvertiserRole_OPERATOR AdvertiserRole = "OPERATOR"
	// AdvertiserRole_ANALYST：分析师。查看广告和投放效果数据。
	AdvertiserRole_ANALYST AdvertiserRole = "ANALYST"
)
