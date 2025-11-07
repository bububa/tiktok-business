package gmvmax

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/file/video"
	"github.com/bububa/tiktok-business/util"
)

// VideoGetRequest 获取商品 GMV Max 推广系列可用的帖子 API Request
type VideoGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StoreID TikTok Shop ID。
	// 若想获取可用于 GMV Max 推广系列的 TikTok Shop，可使用 /gmv_max/store/list/ 并确认返回的 is_gmv_max_available 为 true
	StoreID string `json:"store_id,omitempty"`
	// StoreAuthorizedBcID 有权限访问该 TikTok Shop（store_id）的商务中心的 ID
	StoreAuthorizedBcID string `json:"store_authorized_bc_id,omitempty"`
	// SpuIDList 用于筛选结果的商品 SPU（标准化产品单元）ID 列表。
	// 若您想在商品 GMV Max 推广系列中使用已授权的作品推广此 TikTok Shop 的所有商品，无需传入 spu_id_list 和 custom_posts_eligible 本字段。
	// 若您想在商品 GMV Max 推广系列中使用已授权的作品推广此 TikTok Shop 的特定商品，需通过 spu_id_list 本字段指定这些商品，无需传入 custom_posts_eligible。
	// 若您想在商品 GMV Max 推广系列中使用自定义的作品推广此 TikTok Shop 的某个商品，需通过 spu_id_list 字段指定该商品，并将 custom_posts_eligible 设置为 true。
	// 最大数量：
	// 当 custom_posts_eligible 为 false 或未传入时：50。
	// 当 custom_posts_eligible 为 true 时：1。
	// 若想获取某个 TikTok Shop 中商品的 SPU ID 列表，可使用/store/product/get/。将 ad_creation_eligible 设置为 GMV_MAX 并从返回中挑选 status 为 AVAILABLE 且 gmv_max_ads_status 为 UNOCCUPIED 的 item_group_id 值。
	SpuIDList []string `json:"spu_id_list,omitempty"`
	// CustomPostEligible 是否请求用于自定义的作品中推广某个商品的视频。
	// 若您将本字段设置为 true，您需同时通过 spu_id_list 指定一个商品。
	// 支持的值：true，false。
	// 默认值：false
	CustomPostEligible bool `json:"custom_post_eligible,omitempty"`
	// SortField 仅当 custom_posts_eligible 为 false 或未传入时生效。
	// 对用于自定义的作品的视频的排序字段。
	// 枚举值:
	// GMV：商品交易总额。
	// POST_TIME：发布时间。
	// VIDEO_VIEWS：视频播放量。
	// VIDEO_LIKES：视频点赞数。
	// CLICK_THROUGH_RATE：点击率。
	// PRODUCT_CLICKS：商品点击次数。
	// 默认值：GMV。
	SortField string `json:"sort_field,omitempty"`
	// SortType 对用于已授权的作品或自定义的作品的视频的排序顺序。
	// 枚举值:
	// ASC：升序。
	// DESC：降序。
	// 默认值：DESC。
	// 若 custom_posts_eligible 为 false 或未传入，您可通过同时指定 sort_field 和 sort_type 对用于已授权的作品的视频排序。
	// 若 custom_posts_eligible 为 true，您无需指定 sort_field。在此情况下，默认对用于自定义的作品的视频按发布时间降序排序，最先返回最近发布的作品。若您希望按升序排序（最先返回最早发布的作品），可手动将 sort_type 设置为 ASC
	SortType enum.OrderType `json:"sort_type,omitempty"`
	// Keyword 用于筛选帖子（视频）的搜索关键词。
	// 您可指定帖子 ID 或文案。
	// 若您想根据文案（text）搜索帖子，可提供一个文本字符串。这种搜索类型支持多语言搜索和模糊匹配。
	// 示例：summer。
	// 若您想根据帖子 ID（item_id）搜索帖子，需提供最小长度为 19 个字符的数字字符串。这种搜索类型支持精确匹配。
	// 示例：1234567891234567891
	Keyword string `json:"keyword,omitempty"`
	// NeedAuthCodeVideo 是否在结果中包含与 AUTH_CODE 类型的认证身份绑定的帖子。
	// 支持的值： true，false。
	// 默认值：false。
	// 注意：若未传入 need_auth_code_video 或设置为 false 且未传入 identity_list，返回的结果将为空
	NeedAuthCodeVideo bool `json:"need_auth_code_video,omitempty"`
	// IdentityList 用于筛选结果的 TT_USER、BC_AUTH_TT 或 TTS_TT 认证身份列表。
	// 若想获取可用于商品 GMV Max 推广系列的认证身份列表，可使用 /gmv_max/identity/get/并确认返回的 product_gmv_max_available 为true。
	// 最大数量：20。
	// 注意：若未传入 need_auth_code_video 或设置为 false 且未传入 identity_list，返回的结果将为空
	IdentityList []Identity `json:"identity_list,omitempty"`
	// Page 当前页数。
	// 取值范围： ≥ 1。
	// 默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小。
	// 取值范围：1-50。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *VideoGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("store_id", r.StoreID)
	values.Set("store_authorized_bc_id", r.StoreAuthorizedBcID)
	if len(r.SpuIDList) > 0 {
		values.Set("spu_id_list", string(util.JSONMarshal(r.SpuIDList)))
	}
	if r.CustomPostEligible {
		values.Set("custom_post_eligible", "true")
	}
	if r.SortField != "" {
		values.Set("sort_field", r.SortField)
	}
	if r.SortType != "" {
		values.Set("sort_type", string(r.SortType))
	}
	if r.Keyword != "" {
		values.Set("keyword", r.Keyword)
	}
	if r.NeedAuthCodeVideo {
		values.Set("need_auth_code_video", "true")
	}
	if len(r.IdentityList) > 0 {
		values.Set("identity_list", string(util.JSONMarshal(r.IdentityList)))
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

// VideoGetResponse 获取商品 GMV Max 推广系列可用的帖子 API Response
type VideoGetResponse struct {
	model.BaseResponse
	Data *VideoGetResult `json:"data,omitempty"`
}

type VideoGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// ItemList 使用该 TikTok Shop 的商品 GMV Max 推广系列中可用的 TikTok 视频（帖子）列表
	ItemList []TikTokItem `json:"item_list,omitempty"`
}

// TikTokItem 与该 GMV Max 推广系列绑定的已授权 TikTok 作品（即 TikTok 帖子）或自定义的作品。
type TikTokItem struct {
	// ItemID TikTok 帖子 ID
	ItemID string `json:"item_id,omitempty"`
	// Text TikTok 帖子描述
	Text string `json:"text,omitempty"`
	// SpuIDList 与该 TikTok 帖子绑定的商品 SPU ID 列表
	SpuIDList []string `json:"spu_id_list,omitempty"`
	// CanChangeAnchor 是否支持更改视频中的商品锚点链接。
	// 支持的值：
	// true：您可在此视频中新增商品锚点链接创建自定义的作品。
	// false：不支持使用此视频创建自定义的作品。
	// 注意：若 can_change_anchor 为 false，则不支持将本字段传入 /campaign/gmv_max/create/ 或 /campaign/gmv_max/update/ 中的 custom_anchor_video_list 或 item_list 创建自定义的作品
	CanChangeAnchor bool `json:"can_change_anchor,omitempty"`
	// IdentityInfo 与该 TikTok 帖子绑定的认证身份的有关信息
	IdentityInfo *Identity `json:"identity_info,omitempty"`
	// VideoInfo 帖子中视频的详情
	VideoInfo *video.Video `json:"video_info,omitempty"`
}
