package enum

// CatalogInsightFilterType 筛选条件的类型。
type CatalogInsightFilterType string

const (
	// CatalogInsightFilterType_CATEGORY_ID：商品类目 ID。
	CatalogInsightFilterType_CATEGORY_ID CatalogInsightFilterType = "CATEGORY_ID"
	// CatalogInsightFilterType_BRAND：商品品牌名称。
	CatalogInsightFilterType_BRAND CatalogInsightFilterType = "BRAND"
	// CatalogInsightFilterType_AVAILABILITY：商品库存状态。
	CatalogInsightFilterType_AVAILABILITY CatalogInsightFilterType = "AVAILABILITY"
)
