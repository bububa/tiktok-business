package enum

// ProductSource 商品来源
type ProductSource string

const (
	// ProductSource_CATALOG ：商品库。
	ProductSource_CATALOG ProductSource = "CATALOG"
	// ProductSource_STORE：TikTok Shop 店铺，或 TikTok 橱窗。
	ProductSource_STORE ProductSource = "STORE"
	// ProductSource_SHOWCASE（TikTok 橱窗）
	ProductSource_SHOWCASE ProductSource = "SHOWCASE"
	// ProductSource_UNSET
	ProductSource_UNSET ProductSource = "UNSET"
)
