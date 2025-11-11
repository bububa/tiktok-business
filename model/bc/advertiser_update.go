package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserUpdateRequest 更新广告账户 API Request
type AdvertiserUpdateRequest struct {
	// AdvertiserID 广告主ID
	// 必传 advertiser_id 和 bc_id 其中之一。
	// 若想更新单个广告账号的非预算设置，传入advertiser_id。
	// 若想更新单个或多个广告账号的预算，传入bc_id
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName 广告账号名称
	AdvertiserName string `json:"advertiser_name,omitempty"`
	// Company 公司名称。
	// 若想确认是否支持通过 API 更新该广告账户的公司名称，可调用 /advertiser/info/ 并将fields设置为["company_name_editable"]。
	// 若返回的company_name_editable为true，您可使用本字段更新公司名称。
	// 若返回的company_name_editable为false，您需联系您的 TikTok 代表更新公司名称。
	Company string `json:"company,omitempty"`
	// ContactName 联系人姓名
	ContactName string `json:"contact_name,omitempty"`
	// ContactEmail 联系邮箱
	ContactEmail string `json:"contact_email,omitempty"`
	// ContactNumber 联系电话
	ContactNumber string `json:"contact_number,omitempty"`
	// PromotionLink 推广链接，长度不能超过 255 个字符
	PromotionLink string `json:"promotion_link,omitempty"`
	// LicenseNo 营业执照编号。
	// 注意：注册地为中国大陆，中国香港地区，或巴西，墨西哥时支持本字段
	LicenseNo string `json:"license_no,omitempty"`
	// LicenseImageID 营业执照照片 ID
	LicenseImageID string `json:"license_image_id,omitempty"`
	// QualificationImages 额外资质证明照片
	QualificationImages []QualificationImage `json:"qualification_images,omitempty"`
	// Address 收据上显示的公司地址。
	// 长度范围：1-1,024 字符
	Address string `json:"address,omitempty"`
	// TaxMap 计费和发票税号。不同国家采用不同的字段。例如，法国使用vat, 巴西使用 tax_id
	TaxMap map[string]string `json:"tax_map,omitempty"`
	// NeedSubmitCertificate 是否想要提交新的资质证明进行审核。
	// 支持的值：true，false。
	// 注意: 仅当更新以下任一字段时，将此字段设置为 true：
	// compnay
	// promotion_link
	// license_no
	// license_image_id
	// qualification_images
	// 如果上述字段均未发生变化，请省略此参数，以避免重新触发营业执照审核流程
	NeedSubmitCertificate *bool `json:"need_submit_certificate,omitempty"`
	// BcID 必传 advertiser_id 和 bc_id 其中之一。
	// 若想更新单个广告账号的非预算设置，传入advertiser_id。
	// 若想更新单个或多个广告账号的预算，传入bc_id。
	// 商务中心 ID。
	// 注意： 只有拥有财务经理权限（finance_role 为 MANAGER）的用户才能在商务中心执行更新预算和重置预算消耗操作。您可以使用/bc/member/invite/ 和 /bc/member/update/设置或更改成员权限。
	BcID string `json:"bc_id,omitempty"`
	// BudgetUpdateType 仅当传入 bc_id 时支持本字段。
	// 传入 advertiser_id 时不支持本字段。
	// 您想对预算执行的更改类型。
	// 枚举值：
	// UPDATE：更新预算模式和预算金额。
	// 若将本字段设置为 UPDATE，需传入advertiser_budgets 对象中的 advertiser_id 和 budget，budget_mode 为可选字段。
	// RESET：将广告账号的预算消耗重置为 0。
	// 仅支持对预算模式为自定义（budget_mode 为 CUSTOM_BUDGET）的广告账号执行重置预算消耗操作。
	// 若将本字段设置为 RESET，仅需传入advertiser_budgets 对象中的 advertiser_id。若传入了 budget_mode 和 budget 字段，将忽略budget_mode 和 budget。
	// INCREMENTAL_UPDATE：增加或减少当前广告账户预算。
	// ONE_CLICK_SET：将广告账户预算一键设置为最小值。此时无需传入对象数组 advertiser_budgets 中的 budget 字段。您可通过返回参数 one_click_set_amount 获取更改后的预算。
	// 默认值：UPDATE。
	// 注意：UPDATE和RESET只能对竞价广告账户使用。
	BudgetUpdateType enum.BudgetUpdateType `json:"budget_update_type,omitempty"`
	// AdvertiserBudgets 传入 bc_id 时，本字段必填。
	// 一个或多个广告账户要更新的预算设置信息。
	// 最大数量：
	// 当 budget_update_type 为 UPDATE 或 RESET 时：50。
	// 当 budget_update_type 为 INCREMENTAL_UPDATE 或 ONE_CLICK_SET 时：1。
	// 注意：每个广告账户的月度预算、日预算和自定义预算一天内最多更新 20 次
	AdvertiserBudgets []AdvertiserBudget `json:"advertiser_budgets,omitempty"`
}

