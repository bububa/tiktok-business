package creator

import "github.com/bububa/tiktok-business/model"

// GetResponse 通过访问令牌获取创作者 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// List TikTok 创作者列表
		List []Creator `json:"list,omitempty"`
	} `json:"data"`
}
