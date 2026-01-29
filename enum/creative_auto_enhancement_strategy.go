package enum

// CreativeAutoEnhancementStrategy Automatic enhancements are real-time edits applied to your ads during your campaign. They can improve performance by creating more engaging and impactful visuals, sound, and more.
type CreativeAutoEnhancementStrategy string

const (
	// CreativeAutoEnhancementStrategy_TRANSLATE_AND_DUB: Translate and dub. Connect with global audiences by delivering your ad in 50+ languages.
	CreativeAutoEnhancementStrategy_TRANSLATE_AND_DUB CreativeAutoEnhancementStrategy = "TRANSLATE_AND_DUB"
	// CreativeAutoEnhancementStrategy_MUSIC_REFRESH: Music refresh. Stay on-trend by swapping in music currently popular on TikTok.
	CreativeAutoEnhancementStrategy_MUSIC_REFRESH CreativeAutoEnhancementStrategy = "MUSIC_REFRESH"
	// CreativeAutoEnhancementStrategy_VIDEO_QUALITY: Video quality. Improve overall visual quality with increased resolution and clarity.
	CreativeAutoEnhancementStrategy_VIDEO_QUALITY CreativeAutoEnhancementStrategy = "VIDEO_QUALITY"
	// CreativeAutoEnhancementStrategy_IMAGE_QUALITY: Image quality. Improve overall visual quality with increased resolution and clarity.
	CreativeAutoEnhancementStrategy_IMAGE_QUALITY CreativeAutoEnhancementStrategy = "IMAGE_QUALITY"
	// CreativeAutoEnhancementStrategy_IMAGE_RESIZE: Resize. Resize your image to take advantage of full-screen capabilities.
	CreativeAutoEnhancementStrategy_IMAGE_RESIZE CreativeAutoEnhancementStrategy = "IMAGE_RESIZE"
	// CreativeAutoEnhancementStrategy_CALL_TO_ACTION_ENHANCEMENT: CTA enhancement. Tailor and optimize your call to action based on specific use cases to enhance performance.
	CreativeAutoEnhancementStrategy_CALL_TO_ACTION_ENHANCEMENT CreativeAutoEnhancementStrategy = "CALL_TO_ACTION_ENHANCEMENT"
	// CreativeAutoEnhancementStrategy_AIGC_CARD: Generate ad card. Increase engagement by showcasing key messages and visuals designed to drive clicks and conversions.
	CreativeAutoEnhancementStrategy_AIGC_CARD CreativeAutoEnhancementStrategy = "AIGC_CARD"
)
