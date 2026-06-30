package enum

// ProductCardType The product card display type.
type ProductCardType string

const (
	// ProductCardType_PRODUCT_CARD: Product Card.
	ProductCardType_PRODUCT_CARD ProductCardType = "PRODUCT_CARD"
	// ProductCardType_PRODUCT_TILE: Product Tiles.
	ProductCardType_PRODUCT_TILE ProductCardType = "PRODUCT_TILE"
	// ProductCardType_PRODUCT_INFO_CARD: Info Card.
	ProductCardType_PRODUCT_INFO_CARD ProductCardType = "PRODUCT_INFO_CARD"
	// ProductCardType_PRODUCT_SHOWCASE_TILE: Showcase Tiles.
	ProductCardType_PRODUCT_SHOWCASE_TILE ProductCardType = "PRODUCT_SHOWCASE_TILE"
	// ProductCardType_ANCHOR: Anchor.
	ProductCardType_ANCHOR ProductCardType = "ANCHOR"
	// ProductCardType_CAROUSEL_LABEL: Carousel label.
	ProductCardType_CAROUSEL_LABEL ProductCardType = "CAROUSEL_LABEL"
)
