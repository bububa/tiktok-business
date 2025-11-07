package video

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Video 视频信息
type Video struct {
	// Displayable 视频能否在平台中展示
	Displayable bool `json:"displayable,omitempty"`
	// Width 视频宽度
	Width int `json:"width,omitempty"`
	// Height 视频高度
	Height int `json:"height,omitempty"`
	// VideoCoverURL 视频封面临时URL。 有效期6小时，过期需重新获取。
	// URL中的x-expires参数后会包含过期时间，格式为Epoch/Unix时间戳，单位为秒。
	// 示例：http://p16-sign-sg.tiktokcdn.com/v0201/b99a388e3709470be5c~tplv-noop.image?x-expires=1671742348&x-signature=FziJhvED9NDTDmPofv3I%3D。
	VideoCoverURL string `json:"video_cover_url,omitempty"`
	// BitRate 码率，单位bps
	BitRate int64 `json:"bit_rate,omitempty"`
	// Format 视频格式
	Format string `json:"format,omitempty"`
	// PreviewURL 视频预览链接。有效期6小时，过期需重新获取。
	// 若想了解本预览链接的过期时间，请查看preview_url_expire_time。
	PreviewURL string `json:"preview_url,omitempty"`
	// PreviewURLExpireTime 视频预览链接过期时间，格式为YYYY-MM-DD HH:MM:SS(UTC+0)
	PreviewURLExpireTime model.DateTime `json:"preview_url_expire_time,omitzero"`
	// Duration 视频时长，单位秒
	Duration float64 `json:"duration,omitempty"`
	// Signature 视频文件MD5
	Signature string `json:"signature,omitempty"`
	// VideoID 视频ID，可用于广告投放中创建广告
	VideoID string `json:"video_id,omitempty"`
	// Size 视频大小，单位Byte
	Size int64 `json:"size,omitempty"`
	// MaterialID 素材ID
	MaterialID string `json:"material_id,omitempty"`
	// AllowedPlacements 视频可投放版位。受音乐版权限制，创意工具工具生成的部分素材只能投放到 TikTok 版位。投放到其他版位时将会导致审核不通过。枚举值详见枚举值-广告管理-版位。
	AllowedPlacements []enum.Placement `json:"allowed_placements,omitempty"`
	// AllowDownload 视频是否允许下载。受音乐版权限制，创意工具生成的部分素材只允许预览，禁止下载和传播
	AllowDownload bool `json:"allow_download,omitempty"`
	// FileName 视频名称
	FileName string `json:"file_name,omitempty"`
	// Defintion 视频的清晰度
	Definition string `json:"definition,omitempty"`
	// FPS 视频的每秒帧数（FPS）
	FPS int `json:"fps,omitempty"`
	// CreateTime 创建时间。UTC 时间，格式：2020-06-10T07:39:14Z
	CreateTime model.DateTime `json:"create_time,omitzero"`
	// ModifyTime 修改时间。UTC 时间，格式：2020-06-10T07:39:14Z
	ModifyTime model.DateTime `json:"modify_time,omitzero"`
	// FixTaskID 仅在请求中flaw_detect和 auto_fix_enabled都设置为true，且检测到视频中的问题时返回。
	// 修复任务 ID。
	// 若想获取修复后的视频，可使用以下方法之一：
	// 将本字段的值传给 /video/fix/task/get/ 接口的 task_id 字段。
	// 若请求中已将 auto_bind_enabled 设置为 true ，可以通过调用 /file/video/ad/search/ 检索修复后的视频。
	FixTaskID string `json:"fix_task_id,omitempty"`
	// FlawTypes 仅在请求中flaw_detect和 auto_fix_enabled都设置为true，且检测到视频中的问题时返回。视频问题种类。
	// 枚举值：
	// LOW_RESOLUTION：视频分辨率低于540 x 960 像素, 不符合要求。
	// ILLEGAL_VIDEO_SIZE：视频尺寸的比例不符合竖版（9:16），方版（1:1），横版（16:9）的要求。
	// NO_BGM（已废弃）：广告或视频没有背景音乐，或背景音乐不连贯/不清晰。
	// BLACK_EDGE（已废弃）：视频画面中出现影响用户体验的黑边。
	// ILLEGAL_DURATION（已废弃）：视频长度长于 60 秒或短于 5 秒，不符合要求。
	FlawTypes []enum.VideoFlawType `json:"flaw_types,omitempty"`
}
