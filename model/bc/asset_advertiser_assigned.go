package bc

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// AssetAdvertiserAssignedRequest Get ad accounts linked to a TikTok account in Business Center API Request
type AssetAdvertiserAssignedRequest struct {
	// BcID 商务中心 ID。
	BcID string `json:"bc_id,omitempty"`
	// AssetID 商务中心内某个 TikTok 账号的 ID。
	// 如需获取商务中心内的 TikTok 账号列表，使用 /bc/asset/get/ 接口，并将 asset_type 设为 TT_ACCOUNT
	AssetID string `json:"asset_id,omitempty"`
	// AssetType Asset type.
	// Enum values:
	// TT_ACCOUNT: TikTok account as an ad delivery asset. You can run ads from these TikTok accounts.
	// MANAGED_BUSINESS_ACCOUNT: TikTok account for Business Account management. You can run ads from these TikTok accounts, and manage messages, leads, forms, comments, product links, and insights in your TikTok account.
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// AdvertiserID 要与该 TikTok 账号绑定的广告账号的 ID，广告账号需在同一商务中心中。
	// 如需获取商务中心内的广告账号列表，使用 /bc/asset/get/ 接口，并将 asset_type 设为 ADVERTISER
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Page Current page number.
	// Value range: ≥1.
	// Default value: 1.
	Page int `json:"page,omitempty"`
	// PageSize Page size.
	// Value range: 1-50.
	// Default value: 10.
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest
func (r *AssetAdvertiserAssignedRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("asset_id", r.AssetID)
	values.Set("asset_type", string(r.AssetType))
	values.Set("advertiser_id", r.AdvertiserID)
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

// AssetAdvertiserAssignedResponse Get ad accounts linked to a TikTok account in Business Center API Response
type AssetAdvertiserAssignedResponse struct {
	model.BaseResponse
	Data *AssetAdvertiserAssignedResult `json:"data,omitempty"`
}

type AssetAdvertiserAssignedResult struct {
	// List The list of ad accounts linked to the TikTok account (asset_id) in the Business Center.
	List []AssetAssignedAdvertiser `json:"list,omitempty"`
	// PageInfo Pagination information
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// AssetAssignedAdvertiser ad account linked to the TikTok account (asset_id) in the Business Center.
type AssetAssignedAdvertiser struct {
	// AdvertiserID The ID of the ad account
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AdvertiserName The name of the ad account
	AdvertiserName string `json:"advertiser_name,omitempty"`
}
