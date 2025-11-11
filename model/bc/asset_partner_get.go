package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetPartnerGetRequest 根据资产获取合作伙伴 API Request
type AssetPartnerGetRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetType 资产类型。
	// 枚举值：
	// ADVERTISER：广告账户（广告主账号）。
	// CATALOG：商品库。
	// TIKTOK_SHOP：TikTok Shop（商店）。
	// PIXEL：Pixel。
	// LEAD：线索。
	// TT_ACCOUNT：TikTok 账号。此资产类型对应 BC_AUTH_TT身份，详情见认证身份
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AssetID 资产 ID。
	// 若 asset_type 为 ADVERTISER，本字段代表广告账户 ID。
	// 若 asset_type 为 CATALOG，本字段代表商品库 ID。
	// 若 asset_type 为 TIKTOK_SHOP，本字段代表 TikTok Shop 的 ID。
	// 若 asset_type 为 PIXEL，本字段代表 Pixel ID。
	// 若 asset_type 为 LEAD，本字段代表表单库 ID。若想获取您有权访问的表单库的详情，可使用/page/library/get/。
	// 若 asset_type 为 TT_ACCOUNT，本字段代表商务中心中的 TikTok 账号的 ID。
	AssetID string `json:"asset_id,omitempty"`
	// Filtering 筛选条件
	Filtering *AssetPartnerGetFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// AssetPartnerGetFilter 筛选条件
type AssetPartnerGetFilter struct {
	// Keyword 想要筛选的合作伙伴名称关键词
	Keyword string `json:"keyword,omitempty"`
}

// Encode implements GetRequest
func (r *AssetPartnerGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("asset_type", string(r.AssetType))
	values.Set("asset_id", r.AssetID)
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

// AssetPartnerGetResponse 根据资产获取合作伙伴 API Response
type AssetPartnerGetResponse struct {
	model.BaseResponse
	Data *AssetPartnerGetResult `json:"data,omitempty"`
}

type AssetPartnerGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 合作伙伴列表
	List []AssetPartner `json:"list,omitempty"`
}

// AssetPartner 合作伙伴
type AssetPartner struct {
	// BcID 合作伙伴商务中心ID
	BcID string `json:"bc_id,omitempty"`
	// BcName 合作伙伴名
	BcName string `json:"bc_name,omitempty"`
}
