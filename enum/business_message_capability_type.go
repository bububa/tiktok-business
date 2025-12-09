package enum

// BusinessMessageCapabilityType the capability type that you want to query for the Business Account.
type BusinessMessageCapabilityType string

const (
	// BusinessMessageCapabilityType_IMAGE_SEND: The Business Account can send images and/or receive images in the specified conversation.
	BusinessMessageCapabilityType_IMAGE_SEND BusinessMessageCapabilityType = "IMAGE_SEND"
)
