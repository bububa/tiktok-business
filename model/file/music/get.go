package music

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取音乐列表 API Request
type GetRequest struct {
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// MusicScene 音乐的使用场景。
	// 枚举值：
	// CREATIVE_ASSET：创意工具场景。
	// CAROUSEL_ADS：非购物广告类型的轮播广告场景。
	// CATALOG_CAROUSEL：购物广告类型的轮播广告场景。
	// 默认值：CREATIVE_ASSET
	MusicScene enum.MusicScene `json:"music_scene,omitempty"`
	// SearchType 当 music_scene 设置为 CAROUSEL_ADS或 CATALOG_CAROUSEL时必填。
	// 当 music_scene 设置为 CREATIVE_ASSET 或未传入时不支持本字段。
	// 音乐搜索类型。
	// 枚举值：
	// SEARCH_BY_KEYWORD：通过关键词进行模糊搜索。您需同时传入 keyword。 需注意不支持使用本搜索类型搜索用户自定义上传的音乐。
	// SEARCH_BY_RECOMMEND：
	// music_scene设置为CAROUSEL_ADS时：根据指定的图片URL获取推荐音乐。您需同时传入 image_urls 。最多返回500首推荐音乐。需注意不支持使用本搜索类型搜索用户自定义上传的音乐。
	// music_scene 设置为 CATALOG_CAROUSEL 时：根据指定的商品获取推荐音乐。您需同时传入catalog_id 和 catalog_authorized_bc_id 。此外，您需同时选择以下三种设置之一。 需注意不支持使用本搜索类型搜索用户自定义上传的音乐。
	// 若想为商品库的所有商品获取推荐音乐，不可传入 product_set_id ， item_group_ids 以及 sku_ids 字段。
	// 若想为商品系列中的商品获取推荐音乐，需传入 product_set_id 和 item_group_ids 其中之一。
	// 若想为特定商品获取推荐音乐，需传入 sku_ids 。
	// SEARCH_BY_LIKED：搜索收藏列表中的音乐。最多返回1,000首推荐收藏列表中的音乐。
	// SEARCH_BY_HISTORY：搜索历史列表中的音乐。
	// SEARCH_BY_MUSIC_ID：通过音乐ID搜索音乐。
	// SEARCH_BY_SOURCE：通过音乐来源搜索音乐。您需同时传入sources。需注意您只能筛选出用户自定义上传的、可用于轮播广告的音乐。
	// 注意： 当本字段设置为 SEARCH_BY_RECOMMEND ，SEARCH_BY_LIKED 或 SEARCH_BY_MUSIC_ID，不支持分页，将忽略 page 和 page_size 参数。
	SearchType enum.MusicSearchType `json:"search_type,omitempty"`
	// Filtering 筛选条件
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围: ≥ 1。
	// page 和 page_size 的乘积不可超过 100,000。
	// 注意：当 search_type设置为 SEARCH_BY_RECOMMEND ，SEARCH_BY_LIKED 或 SEARCH_BY_MUSIC_ID，不支持分页，将忽略 page 和 page_size 参数
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 1,000。
	// 取值范围: 1-1,000。
	// page 和 page_size 的乘积不可超过 100,000。
	// 注意：
	// 当 search_type 设置为 SEARCH_BY_RECOMMEND ，SEARCH_BY_LIKED 或 SEARCH_BY_MUSIC_ID 时，不支持分页，将忽略 page 和 page_size 参数。
	// 当 search_type 设置为 SEARCH_BY_KEYWORD 或 SEARCH_BY_HISTORY 时，默认值为200，取值范围[1, 200]
	PageSize int `json:"page_size,omitempty"`
}

