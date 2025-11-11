package enum

// InvoicePayer 账单付费方
type InvoicePayer string

const (
	// InvoicePayer_AGENCY （代理商）
	InvoicePayer_AGENCY InvoicePayer = "AGENCY"
	// InvoicePayer_ADVERTISER （广告主）
	InvoicePayer_ADVERTISER InvoicePayer = "ADVERTISER"
)
