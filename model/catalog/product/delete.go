package product

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DeleteRequest  删除商品 API Request
type DeleteRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedId 更新源ID
	FeedID string `json:"feed_id,omitempty"`
	// SkuIDs 对于电商商品库的商品必填。
	// 想要删除的电商商品的 SKU ID 列表。
	// 最大数量：1,000。
	// 长度限制：每个 SKU ID 最大 100 个字符，不能包含表情符号。
	// 若想获取电商商品库中商品的 SKU ID，可使用/catalog/product/get。商品的 SKU ID 将通过返回参数 sku_id 返回。
	// 注意：不支持重复的 SKU ID。
	SkuIDs []string `json:"sku_ids,omitempty"`
	// HotelIDs 对于酒店商品库的商品必填。
	// 想要删除的酒店的 ID 列表。
	// 最大数量：1,000。
	// 长度限制：每个酒店 ID 最大 100 个字符。
	// 若想获取酒店商品库中酒店的 ID，可使用/catalog/product/get。酒店的 ID 将通过返回参数 hotel_id 返回。
	// 注意：不支持重复的酒店 ID。
	HotelIDs []string `json:"hotel_ids,omitempty"`
	// FlightIDs 对于航班商品库的商品必填。
	// 想要删除的航班的 ID 列表。
	// 最大数量：1,000。
	// 长度限制：每个航班 ID 最大 150 个字符。
	// 若想获取航班商品库中航班的 ID，可使用/catalog/product/get。航班的 ID 将通过返回参数 flight_id 返回。
	// 注意：不支持重复的航班 ID。
	FlightIDs []string `json:"flight_ids,omitempty"`
	// DestinationIDs 对于目的地商品库的商品必填。
	// 想要删除的目的地的 ID 列表。
	// 最大数量：1,000。
	// 长度限制：每个目的地 ID 最大 100 个字符。
	// 若想获取目的地商品库中目的地的 ID，可使用/catalog/product/get。目的地的 ID 将通过返回参数 destination_id 返回。
	// 注意：不支持重复的目的地 ID。
	DestinationIDs []string `json:"destination_ids,omitempty"`
	// SeriesIDs 对于短剧商品库的商品必填。
	// 想要删除的短剧的 ID 列表。
	// 最大数量：1,000。
	// 长度限制：每个短剧 ID 最大 100 个字符。
	// 若想获取短剧商品库中短剧的 ID，可使用 /catalog/product/get/。短剧的 ID 将通过返回参数 series_id 返回。
	// 注意: 不支持重复的短剧 ID。
	SeriesIDs []string `json:"series_ids,omitempty"`
}

// Encode implements PostRequest interface
func (r *DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除商品 API Response
type DeleteResponse struct {
	model.BaseResponse
	Data struct {
		// FeedLogID 商品处理日志ID
		FeedLogID string `json:"feed_log_id,omitempty"`
	} `json:"data"`
}
