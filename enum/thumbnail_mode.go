package enum

// ThumbnailMode he mode of the video thumbnail.
type ThumbnailMode string

const (
	// ThumbnailMode_AUTOMATIC: Smart thumbnail. The system will show a different frame from your video customized for different audiences to closely match what they're searching for and what they're most likely to tap. The video cover specified through image_info will be used as the fallback video thumbnail if no matching thumbnails are found.
	ThumbnailMode_AUTOMATIC ThumbnailMode = "AUTOMATIC"
	// ThumbnailMode_MANUAL: Manual thumbnail. The video cover specified through image_info will be used as the video thumbnail.
	ThumbnailMode_MANUAL ThumbnailMode = "MANUAL"
)
