package order

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CreateRequest 创建 TTCM 订单 API Request
type CreateRequest struct {
	// TCMAccountID TikTok Creator Marketplace 账户ID
	TCMAccountID string `json:"tcm_account_id,omitempty"`
	// BrandName 品牌名称。
	// 长度限制：60个字符。
	BrandName string `json:"brand_name,omitempty"`
	// Description 订单描述。
	// 长度限制：60个字符。
	Description string `json:"description,omitempty"`
	// HandleNames 创作者的用户名列表。
	// 最大数量：100。
	HandleNames []string `json:"handle_names,omitempty"`
	// CampaignCode 推广系列代码。若传入推广系列代码，则订单会在与该推广系列代码绑定的推广系列中生成。
	// 注意：
	// campaign_code 必须属于传入的TikTok Creator Marketplace 账户（tcm_account_id） 和品牌（brand_name）。
	// 若您已通过 deliverables 或 video_number 字段为已与推广系列代码绑定的推广系列中的第一个订单设置了可交付内容（即视频）的数量要求，则同一推广系列中的后续订单需设置相同的可交付内容的数量要求。若您设置的可交付内容的数量要求与第一个订单不同，则将忽略该设置，并重置为第一个订单的设置。
	// 若您想要为新订单设置一个不同于现有订单的可交付内容数量要求，您需要在（与新推广系列代码绑定的）新推广系列下新建一个订单，而不能使用现有的推广系列代码。
	CampaignCode string `json:"campaign_code,omitempty"`
	// OrderSource 订单来源，用于区分订单是否为托管服务提供商（MSP）订单。
	// 枚举值：
	// MSP ：该订单是托管服务提供商订单。
	// OTHER ：该订单非托管服务提供商订单。
	OrderSource enum.TCMOrderSource `json:"order_source,omitempty"`
	// Deliverables 该订单所需的可交付内容（视频）的信息。
	// 最大数量：10。
	// 若您传入了deliverables，则至少需要指定其中的 deliverable_name 字段。
	// 若您同时传入deliverables和 video_number，则会忽略 video_number 。
	// 若未传入deliverables ，则会为订单自动生成基本信息（即可交付内容的名称）。例如，第一个可交付内容的名称将为 “Deliverable 1”，第二个可交付内容的名称将为“Deliverable 2”。
	Deliverables []Deliverable `json:"deliverables,omitempty"`
	// VideoNumber 若已设置 deliverables ，则无需传入本字段。否则，将忽略本字段。
	// 若想为每个可交付内容分别设置具体要求（例如名称和申请授权天数），而不仅仅是可交付内容的数量要求，请使用 deliverables 字段。
	// 订单所需视频的数量。
	// 默认值：10。
	// 取值范围：1-10。
	VideoNumber int `json:"video_number,omitempty"`
	// DefaultSparkAdsRequestedAuthorizationDays 若已通过 deliverables 为每个可交付内容指定了 spark_ads_requested_authorization_days ，则无需传入本字段。否则，将忽略本字段。
	// 若想为每个可交付内容设置不同的申请授权天数，而非为所有的可交付内容设置相同的申请授权天数，请使用 deliverables 字段。
	// 若仅为 deliverables 中的部分可交付内容指定了对应的 spark_ads_requested_authorization_days ，并同时传入了 default_spark_ads_requested_authorization_days ，则对于尚未指定 spark_ads_requested_authorization_days 的可交付内容，将使用由 default_spark_ads_requested_authorization_days 指定的默认申请授权天数。
	// 您想要向创作者请求将其所有视频在 Spark Ads 中推广的授权天数。
	// 允许的值：7，30，60，365。
	// 若设置了本字段，则当创作者接受订单时，创作者会自动授权您在 Spark Ads 中使用其全部视频。
	// 例如，若本字段设置为30，且创作者接受订单并上传多个视频作为交付内容，则您可在30天内将所有视频用于 Spark Ads。
	DefaultSparkAdsRequestedAuthorizationDays int `json:"default_spark_ads_requested_authorization_days,omitempty"`
}

// Deliverable 该订单所需的可交付内容（视频）的信息。
type Deliverable struct {
	// DeliverableName 若传入deliverables，则至少需要指定其中的deliverable_name。
	// 可交付内容的名称。该名称指定了可交付内容（即视频）的预期内容。
	// 长度限制：60个字符。
	DeliverableName string `json:"deliverable_name,omitempty"`
	// SparkAdsRequestedAuthorizationDays 您想要向创作者请求将其单个视频在 Spark Ads 中推广的授权天数。
	// 若不想请求 Spark Ads 授权，可将本字段设置为 0 或不传入本字段。
	// 允许的值：0，7，30，60，365。
	// 若设置了本字段，则当创作者接受订单时，创作者会自动授权您在 Spark Ads 中使用其单个视频。
	// 例如，若本字段设置为30，且创作者接受订单并上传一个视频作为交付内容，则您可在30天内将该视频用于 Spark Ads。
	SparkAdsRequestedAuthorizationDays int `json:"spark_ads_requested_authorization_days"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建 TTCM 订单 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *CreateResult `json:"data,omitempty"`
}

type CreateResult struct {
	// CampaignCode 推广系列代码
	CampaignCode string `json:"campaign_code,omitempty"`
	// InviteLink 与推广系列代码绑定的邀请链接
	InviteLink string `json:"invite_link,omitempty"`
	// RemainingOrders 该推广系列代码中剩余的订单配额。
	// 注意：
	// 单个推广系列代码中的最大订单配额为200。
	// 订单配额仅当创作者接受了与推广系列代码绑定的订单时才会扣除。
	RemainingOrders int `json:"remaining_orders,omitempty"`
	// Orders TTCM订单详情
	Orders []Order `json:"orders,omitempty"`
}
