package dmp

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// CustomAudienceUpdateRequest 修改受众 API Request
type CustomAudienceUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceID 自定义受众ID
	CustomAudienceID string `json:"custom_audience_id,omitempty"`
	// CustomAudienceName 自定义受众名称，限128字符内
	CustomAudienceName string `json:"custom_audience_name,omitempty"`
	// AudienceSubType 受众子类型，表明可以使用的广告类型。枚举值：NORMAL: 常规受众，可用于非覆盖和频次广告。REACH_FREQUENCY：覆盖和频次广告受众，只可用于覆盖和频次广告。默认值：NORMAL
	AudienceSubType enum.AudienceSubType `json:"audience_sub_type,omitempty"`
	// FilePaths 数据文件地址列表。
	// 为确保请求稳定性，推荐数量<50。您可使用/dmp/custom_audience/update/接口为受众增加更多数据文件。
	// 在您通过接口上传受众文件上传文件后，您会在返回中获得file_path
	FilePaths []string `json:"file_paths,omitempty"`
	// Action 数据文件操作类型。当传入file_paths时，该字段有效。枚举值:
	// APPEND：上传文件扩充原始客户文件受众数据。仅增加原始文件中没有的数据。
	// REMOVE：上传文件删除原始客户文件受众数据。仅删除与原始文件重复的数据。
	// REPLACE：上传文件替换原始客户文件受众数据。替换全部原始受众数据。
	// 默认值: REPLACE。
	// 注意：新增（APPEND），移除（REMOVE）或替换（REPLACE）操作成功的前提是操作完成后受众大小不低于1000。否则新增、移除或替换操作将会失败。
	Action string `json:"action,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAudienceUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
