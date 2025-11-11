package bc

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetAccountAuthorizationRequest 获取 TikTok 账号广告投放授权链接 API Request
type AssetAccountAuthorizationRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
}

// Encode implements GetRequest
func (r *AssetAccountAuthorizationRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// AssetAccountAuthorizationResponse 获取 TikTok 账号广告投放授权链接 API Response
type AssetAccountAuthorizationResponse struct {
	model.BaseResponse
	Data struct {
		// BcAuthQRCode TikTok 账号广告投放授权链接。
		// 获取授权链接后，您可以使用二维码生成器将其转换为二维码，并分享给 TikTok 账号所有者。TikTok 账号所有者扫描二维码后，即可将 Spark Ads 拉取的“现有帖子”权限和 Spark Ads 推送的“TikTok 帖子”权限授予给商务中心
		BcAuthQRCode string `json:"bc_auth_qr_code,omitempty"`
	} `json:"data"`
}
