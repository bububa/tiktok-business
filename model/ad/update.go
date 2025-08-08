package ad

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UpdateRequest 更新广告 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdgroupID 广告组 ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// PatchUpdate 是否使用增量更新模式。
	// 本字段设置为 true 时，支持增量更新以下字段：
	// ad_name
	// identity_id
	// identity_type
	// identity_authorized_bc_id
	// video_id
	// image_ids
	// music_id
	// tiktok_item_id
	// ad_text
	//
	// 例如，若想增量更新广告名称，可仅设置以下参数：
	// patch_update 设置为 true 。
	// ad_name 设置为新的广告名称。
	// 传入其他必填参数，包括 Header 参数， advertiser_id 和 ad_id。
	//
	// 重要提示：若您想使用本字段使用增量更新模式，推荐您进行充分测试，以免对您当前的广告更新工作流程造成不必要的影响。
	PatchUpdate bool `json:"patch_update,omitempty"`
	// Creatives 广告创意。
	// 最大数量：20。
	// 单次请求中可在一个广告组中创建最多 20 个广告。若想在同一广告组中创建额外的广告，需多次调用/ad/create/。
	Creatives []Creative `json:"creatives,omitempty"`
}

// Encode implements PostRequest interface
func (r *UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新广告 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *UpdateResult `json:"data,omitempty"`
}

type UpdateResult struct {
	// AdIDs 广告ID
	AdIDs []string `json:"ad_ids,omitempty"`
	// Creatives 广告列表。返回字段根据请求参数中 fields 字段产生，默认返回所有字段。
	Creatives []Ad `json:"creatives,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
