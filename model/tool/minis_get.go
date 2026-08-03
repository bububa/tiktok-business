package tool

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// MinisGetRequest 获取广告账号内 TikTok Minis API Request。
type MinisGetRequest struct {
	// AdvertiserID 广告主 ID。
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Page 当前页数。默认值：1。
	Page int `json:"page,omitempty"`
	// PageSize 每页数量。取值范围：1-50，默认值：10。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest.
func (r *MinisGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
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

// MinisGetResponse 获取广告账号内 TikTok Minis API Response。
type MinisGetResponse struct {
	model.BaseResponse
	Data *MinisGetResult `json:"data,omitempty"`
}

// MinisGetResult 广告账号内 TikTok Minis 列表。
type MinisGetResult struct {
	// List TikTok Minis 列表。
	List []Minis `json:"list,omitempty"`
	// PageInfo 分页信息。
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Minis TikTok Minis 信息。
type Minis struct {
	// MinisID TikTok Minis ID。
	MinisID string `json:"minis_id,omitempty"`
	// MinisName TikTok Minis 名称。
	MinisName string `json:"minis_name,omitempty"`
	// MinisIconURL TikTok Minis 图标 URL。
	MinisIconURL string `json:"minis_icon_url,omitempty"`
	// MinisStatus TikTok Minis 状态。
	MinisStatus enum.MinisStatus `json:"minis_status,omitempty"`
	// MinisType TikTok Minis 类型。
	MinisType enum.MinisType `json:"minis_type,omitempty"`
	// RegionCodes TikTok Minis 可用的地区代码列表。
	RegionCodes []string `json:"region_codes,omitempty"`
}
