package enum

// StoreStatus TikTok Shop 的状态。
type StoreStatus string

const (
	// StoreStatus_ACTIVE：可用。
	StoreStatus_ACTIVE StoreStatus = "ACTIVE"
	// StoreStatus_INACTIVE：不可用。
	StoreStatus_INACTIVE StoreStatus = "INACTIVE"
	// StoreStatus_NEW_CREATE：新建
	StoreStatus_NEW_CREATE StoreStatus = "NEW_CREATE"
)
