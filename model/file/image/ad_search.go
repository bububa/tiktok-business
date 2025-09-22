package image

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdSearchRequest 搜索图片 API Request
type AdSearchRequest struct {
	// AdvertiserID 广告账号
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Filtering 筛选条件。本字段是筛选对象的数组
	Filtering *AdSearchFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围：≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 20。
	// 取值范围：1-1,000。
	PageSize int `json:"page_size,omitempty"`
}

// AdSearchFilter 筛选条件。本字段是筛选对象的数组
type AdSearchFilter struct {
	// ImageIDs 图片ID列表。最多可包含100个ID。
	ImageIDs []string `json:"image_ids,omitempty"`
	// MaterialIDs 素材 ID 列表。
	// 最大数量：1000。
	MaterialIDs []string `json:"material_ids,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
	// Ratio 图片宽高比，例如 [1.7, 2.5]，输入1.7则搜索满足宽高比介于1.65-1.75之间的图片，即精度上下浮动0.05
	Ratio []float64 `json:"ratio,omitempty"`
	// Displayable 期望搜索范围。枚举值：
	// False（默认值）：在所有素材中搜索。
	// True：在平台展示的素材中搜索。
	Displayable bool `json:"displayable,omitempty"`
}

// Encode implements GetRequest interface
func (r AdSearchRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
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

// AdSearchResponse 搜索视频 API Response
type AdSearchResponse struct {
	model.BaseResponse
	Data *AdSearchResult `json:"data"`
}

type AdSearchResult struct {
	// List 图片数据
	List []Image `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
