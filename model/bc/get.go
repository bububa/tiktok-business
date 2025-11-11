package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// GetRequest 获取商务中心列表 API Request
type GetRequest struct {
	// BcID 商务中心ID。不传则默认返回用户所有的商务中心列表，传入时返回指定商务中心账户
	BcID string `json:"bc_id,omitempty"`
	// Page 当前页数，默认值: 1，取值范围: ≥ 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，默认值: 10，取值范围: 1-50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
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

// GetResponse 获取商务中心列表 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 商务中心列表
	List []BcWithUserRole `json:"list,omitempty"`
}

type BcWithUserRole struct {
	// BcInfo 商务中心
	BcInfo *BusinessCenter `json:"bc_info,omitempty"`
	// UserRole 商务中心用户角色。枚举值：ADMIN（管理员）， STANDARD（标准成员）
	UserRole string `json:"user_role,omitempty"`
	// ExtUserRole 用户在商务中心除基本角色外的扩展角色
	ExtUserRole *ExtUserRole `json:"ext_user_role,omitempty"`
}

// ExtUserRole 用户在商务中心除基本角色外的扩展角色
type ExtUserRole struct {
	// FinanceRole 用户财务角色. 枚举: MANAGER(财务经理)， ANALYST(财务分析师)
	FinanceRole enum.UserFinanceRole `json:"finance_role,omitempty"`
}
