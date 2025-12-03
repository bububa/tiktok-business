package video

import (
	"io"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

type AdUploadRequest struct {
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// FileName 视频名称。
	// 长度限制： 1-100 个字符。
	// 默认值：
	// 若未启用智能修复（law_detect 和 auto_fix_enabled 设置为false ）：
	// 上传本地文件时，默认名称将是原始文件名。
	// 提供 URL 时，默认名称将是 URL 路径中的文件名部分。
	// 若原始文件名或 URL 路径中的文件名部分超过 100 个字符，将仅截取前 100 个字符。
	// 若启用了智能修复（law_detect 和 auto_fix_enabled 设置为true）：
	// 默认名称将是原始文件名（上传本地文件时）或 URL 路径中的文件名部分（提供 URL 时）。若您通过 file_name 指定了自定义视频名称，file_name 的值将作为修复后视频的名称。
	// 注意：同一个广告账户下的视频名称不允许重复。您可以调用/file/name/check/查验文件名称是否已存在。
	// 如果您收到了因为视频名称相同而导致的报错信息，请更改视频名称，或者在视频名称后添加时间戳来规避相同的视频名称（例如，以_)，然后再次上传。
	FileName string `json:"file_name,omitempty"`
	// UploadType 视频上传方式。
	// 默认值: UPLOAD_BY_FILE。
	// 枚举值: UPLOAD_BY_FILE，UPLOAD_BY_URL，UPLOAD_BY_FILE_ID， UPLOAD_BY_VIDEO_ID。
	// 注意：若本字段设置为 UPLOAD_BY_FILE，UPLOAD_BY_URL 或 UPLOAD_BY_FILE_ID, 则将返回新视频 ID（video_id）。若使用这三种方式将同一视频上传多次，每次上传均会生成新的视频 ID。
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// VideoFile 当upload_type为UPLOAD_BY_FILE 时必填。
	// 视频文件。
	// 推荐设置：
	// (1) 文件大小： 最大 500 MB。
	// (2) 比例：9:16, 16:9 和 1:1。
	// (3) 格式：.mp4, .mov, .mpeg, .avi。
	// 示例：'video_file=@"/Users/admin/Downloads/sample-mov-file.mov"'
	// 注意：上传视频前，请确保视频文件可播放且格式支持。
	VideoFile      io.Reader `json:"-"`
	VideoSignature string    `json:"video_signature,omitempty"`
	// VideoURL 当 upload_type 为UPLOAD_BY_URL 时必填。
	// 视频 URL 地址，如http://xxx.xxx。
	// (1) 文件大小：建议 10M 以内。
	// (2) 文件校验：校验返回的 header 中的Content-Type。常见的不支持的媒体类型为文本类型（Content-Type = text/*，例如text/html，text/plain）。此外，若返回的 header 中设置了Content -MD5，也会进行校验。
	// (3) 编码：URL需为浏览器中合法的URL。例如，URL中的空格需编码为%20。您可以将URL复制粘贴到浏览器的地址栏，随后URL会自动编码。
	// (4) 比例、格式、分辨率和码率限制同video_file 。
	// 注意：上传视频前，请确保视频URL可播放且格式支持。若URL无效或格式不支持（如txt格式），将会出现报错，或返回的视频ID无效，无法用于创建广告。
	VideoURL string `json:"video_url,omitempty"`
	// FileID 当upload_type为UPLOAD_BY_FILE_ID时必填。
	// 想上传的文件的file_id。
	// 可从文件上传接口获取。
	FileID string `json:"file_id,omitempty"`
	// VideoID 当upload_type 为UPLOAD_BY_VIDEO_ID时必填。
	// 视频 ID。
	// 注意：您需要传入TikTok广告管理平台或API上传的视频video_id。若您传入其他来源的视频ID，将会出现报错，或无法用上传的视频创建广告。
	// 您可以从/file/video/ad/upload/接口的返回值中，或者使用/file/video/ad/search/接口获取视频的video_id。
	VideoID string `json:"video_id,omitempty"`
	// IsThirdParty 视频是否为第三方视频
	IsThirdParty bool `json:"is_third_party,omitempty"`
	// FlawDetect 是否自动检测视频的潜在问题。
	// 默认值： false。
	FlawDetect bool `json:"flaw_detect,omitempty"`
	// AutoFixEnabled 是否自动修复检测到的问题。
	// 默认值 : false。
	// 若检测到视频存在的问题，则：
	// auto_fix_enabled 设置为 false时，会返回报错信息，提示问题类型 。
	// auto_fix_enabled设置为 true时，会仅自动修复 LOW_RESOLUTION 和 ILLEGAL_VIDEO_SIZE 类型的问题，并返回 fix_task_id和flaw_types字段。
	// 注意：
	// 若视频中检测到 LOW_RESOLUTION 问题，视频分辨率将自动增加至 1280x720 像素的标准分辨率，也称为 720p。
	// 若视频中检测到 ILLEGAL_VIDEO_SIZE 问题，视频将自动调整为以下标准宽高比之一：1:1（方版），9:16（竖版），或 16:9（横版）。
	AutoFixEnabled bool `json:"auto_fix_enabled,omitempty"`
	// AutoBindEnabled 此字段只在flaw_detect和auto_fix_enabled 均设置为true时生效。
	// 是否自动将修复后的视频上传至素材库。
	// 默认值： false。
	AutoBindEnabled bool `json:"auto_bind_enabled,omitempty"`
}

type AdPostUploadRequest AdUploadRequest

// Encode implements PostRequest interface
func (r *AdPostUploadRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// Encode implements UploadRequest interface
func (r *AdUploadRequest) Encode() []model.UploadField {
	ret := make([]model.UploadField, 0, 9)
	ret = append(ret, model.UploadField{
		Key:   "advertiser_id",
		Value: r.AdvertiserID,
	}, model.UploadField{
		Key:   "upload_type",
		Value: string(enum.UPLOAD_BY_FILE),
	}, model.UploadField{
		Key:   "file_name",
		Value: r.FileName,
	}, model.UploadField{
		Reader: r.VideoFile,
		Key:    "video_file",
		Value:  r.FileName,
	}, model.UploadField{
		Key:   "video_signature",
		Value: r.VideoSignature,
	})
	if r.IsThirdParty {
		ret = append(ret, model.UploadField{
			Key:   "is_third_party",
			Value: "true",
		})
	}
	if r.FlawDetect {
		ret = append(ret, model.UploadField{
			Key:   "flaw_detect",
			Value: "true",
		})
	}
	if r.AutoFixEnabled {
		ret = append(ret, model.UploadField{
			Key:   "auto_fix_enabled",
			Value: "true",
		})
	}
	if r.AutoBindEnabled {
		ret = append(ret, model.UploadField{
			Key:   "auto_bind_enabled",
			Value: "true",
		})
	}
	return ret
}

// AdUploadResponse 上传视频 API Response
type AdUploadResponse struct {
	model.BaseResponse
	Data []Video `json:"data,omitempty"`
}
