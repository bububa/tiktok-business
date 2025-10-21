package enum

	// ProductActiveStatus 商品的上架状态。
type ProductActiveStatus string

const (
	// ProductActiveStatus_ACTIVATED：已上架。
	ProductActiveStatus_ACTIVATED ProductActiveStatus = "ACTIVATED"
	// ProductActiveStatus_DEACTIVATED：已下架。
  ProductActiveStatus_DEACTIVATED ProductActiveStatus = "DEACTIVATED"
)
