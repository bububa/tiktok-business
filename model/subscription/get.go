package subscription

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest  获取开发者应用的订阅信息 API Request
type GetRequest struct {
	// AppID 开发者应用ID。您可以导航至应用管理 > App Detail > 基本信息，获取应用ID。
	AppID string `json:"app_id,omitempty"`
	// Secret 开发者应用密钥。您可以导航至应用管理 > App Detail > 基本信息，获取密钥。
	Secret string `json:"secret,omitempty"`
	// SubscribeEntity 订阅对象。
	// 枚举值:
	// AD_ACCOUNT_SUSPENSION：广告账户的暂时停用状态。
	// LEAD：线索。
	// AD_GROUP：广告组的审核状态。
	// AD：广告的审核状态。
	// TCM_SPARK_ADS：上传至某个 TCM 工作流程 2.0 订单的视频的 Spark Ads 授权状态。
	// CREATIVE_FATIGUE：单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态。
	SubscribeEntity enum.SubscribeEntity `json:"subscribe_entity,omitempty"`
	// Page 当前页数。
	// 默认值：1。
	// 取值范围：≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 10。
	// 取值范围: 1-1,000。
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("app_id", r.AppID)
	values.Set("secret", r.Secret)
	if r.SubscribeEntity != "" {
		values.Set("subscribe_entity", string(r.SubscribeEntity))
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

// GetResponse 获取开发者应用的订阅信息 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// Subscriptions 订阅信息
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
