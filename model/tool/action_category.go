package tool

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ActionCategoryRequest 获取行为分类 API Request
type ActionCategoryRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：
	// 本字段对于注册地为美国或加拿大的广告主已全量。若您是注册地非美国或加拿大的广告主且想在定向美国或加拿大的广告中使用特殊广告分类，您需申请额外的白名单。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// Language 行为分类语言。
	// 默认值：en.
	// 目前支持的语言有： en, zh, ja, de, es, fr, id, it, ko, ru, th, tr, vi, ar, pt, ms。枚举值详见枚举值-语言。
	Language string `json:"language,omitempty"`
}

// Encode implements GetRequest interface
func (r *ActionCategoryRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if len(r.SpecialIndustries) > 0 {
		values.Set("special_industries", string(util.JSONMarshal(r.SpecialIndustries)))
	}
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ActionCategoryResponse 获取行为分类 API Response
type ActionCategoryResponse struct {
	model.BaseResponse
	Data struct {
		// ActionCategories 行为分类列表
		ActionCategories []ActionCategory `json:"action_categories,omitempty"`
	} `json:"data"`
}

// ActionCategory 行为分类
type ActionCategory struct {
	// Description 行为描述
	Description string `json:"description,omitempty"`
	// ActionCategoryID 行为 ID
	ActionCategoryID string `json:"action_category_id,omitempty"`
	// Level 层级
	Level int `json:"level,omitempty"`
	// SubCategoryIDs 子级行为ID分类
	SubCategoryIDs []string `json:"sub_category_ids,omitempty"`
	// Name 行为名称
	Name string `json:"name,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：
	// 如果返回值为空([])，则表示该分类没有支持的特殊广告分类。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// ActionScene 行为场景。
	// 枚举值：CREATOR_RELATED, VIDEO_RELATED
	ActionScene enum.ActionScene `json:"actiotn_scene,omitempty"`
}
