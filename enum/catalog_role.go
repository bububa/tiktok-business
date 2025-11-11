package enum

	// CatalogRole 商务中心用户对于商品库的权限。
type CatalogRole string

	const (
	// CatalogRole_ADMIN：商品库管理。可以通过自然内容和广告内容推广商品，也可以修改商品和商品库。
	CatalogRole_ADMIN CatalogRole = "ADMIN"
	// CatalogRole_AD_PROMOTE：广告推广。使用商品库创建广告。
	CatalogRole_AD_PROMOTE CatalogRole = "AD_PROMOTE"
)
