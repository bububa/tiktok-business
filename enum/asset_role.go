package enum

// AssetRole 您根据资产类型为成员分配的资产角色。
type AssetRole string

const (
	// AssetRole_ADVERTISER_ROLE_ADMIN（广告账户管理员）：
	AssetRole_ADVERTISER_ROLE_ADMIN AssetRole = "ADVERTISER_ROLE_ADMIN"
	// 管理广告账户：管理广告账户的设置、资金、权限。
	// 管理推广系列：创建、编辑广告。
	// 获取广告表现：获取广告报表，查看广告。
	// AssetRole_ADVERTISER_ROLE_OPERATOR（广告账户操作员）：
	AssetRole_ADVERTISER_ROLE_OPERATOR AssetRole = "ADVERTISER_ROLE_OPERATOR"
	// 管理推广系列：创建、编辑广告。
	// 获取广告表现：获取报表，查看广告。
	// AssetRole_ADVERTISER_ROLE_ANALYST（广告账户分析师）：
	// 获取广告表现：获取报表，查看广告。
	AssetRole_ADVERTISER_ROLE_ANALYST AssetRole = "ADVERTISER_ROLE_ANALYST"
)
