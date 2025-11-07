package enum

// ProductVideoSpecificType 视频选择模式。
type ProductVideoSpecificType string

const (
	// ProductVideoSpecificType_AUTO_SELECTION：自动选择视频。
	ProductVideoSpecificType_AUTO_SELECTION ProductVideoSpecificType = "AUTO_SELECTION"
	// ProductVideoSpecificType_CUSTOM_SELECTION：手动选择视频。
	ProductVideoSpecificType_CUSTOM_SELECTION ProductVideoSpecificType = "CUSTOM_SELECTION"
	// ProductVideoSpecificType_UNSET：未设置。
	ProductVideoSpecificType_UNSET ProductVideoSpecificType = "UNSET"
)
