package creative

// AdPreviewCreateRequest 广告预览 API Request
type AdPreviewCreateRequest struct {
// AdvertiserID 广告主 ID
// 注意：您需对广告账户有操作员或管理员权限。若想查看您对某个广告账户的权限，可使用/bc/asset/get/，返回的 advertiser_role 需为 ADMIN 或 OPERATOR
AdvertiserID string `json:"advertiser_id,omitempty"`
// PreviewType 预览类型。
// 若想根据广告创建参数预览广告，需传入ADS_CREATION
PreviewType enum.AdPreviewType `json:"preview_type,omitempty"`
}
