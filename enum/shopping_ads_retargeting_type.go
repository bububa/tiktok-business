package enum

// ShoppingAdsRetargetingType  购物广告受众再营销类型
type ShoppingAdsRetargetingType string

const (
	// ShoppingAdsRetargetingType_LAB1 : 对浏览过商品或者将商品加过购物车，但未购买过商品的受众进行再营销。
	ShoppingAdsRetargetingType_LAB1 ShoppingAdsRetargetingType = "LAB1"
	// ShoppingAdsRetargetingType_LAB2：对将商品加过购物车,但未购买过商品的受众进行再营销。
	ShoppingAdsRetargetingType_LAB2 ShoppingAdsRetargetingType = "LAB2"
	// ShoppingAdsRetargetingType_LAB3 : 使用自定义条件组合对受众进行再营销。
	ShoppingAdsRetargetingType_LAB3 ShoppingAdsRetargetingType = "LAB3"
	// ShoppingAdsRetargetingType_OFF : 不进行受众再营销。
	ShoppingAdsRetargetingType_OFF ShoppingAdsRetargetingType = "OFF"
)
