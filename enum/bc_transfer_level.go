package enum

// BcTransferLevel Bc Transfer level.
type BcTransferLevel string

const (
	// BcTransferLevel_ADVERTISER: Recharge money to or deduct money from an ad account.
	BcTransferLevel_ADVERTISER BcTransferLevel = "ADVERTISER"
	// BcTransferLevel_BC: Increase or decrease the credit balance of the Business Center.
	BcTransferLevel_BC BcTransferLevel = "BC"
	// BcTransferLevel_PAYMENT_PORTFOLIO: The Payment Portfolio.
	BcTransferLevel_PAYMENT_PORTFOLIO BcTransferLevel = "PAYMENT_PORTFOLIO"
)
