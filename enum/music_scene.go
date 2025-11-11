package enum

// MusicScene 音乐的使用场景。
type MusicScene string

const (
	// MusicScene_CREATIVE_ASSET：创意工具场景。
	MusicScene_CREATIVE_ASSET MusicScene = "CREATIVE_ASSET"
	// MusicScene_CAROUSEL_ADS：非购物广告类型的轮播广告场景。
	MusicScene_CAROUSEL_ADS MusicScene = "CAROUSEL_ADS"
	// MusicScene_CATALOG_CAROUSEL：购物广告类型的轮播广告场景。
	MusicScene_CATALOG_CAROUSEL MusicScene = "CATALOG_CAROUSEL"
)