// GetFilter 筛选条件
type GetFilter struct {
	// Keyword 当 search_type 设置为 SEARCH_BY_KEYWORD 时必填。
	// 当 search_type 未设置为SEARCH_BY_KEYWORD 时不生效。
	// 用于搜索音乐的关键词
	Keyword string `json:"keyword,omitempty"`
	// ImageURLs 当music_scene设置为CAROUSEL_ADS且search_type 设置为 SEARCH_BY_RECOMMEND 时必填。
	// 当 search_type 未设置为 SEARCH_BY_RECOMMEND 时不生效。
	// 图片URL列表，用于获取推荐音乐。
	// 列表长度范围：[2, 35]。
	// 您需指定至少两个使用 /file/image/ad/upload/ 或 /file/image/ad/search/接口获取的图片URL。
	// 所指定的图片还需同时满足：
	// 规格要求：
	// 文件类型: JPG，JPEG，或PNG
	// 图片分辨率：最大为1242*2340 px 或 2340*1242 px
	// 图片宽高比：最大为 9:20 或 20:9
	// 文件大小：小于 50 MB
	// /file/image/ad/search/返回的图片对应 is_carousel_usable 的值为true。
	// 若您在创建非视频购物类型的轮播广告时使用了推荐音乐，您需在该广告中使用用于生成推荐音乐的图片
	ImageURLs []string `json:"image_urls,omitempty"`
	// MusicIDs 当 search_type 设置为 SEARCH_BY_MUSIC_ID 时必填。
	// 音乐ID列表。
	// 最大数量：100
	MusicIDs []string `json:"music_ids,omitempty"`
	// CatalogID music_scene 设置为 CATALOG_CAROUSEL 且 search_type 设置为 SEARCH_BY_RECOMMEND 时必填。
	// 商品库 ID
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogAuthorizedBcID music_scene 设置为 CATALOG_CAROUSEL 且 search_type 设置为 SEARCH_BY_RECOMMEND 时必填。
	// 拥有商品库权限的商务中心的ID。
	// 对于商务中心中的商品库，您必须传入商品库所属商务中心的ID
	CatalogAuthorizedBcID string `json:"catalog_authorized_bc_id,omitempty"`
	// ItemGroupIDs 当 music_scene 设置为 CATALOG_CAROUSEL ，search_type 设置为 SEARCH_BY_RECOMMEND，且您希望为从商品系列中动态选取商品的视频购物类型轮播广告获取推荐音乐时，您需传入 product_set_id 和 item_group_ids 其中之一。
	// 若您想为从整个商品库中动态选取商品的视频购物类型轮播广告获取推荐音乐，无需传入本字段。
	// 商品的SPU（标准化产品单元） ID列表。
	// 最大数量：20。
	ItemGroupIDs []string `json:"item_group_ids,omitempty"`
	// ProductSetID 当 music_scene 设置为 CATALOG_CAROUSEL， search_type 设置为SEARCH_BY_RECOMMEND，且您希望为从商品系列中动态选取商品的视频购物类型轮播广告获取推荐音乐时，您需传入 product_set_id 和 item_group_ids 其中之一。
	// 若您想为从整个商品库中动态选取商品的视频购物类型轮播广告获取推荐音乐，无需传入本字段。
	// 商品系列ID
	ProductSetID string `json:"product_set_id,omitempty"`
	// SkuIDs 当 music_scene 设置为 CATALOG_CAROUSEL ， search_type 设置为 SEARCH_BY_RECOMMEND ，且您希望为使用商品库中指定商品的视频购物类型轮播广告获取推荐音乐时，本字段必填。
	// 若您想为从整个商品库中动态选取商品的视频购物类型轮播广告获取推荐音乐，无需传入本字段。
	// SKU（库存量单位）ID列表。
	// 最大数量：20。
	// 最小数量：2
	SkuIDs []string `json:"sku_ids,omitempty"`
	// CarouselImageIndex 仅当 music_scene 设置为 CATALOG_CAROUSEL ， search_type 设置为SEARCH_BY_RECOMMEND，且所指定的商品已设置对应的附加图片列表时生效。
	// 用于指定视频购物类型的轮播广告中要使用的附加图片的索引。
	// 若未传入本字段，将在视频购物类型的轮播广告中使用商品库商品的主图。
	// 取值范围：[0, 9]。
	// 您通过本字段传入的数值代表每个商品库商品的附加图片列表（ additional_image_urls ） 中要使用的附加图片的位置。例如，若 carousel_image_index 设置为1，代表您将使用每个 additional_image_urls 的值中的第二个图片作为轮播广告中使用的图片。您可通过/catalog/product/get/返回的 additional_image_urls 字段获取商品对应的附加图片列表。
	// 若您指定的数值超过了商品对应的附加图片列表中附加图片 URL 的数量，将忽略本字段，使用商品的主图（image_url）。您可通过/catalog/product/get/返回的 image_url 字段获取商品对应的主图。
	CarouselImageIndex *int `json:"carousel_image_index,omitempty"`
	// MaterialIDs 当 music_scene 设置为 CAROUSEL_ADS 或 CATALOG_CAROUSEL时，不支持使用本字段。
	// 素材 ID列表。
	// 最大数量：100
	MaterialIDs []string `json:"material_ids,omitempty"`
	// Styles 当 music_scene 设置为 CAROUSEL_ADS或 CATALOG_CAROUSEL时，不支持使用本字段。
	// 音乐风格。
	// 最大数量：1
	Styles []string `json:"styles,omitempty"`
	// Sources 当 music_scene 设置为 CREATIVE_ASSET 或未传入时，本字段选填。
	// 当 music_scene 设置为 CAROUSEL_ADS或 CATALOG_CAROUSEL，且search_type设置为SEARCH_BY_SOURCE时，本字段必填，且仅支持设置为USER，从而用于筛选用户自定义上传的、可用于轮播广告的音乐。
	// 音乐来源。
	// USER 表示用户上传的音乐，SYSTEM 表示系统音乐。
	// 最大数量：1
	Sources []string `json:"sources,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.MusicScene != "" {
		values.Set("music_scene", string(r.MusicScene))
	}
	if r.SearchType != "" {
		values.Set("search_type", string(r.SearchType))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// GetResponse 获取音乐列表 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Musics 音乐数据
	Musics []Music `json:"musics,omitempty"`
}
