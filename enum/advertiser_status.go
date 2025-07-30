package enum

// AdvertiserStatus 广告账号状态。
type AdvertiserStatus string

const (
	// Advertiser_STATUS_DISABLE 该广告账户已关户
	Advertiser_STATUS_DISABLE AdvertiserStatus = "STATUS_DISABLE"
	// Advertiser_STATUS_PENDING_CONFIRM 申请待审核
	Advertiser_STATUS_PENDING_CONFIRM AdvertiserStatus = "STATUS_PENDING_CONFIRM"
	// Advertiser_STATUS_PENDING_VERIFIED 待验证
	Advertiser_STATUS_PENDING_VERIFIED AdvertiserStatus = "STATUS_PENDING_VERIFIED"
	// Advertiser_STATUS_CONFIRM_FAIL 审核不通过
	Advertiser_STATUS_CONFIRM_FAIL AdvertiserStatus = "STATUS_CONFIRM_FAIL"
	// Advertiser_STATUS_ENABLE 审核通过
	Advertiser_STATUS_ENABLE AdvertiserStatus = "STATUS_ENABLE"
	// Advertiser_STATUS_CONFIRM_FAIL_END CRM审核不通过
	Advertiser_STATUS_CONFIRM_FAIL_END AdvertiserStatus = "STATUS_CONFIRM_FAIL_END"
	// Advertiser_STATUS_PENDING_CONFIRM_MODIFY 修改待审核
	Advertiser_STATUS_PENDING_CONFIRM_MODIFY AdvertiserStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	// Advertiser_STATUS_CONFIRM_MODIFY_FAIL 修改审核不通过
	Advertiser_STATUS_CONFIRM_MODIFY_FAIL AdvertiserStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	// Advertiser_STATUS_LIMIT 用户被惩罚
	Advertiser_STATUS_LIMIT AdvertiserStatus = "STATUS_LIMIT"
	// Advertiser_STATUS_WAIT_FOR_BPM_AUDIT 等待CRM审核
	Advertiser_STATUS_WAIT_FOR_BPM_AUDIT AdvertiserStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	// Advertiser_STATUS_WAIT_FOR_PUBLIC_AUTH 待对公验证
	Advertiser_STATUS_WAIT_FOR_PUBLIC_AUTH AdvertiserStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	// Advertiser_STATUS_SELF_SERVICE_UNAUDITED 自助开户待验证资质
	Advertiser_STATUS_SELF_SERVICE_UNAUDITED AdvertiserStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	// Advertiser_STATUS_CONTRACT_PENDING 合同未生效
	Advertiser_STATUS_CONTRACT_PENDING AdvertiserStatus = "STATUS_CONTRACT_PENDING"
)
