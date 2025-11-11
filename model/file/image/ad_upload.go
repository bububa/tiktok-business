package image

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AdUploadRequest 上传图片 API Request
type AdUploadRequest struct {
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// FileName 图片名称。
	// 长度限制： 1-100 个字符。
	//
	// 默认值为为文件名称。如果最终长度超过 100 个字符，将会自动截断。
	//
	// 重要提示: 同一个广告账户（advertiser_id）下的图片不能重名。您可以调用/file/name/check/查验文件名称是否已存在。
	// 如果上传图片时返回文件名重复的错误信息，请将图片文件重命名或添加时间戳（例如，以_)，然后再重新上传图片文件。
	FileName string `json:"file_name,omitempty"`
	// UploadType 图片上传方式。
	// 默认值: UPLOAD_BY_FILE。
	// 枚举值: UPLOAD_BY_FILE，UPLOAD_BY_URL，UPLOAD_BY_FILE_ID。
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// ImageFile upload_type为UPLOAD_BY_FILE时必填。
	// 图片文件。
	// 推荐设置:
	// 文件大小：小于 50 MB。
	// 图片分辨率：图片的总像素数，即宽度和高度的乘积，不可超过3,686,400像素（1440 x 2560）。例如，分辨率为 2400 x 1256 像素的图片支持上传，但分辨率为 3600 x 1256 像素的图片将触发报错。
	// 要了解特定用例的推荐图片分辨率，请查看"图片分辨率指南及使用案例"小节。
	// 文件类型: JPG，JPEG，或 PNG。
	ImageFile io.Reader `json:"-"`
	// ImageSignature upload_type为UPLOAD_BY_FILE时必填。
	// 图片的MD5（用于服务端校验）。
	ImageSignature string `json:"image_signature,omitempty"`
	// ImageURL upload_type为UPLOAD_BY_URL时必填。
	// 图片url地址，例如http://xxx.xxx。
	ImageURL string `json:"image_url,omitempty"`
	// FileID 当upload_type为UPLOAD_BY_FILE_ID时必填。
	// 上传图片的file_id。
	// 可从文件上传接口获取。
	FileID string `json:"file_id,omitempty"`
}

type AdPostUploadRequest AdUploadRequest

// Encode implements PostRequest interface
func (r *AdPostUploadRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// Encode implements UploadRequest interface
func (r *AdUploadRequest) Encode() []model.UploadField {
	var reader io.Reader
	if r.ImageSignature == "" {
		buf := util.NewBuffer()
		defer util.ReleaseBuffer(buf)
		_, _ = io.Copy(buf, r.ImageFile)
		h := md5.New()
		h.Write(buf.Bytes())
		r.ImageSignature = hex.EncodeToString(h.Sum(nil))
		reader = buf
	} else {
		reader = r.ImageFile
	}
	return []model.UploadField{{
		Key:   "advertiser_id",
		Value: r.AdvertiserID,
	}, {
		Key:   "upload_type",
		Value: string(enum.UPLOAD_BY_FILE),
	}, {
		Key:   "file_name",
		Value: r.FileName,
	}, {
		Reader: reader,
		Key:    "image_file",
		Value:  r.FileName,
	}, {
		Key:   "image_signature",
		Value: r.ImageSignature,
	}}
}

// AdUploadResponse 上传图片 API Response
type AdUploadResponse struct {
	model.BaseResponse
	Data *Image `json:"data,omitempty"`
}
