package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// TransferUploadRequest 进行分片上传 API Request
type TransferUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// UploadID 上传任务ID
	UploadID string `json:"upload_id,omitempty"`
	// Signature 分片文件MD5，用于服务端验证
	Signature string `json:"signature,omitempty"`
	// StartOffset 该分片的开始偏移量
	StartOffset int64 `json:"start_offset,omitempty"`
	// File 文件分片
	File io.Reader `json:"file,omitempty"`
}

// Encode implement UploadRequest interface
func (r TransferUploadRequest) Encode() []model.UploadField {
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
		Key:   "upload_id",
		Value: r.UploadID,
	}, {
		Key:   "start_offset",
		Value: strconv.FormatInt(r.StartOffset, 10),
	}, {
		Reader: reader,
		Key:    "file",
		Value:  "file",
	}, {
		Key:   "signature",
		Value: r.Signature,
	}}
}

// TransferUploadResponse 进行分片上传 API Response
type TransferUploadResponse struct {
	model.BaseResponse
	Data *TransferUploadResult `json:"data,omitempty"`
}

type TransferUploadResult struct {
	// StartOffset 下一个分片的开始偏移量
	StartOffset int64 `json:"start_offset,omitempty"`
	// EndOffset 下一个分片的结束偏移量
	// 当start_offset和end_offset相等的时候，意味着最后一个分片已经上传完毕。随后您需调用/file/finish/upload/完成分片上传。
	EndOffset int64 `json:"end_offset,omitempty"`
}

func (r TransferUploadResult) Complete() bool {
	return r.StartOffset == r.EndOffset
}
