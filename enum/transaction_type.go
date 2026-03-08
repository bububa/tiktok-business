package enum

// BillingType Transaction type of the transfer.
type BillingType string

const (
	// BillingType_CASH: cash.
	BillingType_CASH BillingType = "CASH"
	// BillingType_GRANT: voucher.
	BillingType_GRANT BillingType = "GRANT"
	// BillingType_CREDIT: credit.
	BillingType_CREDIT BillingType = "CREDIT"
)
