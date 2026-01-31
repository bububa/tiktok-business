package enum

// AppTargetingType App targeting type.
type AppTargetingType string

const (
	// AppTargetingType_PROSPECT: Prospecting. Find prospective customers, including those who have not interacted with your products.
	AppTargetingType_PROSPECT AppTargetingType = "PROSPECT"
	// AppTargetingType_RETARGETING: Retargeting. Show ads to people who have already interacted with your business on and off TikTok.
	AppTargetingType_RETARGETING AppTargetingType = "RETARGETING"
)
