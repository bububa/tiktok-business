package enum

// AssetType 资产类型。
type AssetType string

const (
	// AssetType_ADVERTISER：广告账户（广告主账号）。
	AssetType_ADVERTISER AssetType = "ADVERTISER"
	// AssetType_CATALOG：商品库。
	AssetType_CATALOG AssetType = "CATALOG"
	// AssetType_TIKTOK_SHOP：TikTok Shop（商店）。
	AssetType_TIKTOK_SHOP AssetType = "TIKTOK_SHOP"
	// AssetType_PIXEL：Pixel。
	AssetType_PIXEL AssetType = "PIXEL"
	// AssetType_LEAD：线索。
	AssetType_LEAD AssetType = "LEAD"
	// AssetType_TT_ACCOUNT：TikTok 账号。此资产类型对应 BC_AUTH_TT身份，详情见认证身份
	AssetType_TT_ACCOUNT AssetType = "TT_ACCOUNT"
	// AssetType_STOREFRANT
	AssetType_STOREFRANT AssetType = "STOREFRANT"
)
