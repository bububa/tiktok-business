package product

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// FileRequest 通过文件 URL 上传商品 API Request
type FileRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedId 更新源ID
	FeedID string `json:"feed_id,omitempty"`
	// FileURL 商品上传文件（CSV 格式）的下载地址。建议商品数量不要超过20,000,000个。
	// 注意：
	// 目前 API 只允许上传 CSV 格式的文件。不同类型的商品库的 CSV 模版如下：
	// 电商商品库模板
	// 酒店商品库模板
	// 航班商品库模板
	// 目的地商品库模板
	// 若 CSV 文件包含别名字段，系统将自动将别名字段映射到对应的已有字段。若想了解别名字段和对应的已有字段，可查看航班商品库的别名字段。
	// 娱乐商品库模板
	// 短剧商品库模版
	// 若想获取每个商品参数的详细描述，请查看对象数组products参数。
	// 若想了解 JSON 格式中的字段及其枚举值在 CSV 模版中的对应字段和枚举值，请查看 JSON 结构与 CSV 模版中的商品参数对比 。
	// 商品图片（image_link）规格必须大于等于500x500像素，否则商品将无法通过审核。详细信息请参考商品图片要求。
	FileURL string `json:"file_url,omitempty"`
	// UpdateMode 更新模式。该模式仅单次有效，只应用于当前请求。
	// 枚举值：
	// OVERWRITE：在本模式下，您上传的商品文件中的商品将替换商品库中当前的所有商品。
	// INCREMENTAL：在本模式下，您上传的商品文件中的商品将增加到商品库中。请注意若您上传了商品库中已有的商品（SKU ID重复），则该商品的数据将由您上传文件中的该商品对应数据所覆盖。
	// 默认值：INCREMENTAL。
	// 注意：若您传入了 feed_id，则您通过本字段指定的更新模式不会影响指定的更新源的更新模式。
	UpdateMode enum.CatalogFeedUpdateMode `json:"update_mode,omitempty"`
}

// Encode implements PostRequest interface
func (r *FileRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// FileResponse 通过文件 URL 上传商品 API Response
type FileResponse struct {
	model.BaseResponse
	Data struct {
		// FeedLogID 商品处理日志ID
		FeedLogID string `json:"feed_log_id,omitempty"`
	} `json:"data"`
}
