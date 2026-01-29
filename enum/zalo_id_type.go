package enum

// ZaloIDType The type of Zalo contact format.
type ZaloIDType string

const (
	// ZALO_OFFICIAL_ACCOUNT: Zalo Official Account ID.
	ZALO_OFFICIAL_ACCOUNT ZaloIDType = "ZALO_OFFICIAL_ACCOUNT"
	// ZALO_PHONE_ACCOUNT: Zalo phone number.
	ZALO_PHONE_ACCOUNT ZaloIDType = "ZALO_PHONE_ACCOUNT"
)
