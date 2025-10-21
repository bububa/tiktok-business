package video

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取商品库中上传的商品库视频 API Request
type GetRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您有权访问的商务中心列表，可使用/bc/get/
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库 ID。
	// 若想获取某个商务中心中的电商商品库列表，可使用 /catalog/get/，筛选对应的 catalog_type 为 ECOM 的 catalog_id值。
	// 您需对商品库有商品库管理（目录管理）权限。若想获取您有商品库管理权限的商品库， 可使用/bc/asset/get/。将asset_type 设置为 CATALOG，筛选对应的 catalog_role 为 ADMIN 的 asset_id 值。
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogVideoIDs 想要筛选的商品库视频的 ID。
	// 最大数量：50。
	// 注意：若您为catalog_video_ids传入了值，将忽略分页参数 page 和 page_size。
	CatalogVideoIDs []string `json:"catalog_video_ids,omitempty"`
	// Page 当前页数。
	// 取值范围：≥1。默认值：1。
	// 注意：若您为catalog_video_ids传入了值，将忽略分页参数 page 和 page_size
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-50。默认值：10。
	// 注意：若您为catalog_video_ids传入了值，将忽略分页参数 page 和 page_size
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if len(r.CatalogVideoIDs) > 0 {
		values.Set("catalog_video_ids", string(util.JSONMarshal(r.CatalogVideoIDs)))
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

// GetResponse 获取商品库中上传的商品库视频 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// Videos 商品库视频列表
	Videos []Video `json:"videos,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