// QualificationImage 额外资质证明照片
type QualificationImage struct {
	// ImageID 额外资质证明照片 ID
	ImageID string `json:"image_id,omitempty"`
}

// AdvertiserBudget 广告账户要更新的预算设置信息。
type AdvertiserBudget struct {
	// AdvertiserID 传入 advertiser_budgets 时，本字段必填。
	// 广告账户 ID。广告账号需属于所指定的 bc_id
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// BudgetMode 仅当 budget_update_type 为 UPDATE，RESET 或 ONE_CLICK_SET 时生效。
	// budget_update_type 为 INCREMENTAL_UPDATE 时，不支持本字段。
	// 广告账户预算模式。
	// 枚举值：
	// UNLIMITED：不限。广告账户无预算限制。
	// MONTHLY_BUDGET：月预算。广告账号在月度预算内消耗信用额度。
	// DAILY_BUDGET：日预算。广告账号在日预算内消耗信用额度。
	// CUSTOM_BUDGET：自定义预算。广告账号在单次自定义预算内消耗信用额度。
	// budget_update_type 为 UPDATE 时的默认值：MONTHLY_BUDGET。
	// 注意：
	// 若您未对广告账号的预算模式进行定制，广告账户的预算模式将默认为UNLIMITED。
	// 合约广告账户无法使用月度预算、日预算和自定义预算进行余额共享。
	BudgetMode enum.AdvertiserBudgetMode `json:"budget_mode,omitempty"`
	// Budget budget_mode 为 MONTHLY_BUDGET、DAILY_BUDGET 或 CUSTOM_BUDGET 时本字段必填。
	// 当 budget_update_type 为 UPDATE 或 RESET 时，本字段代表广告账户预算，保留两位小数。
	// 若budget_mode设置为UNLIMITED，则忽略此字段。
	// 若budget_mode设置为MONTHLY_BUDGET，传入月预算。
	// 若budget_mode设置为DAILY_BUDGET，传入日预算。
	// 若budget_mode设置为CUSTOM_BUDGET，传入自定义预算。
	// 当 budget_update_type 为 INCREMENTAL_UPDATE 时，本字段代表广告账户预算的修改金额。
	// 若想增加当前预算，需将 budget 设置为一个正数，例如 100。
	// 若想减少当前预算，需将 budget 设置为一个负数，例如 -10。
	// 当 budget_update_type 为 ONE_CLICK_SET 时，忽略本字段。
	// 注意： 若budget_mode 设置为MONTHLY_BUDGET/DAILY_BUDGET/CUSTOM_BUDGET，budget应不低于105%的当月/当日/当次消耗。您可以使用/report/integrated/get/接口来请求报表，获取广告账户的消耗情况（即指标spend）。
	Budget float64 `json:"budget,omitempty"`
}

// Encode implements PostRequest
func (r *AdvertiserUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdvertiserUpdateResponse 更新广告账户 API Response
type AdvertiserUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// List 仅在请求中指定了bc_id和advertiser_budgets时返回本字段。
		// 更新了预算设置的广告账号列表
		List []AdvertiserUpdateResult `json:"list,omitempty"`
	} `json:"data"`
}

// AdvertiserUpdateResult 更新了预算设置的广告账号
type AdvertiserUpdateResult struct {
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Status 更新的操作状态。
	// 枚举值：
	// SUCCESS：更新成功。
	// FAILED：更新失败
	Status string `json:"status,omitempty"`
	// ErrorMsg 报错信息。
	// status 为 SUCCESS 时，本字段值将为 null
	ErrorMsg string `json:"error_msg,omitempty"`
	// OneClickSetAmount 为该广告账户设置的最低预算金额。
	// 仅当请求中 budget_update_type 设置为 ONE_CLICK_SET 时返回非 null 值
	OneClickSetAmount float64 `json:"one_click_set_amount,omitempty"`
}

func (r AdvertiserUpdateResult) IsError() bool {
	return r.Status != "SUCCESS"
}

// Error implements Error
func (r AdvertiserUpdateResult) Error() string {
	return util.StringsJoin(r.AdvertiserID, ":", r.ErrorMsg)
}
