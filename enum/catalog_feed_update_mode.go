package enum

// CatalogFeedUpdateMode 更新模式。
type CatalogFeedUpdateMode string

const (
	// CatalogFeedUpdateMode_OVERWRITE：覆盖数据源。若历史商品在新的商品库中不存在，则历史商品将被删除。若存在新商品，将在商品库中添加新商品；若商品已存在，则会用上传的商品信息更新原有的商品信息。
	CatalogFeedUpdateMode_OVERWRITE CatalogFeedUpdateMode = "OVERWRITE"
	// CatalogFeedUpdateMode_INCREMENTAL：增量式更新。若历史商品在新的商品库中不存在，历史商品不会被删除。若存在新商品，将在商品库中添加新商品；若商品已存在，则会用上传的商品信息更新原有的商品信息。
	CatalogFeedUpdateMode_INCREMENTAL CatalogFeedUpdateMode = "INCREMENTAL"
)
