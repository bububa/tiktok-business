package bc

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdvertiserCreateRequest 创建广告账户 API Request
type AdvertiserCreateRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您可访问的商务中心，可使用/bc/get/
	BcID string `json:"bc_id,omitempty"`
	// TiedToBillingGroup 是否将此账号添加到一个账单组。默认值: False。
	// 注意：该字段目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表
	TiedToBillingGroup bool `json:"tied_to_billing_group,omitempty"`
	// AdvertiserInfo 广告账户信息
	AdvertiserInfo *Advertiser `json:"advertiser_info,omitempty"`
	// CustomerInfo 业务信息
	CustomerInfo *Customer `json:"customer,omitempty"`
	// QualificationInfo 当商务中心类型为AGENCY或SELF_SERVICE_AGENCY时必填。
	// 资质信息。
	// 注意：当商务中心类型为DIRECT时，请不要填入该字段，否则会报错。
	QualificationInfo *QualificationInfo `json:"qualification_info,omitempty"`
	// ContactInfo 联系信息
	ContactInfo *ContactInfo `json:"contact_info,omitempty"`
	// BillingInfo 账单信息。注册地为法国、巴西或墨西哥的代理商创建广告账户或者创建的广告账户注册地为法国、巴西或墨西哥时必填。
	// 创建墨西哥广告账户代码示例： "billing_info":{"address":"xxxx","tax_field_dict": {"tax_id": "xxxx", "tax_regime": "xxxx"}}
	BillingInfo *BillingInfo `json:"billing_info,omitempty"`
	// BillingGroupInfo 账单组信息
	BillingGroupInfo *BillingGroup `json:"billing_group_info,omitempty"`
}

// Advertiser 广告主
type Advertiser struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Name 广告账户名称，长度不能超过100个字符
	Name string `json:"name,omitempty"`
	// CUrrency 广告账户币种，注意：需与商务中心保持一致，取值范围见附录-币种
	Currency string `json:"currency,omitempty"`
	// Timezone 广告账户时区代码，取值范围见附录-时区信息
	Timezone string `json:"timezone,omitempty"`
	// Type 您想创建的广告账户类型。枚举值: RESERVATION（合约广告）, AUCTION (竞价广告，默认值）。
	// 注意：DIRECT类型的商务中心只能创建AUCTION广告账户
	Type enum.ServiceType `json:"type,omitempty"`
}

// Customer 业务信息
type Customer struct {
	// Company 广告账户公司名称，长度不能超过255个字符
	Company string `json:"company,omitempty"`
	// Industry 广告账户行业ID，取值范围见附录-行业信息
	Industry string `json:"industry,omitempty"`
	// RegisteredArea 广告账户注册地代码， 取值范围见附录-地区代码
	RegisteredArea string `json:"registered_area,omitempty"`
}

// QualificationInfo 资质信息。
type QualificationInfo struct {
	// PromotionLink 当商务中心类型为AGENCY或SELF_SERVICE_AGENCY时必填。
	// 推广链接。该值必须为公司网站 URL。
	// 长度限制：255 个字符
	PromotionLink string `json:"promotion_link,omitempty"`
	// LicenseNo 广告账户注册地为中国大陆，中国香港地区，巴西或墨西哥时必填。
	// 营业执照编号。
	// 注意：若广告账户的注册地在中国大陆，则您可能需要为营业执照提交银联验证。若想查询您的营业执照是否需要银联验证，可使用 /bc/advertiser/unionpay_info/check/。若想了解银联验证和在商务中心创建广告账户的详情，可查看创建广告账户
	LicenseNo string `json:"license_no,omitempty"`
	// LicenseImageID 广告账户注册地为中国大陆或中国香港地区时必填。
	// 资质证书图片 ID。
	// 可由上传资质证书接口获取
	LicenseImageID string `json:"license_image_id,omitempty"`
	// QualificationImageIDs 广告账户或商务中心注册地为法国、巴西或墨西哥时必填。
	// 其他资质图片 ID。
	// 可由上传资质证书接口获取
	QualificationImageIDs []string `json:"qualification_image_ids,omitempty"`
	// QualificationID 资质 ID。
	// 您可以通过该字段复用其他广告账户的资质信息。
	// 资质 ID 可以通过/bc/advertiser/qualification/get/ 接口的返回获得。
	// 注意: 若传入该字段，则不能传入qualification_info下的其他字段
	QualificationID string `json:"qualification_id,omitempty"`
}

// ContactInfo 联系信息
type ContactInfo struct {
	// Name 联系人姓名，长度不能超过100个字符
	Name string `json:"name,omitempty"`
	// Email 联系人邮箱。注册地为法国、巴西或墨西哥的代理商创建广告账户或者创建的广告账户注册地为法国、巴西或墨西哥时必填
	Email string `json:"email,omitempty"`
	// Number 联系人电话
	Number string `json:"number,omitempty"`
}

// BillingInfo 账单信息
type BillingInfo struct {
	// Address 账单开票地址，长度不能超过512个字符
	Address string `json:"address,omitempty"`
	// TaxFieldDict 账单开票税号。不同国家使用不同的税号字段。
	// 欧洲国家，例如法国，使用税号字段vat 。
	// 巴西，墨西哥，阿联酋，埃及，沙特阿拉伯，以色列，土耳其，加拿大，美国使用税号字段tax_id。且对于巴西和墨西哥tax_id为必填字段。
	// 澳大利亚和新西兰使用税号字段abn
	TaxFieldDict map[string]string `json:"tax_field_dict,omitempty"`
}

// BillingGroup 账单组信息
type BillingGroup struct {
	// InvoiceGroupBy 出账模式。枚举值: ACCOUNT: 合并出账。ADVERTISER: 分别出账
	InvoiceGroupBy enum.InvoiceGroupBy `json:"invoice_group_by,omitempty"`
	// BillingGroupID 账单组ID。在invoice_group_by 为ACCOUNT时有效
	BillingGroupID string `json:"billing_group_id,omitempty"`
	// InvoicePayer 账单付费方。枚举 AGENCY （代理商）, ADVERTISER （广告主）。在广告账户注册地为FR（法国）时必填
	InvoicePayer enum.InvoicePayer `json:"invoice_payer,omitempty"`
}

// Encode implements PostRequest
func (r *AdvertiserCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdvertiserCreateResponse 创建广告账户 API Response
type AdvertiserCreateResponse struct {
	model.BaseResponse
	Data struct {
		// AdvertiserID 广告账户ID
		AdvertiserID string `json:"advertiser_id,omitempty"`
	} `json:"data"`
}
