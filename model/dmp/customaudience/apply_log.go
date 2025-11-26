package customaudience

import (
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ApplyLogRequest 获取受众应用记录 API Request
type ApplyLogRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CustomAudienceIDs 自定义受众ID列表。
	// 注意
	// 广告主必须是custom_audience_ids的拥有者，否则将会报错。您可以使用/dmp/custom_audience/list/和 /dmp/custom_audience/get/返回的is_creator字段查看该广告主是否为该受众的拥有者。
	CustomAudienceIDs []string `json:"custom_audience_ids,omitempty"`
	// Timezone 返回值的时区。获取枚举值，请参考附录-时区信息。
	// 若不传，则使用UTC时区。
	Timezone string `json:"timezone,omitempty"`
	// Page 当前页数，默认值: 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *ApplyLogRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("custom_audience_ids", string(util.JSONMarshal(r.CustomAudienceIDs)))
	if r.Timezone != "" {
		values.Set("timezone", r.Timezone)
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

// ApplyLogResponse 获取受众应用记录 API Response
type ApplyLogResponse struct {
	model.BaseResponse
	Data *ApplyLogResult `json:"data,omitempty"`
}

type ApplyLogResult struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 应用数据的数组
	List []ApplyLog `json:"list,omitempty"`
}

type ApplyLog struct {
	// AudienceID 受众ID
	AudienceID string `json:"audience_id,omitempty"`
	// AdgroupID 广告组ID
	AdgroupID string `json:"adgroup_id,omitempty"`
	// AdgroupName 广告组名称
	AdgroupName string `json:"adgroup_name,omitempty"`
	// UsageMode 该受众在广告组中的使用方式。枚举值：Include，Exclude
	UsageMode string `json:"usage_mode,omitempty"`
	// Editor 操作者
	Editor string `json:"editor,omitempty"`
	// ActionTimestamp 操作发生的时间。格式："YYYY-MM-DD HH:MM:SS"
	ActionTimestamp model.DateTime `json:"action_timestamp,omitzero"`
}
