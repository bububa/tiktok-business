package enum

// BcTransferType How you'd like to process payments for an ad account or the Business Center.
type BcTransferType string

const (
	// BcTransferType_RECHARGE: Recharge money to an ad account or increase the credit balance of a Business Center.
	BcTransferType_RECHARGE BcTransferType = "RECHARGE"
	// BcTransferType_REFUND: Deduct money from an ad account or decrease the credit balance of a Business Center.
	BcTransferType_REFUND BcTransferType = "REFUND"
)
