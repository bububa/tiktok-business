package video

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// LogRequest 获取商品库视频的操作结果 API Request
type LogRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您有权访问的商务中心列表，可使用/bc/get/
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库 ID。
	// 若想获取某个商务中心中的电商商品库列表，可使用 /catalog/get/，筛选对应的 catalog_type 为 ECOM 的 catalog_id值。
	// 您需对商品库有商品库管理（目录管理）权限。若想获取您有商品库管理权限的商品库， 可使用/bc/asset/get/。将asset_type 设置为 CATALOG，筛选对应的 catalog_role 为 ADMIN 的 asset_id 值。
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedLogID 商品库视频处理日志的 ID。
	// 若想上传商品库视频并获取日志 ID，可使用/catalog/video/file/。
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

// LogResponse 获取商品库视频的操作结果 API Response
type LogResponse struct {
	model.BaseResponse
	Data *LogResult `json:"data,omitempty"`
}

type LogResult struct {
	// VideoFeedLog 商品库视频处理结果
	VideoFeedLog *Log `json:"video_feed_log,omitempty"`
	// FeedLogData 详细处理信息
	FeedLogData *LogData `json:"feed_log_data,omitempty"`
}

// Log 商品库视频处理结果
type Log struct {
	// CatalogID 商品库 ID
	CatalogID string `json:"catalog_id,omitempty"`
	// FeedID 更新源 ID
	FeedID string `json:"feed_id,omitempty"`
	// AddCount 新增的视频总数
	AddCount int `json:"add_count,omitempty"`
	// UpdateCount 修改的视频总数。
	// 注意：本字段值始终为 0。
	UpdateCount int `json:"update_count,omitempty"`
	// DeleteCount 删除的商品总数
	DeleteCount int `json:"delete_count,omitempty"`
	// ErrorCount 出现上传错误的商品数量
	ErrorCount int `json:"error_count,omitempty"`
	// WarnCount 出现上传提示的商品数量
	WarnCount int `json:"warn_count,omitempty"`
	// ProcessStatus 处理状态。
	// 枚举值：PROCESSING，SUCCESS，FAILED，WAITING
	ProcessStatus string `json:"process_status,omitempty"`
	// StartTime 视频处理任务的开始时间，格式为YYYY-MM-DD HH:MM:SS(UTC+0)。
	// 示例：2024-01-01 00:00:00
	StartTime model.DateTime `json:"start_time,omitzero"`
	// EndTime 视频处理任务的结束时间，格式为 YYYY-MM-DD HH:MM:SS (UTC+0)。
	// 示例：2024-01-01 00:00:00
	EndTime model.DateTime `json:"end_time,omitzero"`
}

// LogData 详细处理信息
type LogData struct {
	// DownloadPath 文件下载地址，包含所有错误/警告信息，key是语言，value是文件地址
	DownloadPath map[string]string `json:"download_path,omitempty"`
	// ErrorAffectedVideos 错误信息
	ErrorAffectedVideos []AffectedVideo `json:"error_affected_videos,omitempty"`
	// WarnAffectedVideos 警告信息
	WarnAffectedVideos []AffectedVideo `json:"warn_affected_videos,omitempty"`
}

type AffectedVideo struct {
	// AffectedVideoCount 受影响商品数
	AffectedVideoCount int `json:"affected_video_count,omitempty"`
	// AffectedVideoItemList 受影响商品列表，只展示前五条
	AffectedVideoItemList []AffectedVideoItem `json:"affected_video_item_list,omitempty"`
	// Field 问题字段
	Field string `json:"field,omitempty"`
	// Issue 问题信息
	Issue string `json:"issue,omitempty"`
	// Suggestion 解决建议
	Suggestion string `json:"suggestion,omitempty"`
}

type AffectedVideoItem struct {
	// Index 视频索引，代表该视频在所上传视频列表中的位置。
	// 例如，索引值为 2 代表该商品为通过 /catalog/video/file/上传的视频列表中的第 1 个商品，位于 CSV 文件中的第二行。
	Index int `json:"index,omitempty"`
	// VideoName 视频名称
	VideoName string `json:"video_name,omitempty"`
	// VideoLink 视频文件的 URL。
	// 示例： https://www.tiktok.com/tiktok_t_shirt.mp4
	VideoLink string `json:"video_link,omitempty"`
	// SkuIDList 要与该视频（video_link）绑定的电商商品库商品的 SKU ID 列表。
	// 示例： "SKU_ID_1, SKU_ID_2"。
	SkuIDList string `json:"sku_id_list,omitempty"`
	// Description 视频的简短描述
	Description string `json:"description,omitempty"`
	// Category 视频的商品类别。
	// 示例：Apparel & Accessories> Clothing> Shirts
	Category string `json:"category,omitempty"`
}
