package enum

// MusicSearchType 音乐搜索类型
type MusicSearchType string

const (
	// MusicSearchType_SEARCH_BY_KEYWORD：通过关键词进行模糊搜索。您需同时传入 keyword。 需注意不支持使用本搜索类型搜索用户自定义上传的音乐。
	MusicSearchType_SEARCH_BY_KEYWORD MusicSearchType = "SEARCH_BY_KEYWORD"
	// MusicSearchType_SEARCH_BY_RECOMMEND：
	// music_scene设置为CAROUSEL_ADS时：根据指定的图片URL获取推荐音乐。您需同时传入 image_urls 。最多返回500首推荐音乐。需注意不支持使用本搜索类型搜索用户自定义上传的音乐。
	// music_scene 设置为 CATALOG_CAROUSEL 时：根据指定的商品获取推荐音乐。您需同时传入catalog_id 和 catalog_authorized_bc_id 。此外，您需同时选择以下三种设置之一。 需注意不支持使用本搜索类型搜索用户自定义上传的音乐。
	// 若想为商品库的所有商品获取推荐音乐，不可传入 product_set_id ， item_group_ids 以及 sku_ids 字段。
	// 若想为商品系列中的商品获取推荐音乐，需传入 product_set_id 和 item_group_ids 其中之一。
	// 若想为特定商品获取推荐音乐，需传入 sku_ids 。
	MusicSearchType_SEARCH_BY_RECOMMEND MusicSearchType = "SEARCH_BY_RECOMMEND"
	// MusicSearchType_SEARCH_BY_LIKED：搜索收藏列表中的音乐。最多返回1,000首推荐收藏列表中的音乐。
	MusicSearchType_SEARCH_BY_LIKED MusicSearchType = "SEARCH_BY_LIKED"
	// MusicSearchType_SEARCH_BY_HISTORY：搜索历史列表中的音乐。
	MusicSearchType_SEARCH_BY_HISTORY MusicSearchType = "SEARCH_BY_HISTORY"
	// MusicSearchType_SEARCH_BY_MUSIC_ID：通过音乐ID搜索音乐。
	MusicSearchType_SEARCH_BY_MUSIC_ID MusicSearchType = "SEARCH_BY_MUSIC_ID"
	// MusicSearchType_SEARCH_BY_SOURCE：通过音乐来源搜索音乐。您需同时传入sources。需注意您只能筛选出用户自定义上传的、可用于轮播广告的音乐。
	MusicSearchType_SEARCH_BY_SOURCE MusicSearchType = "SEARCH_BY_SOURCE"
)
