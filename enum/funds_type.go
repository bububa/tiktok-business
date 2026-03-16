package enum

// FundsType Fund type.
type FundsType string

const (
	// FUNDS_TYPE_CASH cash
	FUNDS_TYPE_CASH FundsType = "FUNDS_TYPE_CASH"
	// FUNDS_TYPE_GRANT coupon/voucher
	FUNDS_TYPE_GRANT FundsType = "FUNDS_TYPE_GRANT"
)
