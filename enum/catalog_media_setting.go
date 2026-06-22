package enum

// CatalogMediaSetting The types of creatives from your E-commerce catalog to use in the ad.
type CatalogMediaSetting string

const (
	// CatalogMediaSetting_VIDEO: Video.
	CatalogMediaSetting_VIDEO CatalogMediaSetting = "VIDEO"
	// CatalogMediaSetting_IMAGE: Image.
	CatalogMediaSetting_IMAGE CatalogMediaSetting = "IMAGE"
	// CatalogMediaSetting_TEMPLATE_VIDEO: Video templates. If you include this value in catalog_media_settings, you can optionally specify catalog_template_video_id at the same time.
	// If catalog_template_video_id is not specified, all video templates and video URLs from your catalog will be used to generate the ad.
	// If catalog_template_video_id is specified, the selected video template will be used to generate the ad.
	CatalogMediaSetting_TEMPLATE_VIDEO CatalogMediaSetting = "TEMPLATE_VIDEO"
	// CatalogMediaSetting_MULTI_SHOW: Multi-Show. Multi-Show Experience is an auto-play video carousel experience designed to drive user exploration and engagement across a breadth of personally-relevant title offerings within your content library. Multi-Show Experience consists of 4 video tiles delivered from your catalog. Each tile's clip plays sequentially (starting in the upper left-hand corner), and offers you the ability to assign a corresponding title-specific landing page.
	CatalogMediaSetting_MULTI_SHOW CatalogMediaSetting = "MULTI_SHOW"
)
