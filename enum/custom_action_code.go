package enum

// CustomActionCode 用于筛选出受众再营销对象的自定义行为。
type CustomActionCode string

const (
	// CustomActionCode_VIEW_PRODUCT：受众浏览过商品。
	CustomActionCode_VIEW_PRODUCT CustomActionCode = "VIEW_PRODUCT"
	// CustomActionCode_ADD_TO_CART : 受众将商品加过购物车。
	CustomActionCode_ADD_TO_CART CustomActionCode = "ADD_TO_CART"
	// CustomActionCode_PURCHASE : 受众购买过商品。
	CustomActionCode_PURCHASE CustomActionCode = "PURCHASE"
)
