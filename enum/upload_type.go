package enum

// UploadType 上传方式。
type UploadType string

const (
	UPLOAD_BY_FILE     UploadType = "UPLOAD_BY_FILE"
	UPLOAD_BY_URL      UploadType = "UPLOAD_BY_URL"
	UPLOAD_BY_FILE_ID  UploadType = "UPLOAD_BY_FILE_ID"
	UPLOAD_BY_VIDEO_ID UploadType = "UPLOAD_BY_VIDEO_ID"

	UploadType_FILE UploadType = "FILE"
	UploadType_URL  UploadType = "URL"
)

// UploadContentType 文件类型。
type UploadContentType string

const (
	// UploadContentTypeImage
	UploadContentTypeImage UploadContentType = "image"
	// UploadContentTypeMusic
	UploadContentTypeMusic UploadContentType = "music"
	// UploadContentTypeVideo
	UploadContentTypeVideo UploadContentType = "video"
	// UploadContentTypePlayable
	UploadContentTypePlayable UploadContentType = "playable"
)
