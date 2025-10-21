package product

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// LogRequest 获取商品操作结果 API Request
type LogRequest struct {
	// BcID 商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedLogID 商品处理日志 ID
	FeedLogID string `json:"feed_log_id,omitempty"`
	// Language 支持的语言。
	// 枚举值：
	// en：英文
	// zh-CN：中文
	// ja-JP：日文
	// de：德语
	// es：西班牙语
	// fr：法语
	// hi-IN：印地语（印度）
	// id：印尼语
	// it：意大利语
	// ko：韩语
	// ms：马来语
	// ru：俄语
	// th：泰语
	// tr：土耳其语
	// vi：越南语
	//
	// 默认值：en。
	Language string `json:"language,omitempty"`
}

// Encode implements GetRequest interface
func (r *LogRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	values.Set("feed_log_id", r.FeedLogID)
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// LogResponse 获取商品操作结果 API Response
type LogResponse struct {
	model.BaseResponse
	Data struct {
		// ProductFeedLog 商品处理结果
		ProductFeedLog *Log `json:"product_feed_log,omitempty"`
	} `json:"data"`
}

// Log 商品处理结果
type Log struct {
	// CatalogID 商品库ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedID 更新源ID
	FeedID string `json:"feed_id,omitempty"`
	// AddCount 新增的商品总数
	AddCount int `json:"add_count,omitempty"`
	// UpdateCount 修改的商品总数
	UpdateCount int `json:"update_count,omitempty"`
	// DeleteCount 删除的商品总数
	DeleteCount int `json:"delete_count,omitempty"`
	// ErrorCount 上传失败的商品数量
	ErrorCount int `json:"error_count,omitempty"`
	// WarnCount 被警告的商品数量
	WarnCount int `json:"warn_count,omitempty"`
	// ProcessStatus 处理状态。
	// 枚举值：
	// WAITING：等待中，有其他任务待处理完成。
	// PROCESSING：处理中。
	// SUCCESS：处理成功。
	// FAILED：处理失败。
	ProcessStatus string `json:"process_status,omitempty"`
	// UpdateMode 更新模式。
	// 枚举值：
	// INCREASE：增量。
	// REPLACE：全量。
	UpdateMode enum.ProductUpdateMode `json:"update_mode,omitempty"`
	// StartTime 开始时间
	StartTime model.DateTime `json:"start_time,omitzero"`
	// EndTime 结束时间
	EndTime model.DateTime `json:"end_time,omitzero"`
	// FeedLogData 详细处理信息
	FeedLogData *LogData `json:"feed_log_data,omitempty"`
}

// LogData 详细处理信息
type LogData struct {
	// DownloadPath 文件下载地址，包含所有错误/警告信息，key是语言，value是文件地址
	DownloadPath map[string]string `json:"download_path,omitempty"`
	// ErrorAffectedProducts 错误信息
	ErrorAffectedProducts []AffectedProduct `json:"error_affected_products,omitempty"`
	// WarnAffectedProducts 警告信息
	WarnAffectedProducts []AffectedProduct `json:"warn_affected_products,omitempty"`
}

type AffectedProduct struct {
	// AffectedProductCount 受影响商品数
	AffectedProductCount int `json:"affected_product_count,omitempty"`
	// AffectedProductItemList 受影响商品列表，只展示前五条
	AffectedProductItemList []AffectedProductItem `json:"affected_product_item_list,omitempty"`
	// Field 问题字段
	Field string `json:"field,omitempty"`
	// Issue 问题信息
	Issue string `json:"issue,omitempty"`
	// Suggestion 解决建议
	Suggestion string `json:"suggestion,omitempty"`
}

// AffectedProductItem 受影响商品
type AffectedProductItem struct {
	// Index 商品索引，代表该商品在所上传商品列表中的位置。
	// 例如，索引值为 2 代表该商品为通过/catalog/product/file/或/catalog/product/upload/上传的商品列表中的第 1 个商品
	Index int `json:"index,omitempty"`
	// Title 仅为电商商品库、娱乐商品库或短剧商品库中的商品返回此字段。
	// 电商商品库或娱乐商品库中商品的标题或短剧商品库中短剧的名称。
	Title string `json:"title,omitempty"`
	// SkuID 仅为电商商品库中的商品返回此字段。
	// 商品的唯一 ID，例如 SKU。
	SkuID string `json:"sku_id,omitempty"`
	// HotelID 仅为酒店商品库中的商品返回此字段。
	// 酒店的唯一ID
	HotelID string `json:"hotel_id,omitempty"`
	// Name 仅为酒店商品库或目的地商品库中的商品返回此字段。
	// 酒店或目的地的名称。
	Name string `json:"name,omitempty"`
	// FlightID 仅为航班商品库中的商品返回此字段。
	// 航班的唯一ID。
	FlightID string `json:"flight_id,omitempty"`
	// AirlineCompany 仅为航班商品库中的商品返回此字段。
	// 运营该航班的航空公司的名称。
	AirlineCompany enum.AirlineCompany `json:"airline_company,omitempty"`
	// MediaTitleID 仅为娱乐商品库中的商品返回此字段。
	// 媒体标题项的唯一ID。
	MediaTitleID string `json:"media_title_id,omitempty"`
	// SeriesID 仅为短剧商品库中的商品返回此字段。
	// 短剧的自定义唯一 ID。
	SeriesID string `json:"series_id,omitempty"`
	// ProductURL 商品链接
	ProductURL string `json:"product_url,omitempty"`
	// Description 商品描述
	Description string `json:"description,omitempty"`
	// Value 字段值
	Value string `json:"value,omitempty"`
}
