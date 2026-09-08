package enum

// BidProtectionDailyStatus The bid protection status on the date.
type BidProtectionDailyStatus string

const (
	// BidProtectionDailyStatus_UNDER_PROTECTION: Active. The campaign is active today; results are calculated on a daily basis.
	BidProtectionDailyStatus_UNDER_PROTECTION BidProtectionDailyStatus = "UNDER_PROTECTION"
	// BidProtectionDailyStatus_INELIGIBLE: Ineligible. Bid protection is ineligible for this date because this campaign or ad group status was paused or deleted.
	BidProtectionDailyStatus_INELIGIBLE BidProtectionDailyStatus = "INELIGIBLE"
	// BidProtectionDailyStatus_CONFIRMING: Pending. Bid protection results are pending and most ad credits will be updated in 8 to 10 days.
	BidProtectionDailyStatus_CONFIRMING BidProtectionDailyStatus = "CONFIRMING"
	// BidProtectionDailyStatus_PAYMENT_COMPLETE: Credit issued. Ad credits have been issued.
	BidProtectionDailyStatus_PAYMENT_COMPLETE BidProtectionDailyStatus = "PAYMENT_COMPLETE"
	// BidProtectionDailyStatus_TARGET_MET: Target met. The bid target was met, so no ad credits were issued.
	BidProtectionDailyStatus_TARGET_MET BidProtectionDailyStatus = "TARGET_MET"
)
