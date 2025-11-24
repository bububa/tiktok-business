package dmp

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CustomAudienceCreateRequest 通过文件创建受众 API Request
type CustomAudienceCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceName 自定义受众名称，限128字符内
	CustomAudienceName string `json:"custom_audience_name,omitempty"`
	// AudienceSubType 受众子类型，表明可以使用的广告类型。枚举值：NORMAL: 常规受众，可用于非覆盖和频次广告。REACH_FREQUENCY：覆盖和频次广告受众，只可用于覆盖和频次广告。默认值：NORMAL
	AudienceSubType enum.AudienceSubType `json:"audience_sub_type,omitempty"`
	// FilePaths 数据文件地址列表。
	// 为确保请求稳定性，推荐数量<50。您可使用/dmp/custom_audience/update/接口为受众增加更多数据文件。
	// 在您通过接口上传受众文件上传文件后，您会在返回中获得file_path
	FilePaths []string `json:"file_paths,omitempty"`
	// CaculateType 加密类型。加密类型需要和实际文件数据类型保持一致，否则上传将会失败，后续受众人群创建也会失败。枚举值详见枚举值-加密类型
	CaculateType enum.CalculateType `json:"caculate_type,omitempty"`
	// RetentionInDays 保留该受众的天数。
	// 取值范围：[1, 365]。
	// 注意：
	// 若设置了本字段，受众将在距离创建时间的指定保留天数后过期，对该受众的任何操作均不会重置过期日期。
	// 若未设置本字段，受众将在连续12个月未应用于任何活跃广告组且未进行任何修改后过期，将该受众应用于活跃广告组或修改该广告组会重置过期日期。欲了解会重置过期日期的操作，请查看帮助中心文档受众过期政策。
	RetentionInDays int `json:"retention_in_days,omitempty"`
}

// Encode implements PostRequest
func (r *CustomAudienceCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CustomAudienceCreateResponse 通过文件创建受众 API Response
type CustomAudienceCreateResponse struct {
	model.BaseResponse
	Data struct {
		// CustomAudienceID 自定义受众ID
		CustomAudienceID string `json:"custom_audience_id,omitempty"`
	} `json:"data"`
}
