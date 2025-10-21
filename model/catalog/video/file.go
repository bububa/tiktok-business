package video

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// FileRequest 通过文件 URL 上传商品库视频 API Request
type FileRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您有权访问的商务中心列表，可使用/bc/get/
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库 ID。
	// 若想获取某个商务中心中的电商商品库列表，可使用 /catalog/get/，筛选对应的 catalog_type 为 ECOM 的 catalog_id值。
	// 您需对商品库有商品库管理（目录管理）权限。若想获取您有商品库管理权限的商品库， 可使用/bc/asset/get/。将asset_type 设置为 CATALOG，筛选对应的 catalog_role 为 ADMIN 的 asset_id 值。
	CatalogID string `json:"catalog_id,omitempty"`
	// FileURL CSV 文件的下载地址。
	// 商品库视频上传的 CSV 模板为catalog_video_template。
	// 若想查看每个商品库视频参数的具体描述，请查看商品库视频参数。
	FileURL string `json:"file_url,omitempty"`
	// AdvertiserIDs 视频同步的目标广告账号 ID 列表。
	// 同步完成后，视频将存在于广告账号的素材库中。若您不想将商品库视频同步至广告账号中，则无需传入本字段。
	// 最大数量：100。
	// 广告账号和商品库（catalog_id）需在同一商务中心（bc_id）中，且您需对广告账号有管理员或操作员权限。不合法的广告账号将被忽略。若想获取您有管理员或操作员权限的 广告账号，可使用/bc/asset/get/。将 asset_type 设置为 ADVERTISER，并筛选 advertiser_role为 ADMIN 或 OPERATOR 的 asset_id 值。
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
}

// Encode implements PostRequest interface
func (r *FileRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// FileResponse 通过文件 URL 上传商品库视频 API Response
type FileResponse struct {
	model.BaseResponse
	Data struct {
		// FeedLogID 商品库视频处理日志的 ID。
		// 若想查询商品库视频上传的处理状态和最终结果，可使用/catalog/video/log/。
		FeedLogID string `json:"feed_log_id,omitempty"`
	} `json:"data"`
}
