package enum

// UploadType 视频上传方式。
type UploadType string

const (
	UPLOAD_BY_FILE     UploadType = "UPLOAD_BY_FILE"
	UPLOAD_BY_URL      UploadType = "UPLOAD_BY_URL"
	UPLOAD_BY_FILE_ID  UploadType = "UPLOAD_BY_FILE_ID"
	UPLOAD_BY_VIDEO_ID UploadType = "UPLOAD_BY_VIDEO_ID"
)
