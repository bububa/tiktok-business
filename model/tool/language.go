package tool

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// LanguageRequest 获取受众语言列表 API Request
type LanguageRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest interface
func (r *LanguageRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// LanguageResponse 获取受众语言列表 API Response
type LanguageResponse struct {
	model.BaseResponse
	Data struct {
		// Languages 受众语言列表
		Languages []Language `json:"languages,omitempty"`
	}
}

// Language 受众语言列表
type Language struct {
	// Code 语言代码。您可将本字段的值将语言代码传入 /adgroup/create/ 中的 languages 字段用于受众语言定向。
	Code string `json:"code,omitempty"`
	// Name 语言名称。
	// 示例：English。
	Name string `json:"name,omitempty"`
}
