package video

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdSearchRequest 搜索视频 API Request
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
	// VideoIDs 视频 ID 列表。
	// 最大数量：100。
	// 注意：
	// 若视频不可用于创建广告，视频 ID 将被忽略。
	// 若未传入本筛选字段，且由于多次上传，同一视频生成了多个视频 ID，则仅会为该视频返回一个有效的视频 ID。
	VideoIDs []string `json:"video_ids,omitempty"`
	// MaterialIDs 素材 ID 列表。
	// 最大数量：20。
	MaterialIDs []string `json:"material_ids,omitempty"`
	// VideoName 视频名称。您可以输入视频名称进行模糊搜索
	VideoName string `json:"video_name,omitempty"`
	// VideoMaterialSource 视频素材来源列表。您可以通过选择视频素材来源列表来进行模糊搜索。
	// 枚举值：
	// UPLOADED_TO_TIKTOK_ADS_MANAGER：上传至 TikTok 广告管理平台的视频。
	// UPLOADED_TO_CATALOG：上传至商品库的视频。
	// CREATIVE_TOOL_SMART_VIDEO：创意工具 - 微电影。
	// CREATIVE_TOOL_QUICK_OPTIMIZATION：创意工具 - 一键优化。
	// CREATIVE_TOOL_VIDEO_TEMPLATE：创意工具 - 模版视频。
	// CREATIVE_TOOL_SMART_VIDEO_SOUNDTRACK：创意工具 - 智能配乐。
	// CREATIVE_TOOL_TIKTOK_VIDEO_EDITOR：创意工具 - 视频编辑器。
	// TIKTOK_CREATIVE_EXCHANGE：TikTok Creative Exchange。
	// CATALOG_VIDEO_TEMPLATE：商品库视频模版。
	// DYNAMIC_VIDEO_EDITOR：动态视频编辑器。
	// CREATIVE_CHALLENGE：创意挑战。
	// AUTOMATED_CREATIVE_OPTIMIZATION：程序化创意。
	// OTHER：其他。
	// QUICK_GENERATION：快速生成。
	// CREATIVE_CENTER_VIDEO_UPLOAD：创意中心 - 视频上传。
	// CREATIVE_CENTER_TIKTOK_VIDEO_EDITOR：创意中心 - TikTok 视频编辑器。
	// CREATIVE_CENTER_VIDEO_TEMPLATE：创意中心 - 视频模板。
	// DYNAMIC_SCENE：动态场景。
	// SMART_OPTIMIZATION_TOOL：智能优化工具。
	VideoMaterialSource []enum.VideoMaterialSource `json:"video_material_source,omitempty"`
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
	// List 视频数据
	List []Video `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
