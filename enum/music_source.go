package enum

// MusicSource 音乐来源。
type MusicSource string

const (
	// MusicSource_USER 表示用户上传的音乐
	MusicSource_USER MusicSource = "USER"
	// MusicSource_SYSTEM 表示系统音乐
	MusicSource_SYSTEM MusicSource = "SYSTEM"
)
