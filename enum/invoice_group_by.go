package enum

// InvoiceGroupBy 出账模式
type InvoiceGroupBy string

const (
	// InvoiceGroupBy_ACCOUNT: 合并出账
	InvoiceGroupBy_ACCOUNT InvoiceGroupBy = "ACCOUNT"
	// InvoiceGroupBy_ADVERTISER: 分别出账
	InvoiceGroupBy_ADVERTISER InvoiceGroupBy = "ADVERTISER"
)
