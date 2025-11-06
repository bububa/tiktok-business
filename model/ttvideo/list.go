package ttvideo

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取 Spark Ads 帖子列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// ItemTypes 要筛选的 Spark Ads 帖子类型。
	// 枚举值：
	// VIDEO：视频帖子。
	// CAROUSEL：照片帖子。
	// 默认值：["VIDEO", "CAROUSEL"]。
	ItemTypes []string `json:"item_types,omitempty"`
	// Keyword 可以作为搜索TikTok帖子的关键词，可以为文本或TikTok帖子ID。
	// 如果您想根据文本搜索TikTok帖子，提供最大长度为500个字符的文本字符串。这种搜索类型支持多语言搜索和模糊匹配。
	// 示例："summer"。
	// 如果你想根据item_id搜索TikTok帖子，提供最小长度为19个字符的数字字符串。这种搜索类型支持精确匹配。
	// 示例："1234567891234567891"
	Keyword string `json:"keyword,omitempty"`
	// Page 当前页数。默认值: 1，取值范围: ≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。默认值: 20，取值范围: 1-50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.ItemTypes) > 0 {
		values.Set("item_types", string(util.JSONMarshal(r.ItemTypes)))
	}
	if r.Keyword != "" {
		values.Set("keyword", r.Keyword)
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

// ListResponse 获取 Spark Ads 帖子列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List Spark Ads 帖子列表信息
	List []Item `json:"list,omitempty"`
}
