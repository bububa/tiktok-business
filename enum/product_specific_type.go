package enum

// ProductSpecificType 选择商品的维度。
type ProductSpecificType string

const (
	// ProductSpecificType_ALL：允许 TikTok 从所有商品中动态选择。
	ProductSpecificType_ALL ProductSpecificType = "ALL"
	// ProductSpecificType_PRODUCT_SET：请选择一个商品系列。TikTok 将动态选择该商品系列中的商品。
	ProductSpecificType_PRODUCT_SET ProductSpecificType = "PRODUCT_SET"
	// ProductSpecificType_CUSTOMIZED_PRODUCTS：选择自定义数量的特定商品。
	ProductSpecificType_CUSTOMIZED_PRODUCTS ProductSpecificType = "CUSTOMIZED_PRODUCTS"
)
