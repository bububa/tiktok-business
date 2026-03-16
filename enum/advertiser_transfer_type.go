package enum

// AdvertiserTransferType Billing type
type AdvertiserTransferType string

const (
	// TRANS_TYPE_TRANSFER transfer
	TRANS_TYPE_TRANSFER AdvertiserTransferType = "TRANS_TYPE_TRANSFER"
	// TRANS_TYPE_TAX tax
	TRANS_TYPE_TAX AdvertiserTransferType = "TRANS_TYPE_TAX"
	// TRANS_TYPE_COST consumption
	TRANS_TYPE_COST AdvertiserTransferType = "TRANS_TYPE_COST"
	// TRANS_TYPE_PAYMENT payment
	TRANS_TYPE_PAYMENT AdvertiserTransferType = "TRANS_TYPE_PAYMENT"
)
