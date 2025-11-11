package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PixelLinkGetRequest 获取与pixel绑定的广告账户列表 API Request
type PixelLinkGetRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// PixelCode Pixel代码
	PixelCode string `json:"pixel_code,omitempty"`
}

// Encode implenents GetRequest
func (r *PixelLinkGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("pixel_code", r.PixelCode)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// PixelLinkGetResponse 获取与pixel绑定的广告账户列表 API Response
type PixelLinkGetResponse struct {
	model.BaseResponse
	Data *PixelLinkGetResult `json:"data,omitempty"`
}

type PixelLinkGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 广告账户列表
	List []Advertiser `json:"list,omitempty"`
}
