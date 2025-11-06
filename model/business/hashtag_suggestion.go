package business

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// HashtagSuggestionRequest 为 TikTok 账号获取推荐话题标签 API Request
type HashtagSuggestionRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// Keyword 要获取对应推荐话题标签的关键词
	Keyword string `json:"keyword,omitempty"`
	// Language 关键词语言。推荐话题标签的语言将与指定的关键词语言一致。
	// 默认值： en（英语）
	// 枚举值：
	// en：英语
	// ar：阿拉伯语
	// bn-IN：孟加拉语
	// ceb-PH：宿务语
	// cs-CZ：捷克语
	// de-DE：德语
	// el-GR：希腊语
	// es：西班牙语
	// fi-FI：芬兰语
	// fil-PH：菲律宾语
	// fr：法语
	// he-IL：希伯来语
	// hi-IN：印地语
	// hu-HU：匈牙利语
	// id-ID：印尼语
	// it-IT：意大利语
	// ja-JP：日语
	// jv-ID：爪哇语
	// km-KH：柬埔寨语
	// ko-KR：韩语
	// ms-MY：马来语
	// my-MM：缅甸语
	// nl-NL：芬兰语
	// pl-PL：波兰语
	// pt-BR：葡萄牙语
	// ro-RO：罗马尼亚语
	// ru-RU：俄语
	// sv-SE：瑞典语
	// th-TH：泰语
	// tr-TR：土耳其语
	// uk-UA：乌克兰语
	// ur：乌尔都语
	// vi-VN：越南语
	// zh-Hans：简体中文
	// zh-Hant-TW：繁体中文
	Language string `json:"language,omitempty"`
}

// Encode implements GetRequest
func (r *HashtagSuggestionRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("business_id", r.BusinessID)
	values.Set("keyword", r.Keyword)
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// HashtagSuggestionResponse 为 TikTok 账号获取推荐话题标签 API Response
type HashtagSuggestionResponse struct {
	model.BaseResponse
	Data struct {
		// Suggestions 推荐话题标签列表。共计返回十个推荐话题标签，话题标签先按精确匹配，其次按照浏览量排序。
		// 注意：若针对您传入的关键词无法匹配到推荐话题标签，将在name字段返回传入的关键词，且 view_count将显示为0。
		Suggestions []Hashtag `json:"suggestions,omitempty"`
	} `json:"data"`
}

// Hashtag 推荐话题标签
type Hashtag struct {
	// Name 推荐的话题标签名称
	Name string `json:"name,omitempty"`
	// ViewCount 推荐的话题标签对应的浏览量
	ViewCount int64 `json:"view_count,omitempty"`
}
