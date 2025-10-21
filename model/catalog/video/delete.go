package video

import (
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest 删除上传的商品库视频 API Request
type DeleteRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您有权访问的商务中心列表，可使用/bc/get/
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库 ID。
	// 若想获取某个商务中心中的电商商品库列表，可使用 /catalog/get/，筛选对应的 catalog_type 为 ECOM 的 catalog_id值。
	// 您需对商品库有商品库管理（目录管理）权限。若想获取您有商品库管理权限的商品库， 可使用/bc/asset/get/。将asset_type 设置为 CATALOG，筛选对应的 catalog_role 为 ADMIN 的 asset_id 值。
	CatalogID string `json:"catalog_id,omitempty"`
	// CatalogVideoIDs 要删除的商品库视频的 ID 列表。
	// 最大数量：1。
	// 若想获取商品库中的商品库视频，可使用 /catalog/video/get/。
	// 注意：商品库视频同步至广告账号的素材库后，即使从相关联的商品库中删除这些视频，广告账号的素材库中依然保留这些视频。
	CatalogVideoIDs []string `json:"catalog_video_ids,omitempty"`
}

// Encode implements PostRequest interface
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
