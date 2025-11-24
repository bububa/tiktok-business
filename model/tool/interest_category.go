package tool

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// InterestCategoryRequest 获取一般兴趣分类列表 API Request
type InterestCategoryRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Version 兴趣分类版本。
	// 枚举值：1(对应 interest_category), 2 (对应 interest_category_v2)。默认值: 2。
	Version int `json:"version,omitempty"`
	// Language 兴趣分类语言。
	// 默认值：en。
	// >目前支持的语言有： en, zh, ja, de, es, fr, id, it, ko, ru, th, tr, vi, ar, pt, ms。枚举值详见枚举值-语言
	Language string `json:"language,omitempty"`
	// Placements 投放版位。枚举值详见附录-版位。
	// 不同版位可选的兴趣分类可能不同，仅当 version为2 时有效
	Placements []enum.Placement `json:"placements,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值:
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：本字段对于注册地为美国或加拿大的广告主已全量。若您是注册地非美国或加拿大的广告主且想在定向美国或加拿大的广告中使用特殊广告分类，您需申请额外的白名单。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
}

// Encode implements GetRequest interface
func (r *InterestCategoryRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	if r.Version > 0 {
		values.Set("version", strconv.Itoa(r.Version))
	}
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if len(r.Placements) > 0 {
		values.Set("placements", string(util.JSONMarshal(r.Placements)))
	}
	if len(r.SpecialIndustries) > 0 {
		values.Set("special_industries", string(util.JSONMarshal(r.SpecialIndustries)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// InterestCategoryResponse 获取一般兴趣分类列表 API Response
type InterestCategoryResponse struct {
	model.BaseResponse
	Data struct {
		// InterestCategories 兴趣分类列表。
		// 若需确定游戏相关兴趣分类ID对应的兴趣分类名称，请查看兴趣分类。
		InterestCategories []InterestCategory `json:"interest_categories,omitempty"`
	} `json:"data"`
}

// InterestCategory 兴趣分类
type InterestCategory struct {
	// InterestCategoryID 兴趣分类ID
	InterestCategoryID string `json:"interest_category_id,omitempty"`
	// InterestCategoryName 兴趣分类名称
	InterestCategoryName string `json:"interest_category_name,omitempty"`
	// Level 层级
	Level int `json:"level,omitempty"`
	// SubCategoryIDs 子级兴趣分类 ID 列表
	SubCategoryIDs []string `json:"sub_category_ids,omitempty"`
	// SpecialIndustries 特殊广告分类。
	// 枚举值：
	// HOUSING：房地产，房屋保险，抵押贷款或相关的广告。
	// EMPLOYMENT：工作机会，实习机会，职业认证项目或相关的广告。
	// CREDIT：信用卡申请，汽车贷款，长期融资或相关的广告。
	// 注意：
	// 如果返回值为空([])，则表示该分类没有支持的特殊广告分类。
	SpecialIndustries []enum.SpecialIndustry `json:"special_industries,omitempty"`
	// Placements 该兴趣分类可用版位。返回空列表代表无限制。枚举值详见枚举值-版位
	Placements []enum.Placement `json:"placements,omitempty"`
}
