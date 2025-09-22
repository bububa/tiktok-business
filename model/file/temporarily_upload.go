package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TemporarilyUploadRequest 上传单个文件 API Request
type TemporarilyUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// UploadType 文件上传方式。
	// 枚举值: FILE, URL。
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// ContentType 文件类型。
	// 枚举值：image，music，video，playable
	ContentType enum.UploadContentType `json:"content_type,omitempty"`
	// File 当upload_type为FILE时必填。
	// 您希望上传的文件。
	// 单个文件大小不可超过100MB。
	// 注意：若想上传可用于轮播广告的自定义音乐文件，您通过本字段指定的文件需满足以下规格要求：
	//
	// 文件类型：MP3。
	// 文件大小：不超过 10 MB。
	File io.Reader `json:"file,omitempty"`
	// URL 文件URL。格式为“http://xxx.xxx”。当upload_type为URL时必填。
	URL string `json:"url,omitempty"`
	// Signature 文件的的MD5(用于服务端校验)，当upload_type为FILE时必填
	Signature string `json:"signature,omitempty"`
	// Name 文件名称，字符上限为100。
	// 注意: 中文或日文每个字占用3个字符，每个英语字符占用1个字符。
	Name string `json:"name,omitempty"`
}

type TemporarilyPostUploadRequest TemporarilyUploadRequest

// Encode implements PostRequest interface
func (r *TemporarilyPostUploadRequest) Encode() []byte {
	r.UploadType = enum.UploadType_URL
	return util.JSONMarshal(r)
}

// Encode implements UploadRequest interface
func (r *TemporarilyUploadRequest) Encode() []model.UploadField {
	var reader io.Reader
	if r.Signature == "" {
		buf := util.NewBuffer()
		defer util.ReleaseBuffer(buf)
		io.Copy(buf, r.File)
		h := md5.New()
		h.Write(buf.Bytes())
		r.Signature = hex.EncodeToString(h.Sum(nil))
		reader = buf
	} else {
		reader = r.File
	}
	return []model.UploadField{{
		Key:   "advertiser_id",
		Value: r.AdvertiserID,
	}, {
		Key:   "upload_type",
		Value: string(enum.UploadType_FILE),
	}, {
		Key:   "content_type",
		Value: string(r.ContentType),
	}, {
		Key:   "name",
		Value: r.Name,
	}, {
		Reader: reader,
		Key:    "file",
		Value:  r.Name,
	}, {
		Key:   "signature",
		Value: r.Signature,
	}}
}

// TemporarilyUploadResponse 上传单个文件 API Response
type TemporarilyUploadResponse struct {
	model.BaseResponse
	Data *File `json:"data,omitempty"`
}
