package negativekeyword

import (
	"io"

	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// DownloadRequest 下载否定关键词 API Request
type DownloadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// 	ObjectType 对象类型。请通过本字段指定您想获取推广系列还是广告组对应的否定关键词列表。
	// 枚举值：
	// CAMPAIGN：推广系列。
	// ADGROUP：广告组。
	ObjectType string `json:"object_type,omitempty"`
	// ObjectID 对象ID。请通过本字段传入推广系列ID或广告组ID，从而获取对应的否定关键词列表。
	// 您需根据 object_type字段传入本字段：
	// 若 object_type设置为CAMPAIGN，您需通过本字段传入推广系列ID。
	// 若 object_type 设置为 ADGROUP , 您需通过本字段传入广告组ID。
	ObjectID string `json:"object_id,omitempty"`
}

// Encode implements GetRequest
func (r *DownloadRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("object_type", r.ObjectType)
	values.Set("object_id", r.ObjectID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// DownloadResponse 下载否定关键词 API Response
type DownloadResponse struct {
	model.BaseResponse
	w io.Writer
}

func NewDownloadResponse(w io.Writer) *DownloadResponse {
	return &DownloadResponse{w: w}
}

// Read implements DownloadResponse
func (resp *DownloadResponse) Read(r io.Reader) error {
	_, err := io.Copy(resp.w, r)
	return err
}
