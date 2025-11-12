package dmp

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CustomAudienceFileUploadRequest 上传受众文件 API Request
type CustomAudienceFileUploadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// File 数据文件。仅支持.txt或者.csv文件
	File io.Reader `json:"-"`
	// FileSignature 文件的MD5，用于服务端校验文件完整性
	FileSignature string `json:"file_signature,omitempty"`
	// CaculateType 加密类型。加密类型需要和实际文件数据类型保持一致，否则上传将会失败，后续受众人群创建也会失败。枚举值详见枚举值-加密类型
	CaculateType enum.CalculateType `json:"caculate_type,omitempty"`
}

func (r *CustomAudienceFileUploadRequest) Encode() []model.UploadField {
	var reader io.Reader
	if r.FileSignature == "" {
		buf := util.NewBuffer()
		defer util.ReleaseBuffer(buf)
		_, _ = io.Copy(buf, r.File)
		h := md5.New()
		h.Write(buf.Bytes())
		r.FileSignature = hex.EncodeToString(h.Sum(nil))
		reader = buf
	} else {
		reader = r.File
	}
	return []model.UploadField{{
		Key:   "advertiser_id",
		Value: r.AdvertiserID,
	}, {
		Key:   "caculate_type",
		Value: string(r.CaculateType),
	}, {
		Reader: reader,
		Key:    "file",
		Value:  "file",
	}, {
		Key:   "file_signature",
		Value: r.FileSignature,
	}}
}

// CustomAudienceFileUploadResponse 上传受众文件 API Response
type CustomAudienceFileUploadResponse struct {
	model.BaseResponse
	Data struct {
		// FilePath 在TikTok内部的文件地址。
		// 注意：只有最初用于上传受众文件的广告主 ID，才能随后使用返回的 file_path 来创建受众
		FilePath string `json:"file_path,omitempty"`
	} `json:"data"`
}
