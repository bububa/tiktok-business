package adgroup

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// QuotaRequest 获取投放中广告组的动态配额 API Request
type QuotaRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implemnts GetRequest interface
func (r *QuotaRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// QuotaResponse 获取投放中广告组的动态配额 API Response
type QuotaResponse struct {
	model.BaseResponse
	Data *Quota `json:"data,omitempty"`
}

// Quota 广告组的动态配额
// 注意：
// 若广告主（advertiser_id）的投放中广告组配额固定，返回中将显示total_adgroup_quota 为 5000，used_adgroup_quota为 0。
// 若广告主（advertiser_id）处在惩罚状态，则不允许创建投放中广告组，此时 total_adgroup_quota和used_adgroup_quota的返回值将为0。您可通过/advertiser/info/返回的status字段查看广告账户是否处在惩罚状态。
type Quota struct {
	// TotalAdgroupQuota 广告主（advertiser_id）可拥有的投放中广告组最大数。
	TotalAdgroupQuota int `json:"total_adgroup_quota,omitempty"`
	// UsedAdgroupQuota 广告主（advertiser_id）目前拥有的投放中广告组数量。
	// 注意：若所指定的广告主已用尽投放中广告组的动态配额（total_adgroup_quota = used_adgroup_quota），则该广告主可以继续创建额外广告组，但新建的额外广告组在数秒后将自动变为暂停状态。此外，广告主还将无法再使用 /adgroup/status/update/将处于暂停状态的广告组更新为投放中状态。
	UsedAdgroupQuota int `json:"used_adgroup_quota,omitempty"`
}
