package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetGroupGetRequest 获取资产组信息 API Request
type AssetGroupGetRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetGroupID 您想要更新的资产组的ID
	AssetGroupID string `json:"asset_group_id,omitempty"`
	// QueryEntity 您想要获取的对象。枚举值：ASSET，MEMBER
	QueryEntity enum.AssetGroupEntity `json:"update_entity,omitempty"`
	// Filtering 筛选条件
	Filtering *AssetGroupGetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 取值范围：1-50。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

type AssetGroupGetFilter struct {
	// Keyword 筛选关键词。
	// 如果query_entity为ASSET，您可以传入asset_name的相关关键词进行模糊搜索，或传入asset_id进行精确搜索。
	// 如果query_entity为MEMBER，您可以传入user_name的相关关键词进行模糊搜索，或传入user_email进行精确搜索。
	Keyword string `json:"keyword,omitempty"`
}

// Encode implements GetRequest
func (r *AssetGroupGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("asset_group_id", r.AssetGroupID)
	values.Set("query_entity", string(r.QueryEntity))
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

// AssetGroupGetResponse 获取资产组信息 API Response
type AssetGroupGetResponse struct {
	model.BaseResponse
	Data *AssetGroupGetResult `json:"data,omitempty"`
}

type AssetGroupGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AssetGroup 资产组列表
	AssetGroup []AssetGroup `json:"asset_group,omitempty"`
}

// AssetGroup 资产组
type AssetGroup struct {
	// AssetGroupName 资产组名称
	AssetGroupName string `json:"asset_group_name,omitempty"`
	// AssetGroupID 资产组ID
	AssetGroupID string `json:"asset_group_id,omitempty"`
	// Assets 资产信息。当query_entity为ASSET时返回
	Assets []Asset `json:"assets,omitempty"`
	// Members 成员信息。当query_entity为ASSET时返回
	Members []AssetGroupMember `json:"members,omitempty"`
}

// AssetGroupMember 成员信息
type AssetGroupMember struct {
	// MebmerID 成员ID
	MemberID string `json:"member_id,omitempty"`
	// AssetRoles 您根据资产类型为成员分配的资产角色。
	// 键：资产类型。枚举值：ADVERTISER。
	// 值：分配的角色。枚举值：
	// ADVERTISER_ROLE_ADMIN（广告账户管理员）：
	// 管理广告账户：管理广告账户的设置、资金、权限。
	// 管理推广系列：创建、编辑广告。
	// 获取广告表现：获取广告报表，查看广告。
	// ADVERTISER_ROLE_OPERATOR（广告账户操作员）：
	// 管理推广系列：创建、编辑广告。
	// 获取广告表现：获取报表，查看广告。
	// ADVERTISER_ROLE_ANALYST（广告账户分析师）：
	// 获取广告表现：获取报表，查看广告。
	// 默认值："ADVERTISER: ADVERTISER_ROLE_ADMIN"。
	AssetRoles map[enum.AssetType]enum.AssetRole `json:"asset_roles,omitempty"`
}
