package music

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// UploadRequest 上传音乐 API Request
type UploadRequest struct {
	// AdvertiserID 广告账户 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// MusicScene 音乐的使用场景。
	// 枚举值：
	// CREATIVE_ASSET：创意工具场景。
	// CAROUSEL_ADS：非购物广告类型的轮播广告场景。
	// CATALOG_CAROUSEL：购物广告类型的轮播广告场景。
	// 默认值：CREATIVE_ASSET
	MusicScene enum.MusicScene `json:"music_scene,omitempty"`
	// UploadType 当 music_scene 设置为 CAROUSEL_ADS或 CATALOG_CAROUSEL时，需传入 upload_type 和 material_action 其中之一。
	// 当 music_scene 设置为 CREATIVE_ASSET 或未传入时，本字段非必传，默认值为UPLOAD_BY_FILE。
	// 音乐上传方式。
	// 枚举值:
	// UPLOAD_BY_FILE：通过文件将音乐上传。
	// UPLOAD_BY_FILE_ID： 通过文件ID将音乐上传
	UploadType enum.UploadType `json:"upload_type,omitempty"`
	// MaterialAction 当 music_scene 设置为 CAROUSEL_ADS 或 CATALOG_CAROUSEL时，需传入 upload_type 和 material_action 其中之一。
	// 当 music_scene 设置为 CREATIVE_ASSET 或未传入 music_scene 时不支持本字段。
	// 您想要对音乐进行的操作。您需同时传入 material_id 字段。
	// 枚举值：
	// ADD_TO_LIKED：将音乐添加至收藏列表。
	// ADD_TO_HISTORY：将音乐添加至历史列表。
	// REMOVE_FROM_LIKED：将音乐从收藏列表中移除。
	// 注意：单个广告账户的收藏列表中至多可添加1,000首音乐数量。到达数量上限后，不支持再将此字段设置为 ADD_TO_LIKED 再向收藏列表中添加音乐
	MaterialAction enum.MaterialAction `json:"material_action,omitempty"`
	// MusicFile upload_type 为UPLOAD_BY_FILE或未传入upload_type 时必填。
	// 音乐文件。
	// 注意：当 music_scene 设置为 CAROUSEL_ADS或 CATALOG_CAROUSEL时，您通过本字段指定的自定义音乐文件需满足以下规格要求：
	// 文件类型：MP3。
	// 文件大小：不超过 10 MB。
	MusicFile io.Reader `json:"-"`
	// MusicSignature upload_type 设置为 UPLOAD_BY_FILE 或未传入 upload_type 时必填。
	// 音乐文件的MD5(用于服务端校验)
	MusicSignature string `json:"music_signature,omitempty"`
	// FileName 仅当传入 upload_type 时生效。
	// 音乐名称。
	// 长度限制： 1-100 个字符。缺省时为本地文件名称
	FileName string `json:"file_name,omitempty"`
	// FileID 当upload_type为UPLOAD_BY_FILE_ID时必填。
	// 想要上传的文件的file_id。可从文件上传接口获取。
	// 注意：当 music_scene 设置为 CAROUSEL_ADS或 CATALOG_CAROUSEL时，您通过本字段指定的ID对应的自定义音乐文件需满足以下规格要求：
	// 文件类型：MP3。
	// 文件大小：不超过 10 MB。
	FileID string `json:"file_id,omitempty"`
	// MaterialID 当 music_scene 设置为 CAROUSEL_ADS或 CATALOG_CAROUSEL且传入 material_action 时必填。
	// 当 music_scene 设置为 CREATIVE_ASSET 或未传入 music_scene 时不支持本字段。
	// 音乐的素材ID，用于将音乐添加至收藏列表或历史列表，或将音乐从收藏列表中移除。
	// 若想获取有效的素材ID，您可以：
	// 通过本接口上传可用于轮播广告的自定义音乐，或
	// 通过/file/music/get/返回的 material_id ，获取广告账号下可用于轮播广告的音乐对应的素材ID
	MaterialID string `json:"material_id,omitempty"`
}

type PostUploadRequest UploadRequest

// Encode implements PostRequest interface
func (r *PostUploadRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// Encode implements UploadRequest interface
func (r *UploadRequest) Encode() []model.UploadField {
	var reader io.Reader
	if r.MusicSignature == "" {
		buf := util.NewBuffer()
		defer util.ReleaseBuffer(buf)
		_, _ = io.Copy(buf, r.MusicFile)
		h := md5.New()
		h.Write(buf.Bytes())
		r.MusicSignature = hex.EncodeToString(h.Sum(nil))
		reader = buf
	} else {
		reader = r.MusicFile
	}
	ret := make([]model.UploadField, 0, 8)
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
		Key:   "music_signature",
		Value: r.MusicSignature,
	}, model.UploadField{
		Reader: reader,
		Key:    "music_file",
		Value:  r.FileName,
	})
	if r.MusicScene != "" {
		ret = append(ret, model.UploadField{
			Key:   "music_scene",
			Value: string(r.MusicScene),
		})
	}
	if r.MaterialAction != "" {
		ret = append(ret, model.UploadField{
			Key:   "material_action",
			Value: string(r.MaterialAction),
		})
	}
	if r.MaterialID != "" {
		ret = append(ret, model.UploadField{
			Key:   "material_id",
			Value: r.MaterialID,
		})
	}
	return ret
}

// UploadResponse 上传音乐 API Response
type UploadResponse struct {
	model.BaseResponse
	Data *Music `json:"data,omitempty"`
}
