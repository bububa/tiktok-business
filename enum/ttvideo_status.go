package enum

// TTVideoStatus 视频状态
type TTVideoStatus string

const (
	// TTVideoStatus_INIT
	TTVideoStatus_INIT TTVideoStatus = "INIT"
	// TTVideoStatus_UPLOAD
	TTVideoStatus_UPLOAD TTVideoStatus = "UPLOAD"
	// TTVideoStatus_RELEASE
	TTVideoStatus_RELEASE TTVideoStatus = "RELEASE"
	// TTVideoStatus_DELETE
	TTVideoStatus_DELETE TTVideoStatus = "DELETE"
)
