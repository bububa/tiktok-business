package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// MemberGetRequest 获取成员列表 API Request
type MemberGetRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 最大值：20
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
	// Filtering 筛选条件
	Filtering *MemberGetFilter `json:"filtering,omitempty"`
}

// MemberGetFilter 筛选条件
type MemberGetFilter struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// UserRole 用户在商务中心的基础角色。枚举值：
	// ADMIN：管理员可以管理商务中心的成员和资产。
	// STANDARD：标准成员只能处理分配给他们的广告账户
	UserRole enum.UserRole `json:"user_role,omitempty"`
	// RelationStatus 用户与商务中心的关系状态，显示用户是否已加入商务中心。枚举值：
	// BOUND: 管理资产的申请已获拥有者批准。
	// PENDING: 管理资产的申请正在等待拥有者批准。
	// REJECTED: 管理资产的申请已被拥有者拒绝。
	RelationStatus enum.BcRelationStatus `json:"relation_status,omitempty"`
}

// Encode implements GetRequest
func (r *MemberGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// MemberGetResponse 获取成员列表 API Response
type MemberGetResponse struct {
	model.BaseResponse
	Data *MemberGetResult `json:"data,omitempty"`
}

type MemberGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 成员列表
	List []Member `json:"list,omitempty"`
}

// Member 成员
type Member struct {
	// UserID 用户ID
	UserID string `json:"user_id,omitempty"`
	// UserName 用户名称
	UserName string `json:"user_name,omitempty"`
	// UserRole 用户在商务中心中的角色
	UserRole enum.UserRole `json:"user_role,omitempty"`
	// RelationStatus 用户在商务中心的状态
	RelationStatus enum.BcRelationStatus `json:"relation_status,omitempty"`
	// AdvertiserRole 授予合作伙伴的广告账号角色。枚举值： ADMIN, OPERATOR, ANALYST。只有当asset_type 为ADVERTISER时返回
	AdvertiserRole enum.AdvertiserRole `json:"advertiser_role,omitempty"`
	// CatalogRole 授予合作伙伴的商品库角色。只有当asset_type 为 CATALOG时有效。枚举值：
	// ADMIN: 商务中心用户拥有商品库的管理权限。
	// AD_PROMOTE: 商务中心用户只能使用商品库进行推广活动。
	CatalogRole enum.CatalogRole `json:"catalog_role,omitempty"`
	// UserEmail 用户邮箱地址
	UserEmail string `json:"user_email,omitempty"`
	// ExtUserRole 用户在商务中心除基本角色外的扩展角色
	ExtUserRole *ExtUserRole `json:"ext_user_role,omitempty"`
}
