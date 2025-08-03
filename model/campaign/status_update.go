package campaign

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// StatusUpdateRequest 更新推广系列操作状态 API Request
type StatusUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignIDs 推广系列ID列表。允许数量范围：1-20
	CampaignIDs []string `json:"campaign_ids,omitempty"`
	// OperationStatus 要进行的操作。
	// 枚举值：DELETE（删除）,DISABLE（暂停）,ENABLE（启用）。
	// 注意：您无法修改已删除推广系列的状态。
	OperationStatus enum.OperationStatus `json:"operation_status,omitempty"`
	// PostbackWindowMode 仅当以下条件均满足时生效：
	// 推广系列（campaign_ids）为专属推广系列（campaign_type为IOS14_CAMPAIGN ）。
	// 你在请求中同时将operation_status设置为 DISABLE 。
	// 推广系列未设置postback_window_mode 。
	// 您已为应用启用了 SKAN 4.0。
	//
	// 选定的模式决定您想要确保能接收到的 SKAN (SKAdNetwork) 4.0 回传。对应的窗口期较长的回传需要更多时间接收回传，推广系列配额释放的时间也相应较长。
	// 枚举值：
	// POSTBACK_WINDOW_MODE1：此模式确保能接收到第一次回传，对应的归因窗口期为 0-2 天。回传数据最长需要 4 天时间接收，推广系列配额释放的时间最长也为 4 天。
	// POSTBACK_WINDOW_MODE2：此模式确保能接收到第一次及第二次回传，对应的归因窗口期为 3-7 天。回传数据最长需要 13 天时间接收，推广系列配额释放的时间最长也为 13 天。
	// POSTBACK_WINDOW_MODE3：此模式确保能接收到三次回传，对应的归因窗口期为 8-35 天。回传数据最长需要 41 天时间接收，推广系列配额释放的时间最长也为 41 天。
	//
	// 若想了解 SKAN 4.0 归因窗口期和推广系列配额详情，请查看专属推广系列。
	//
	// 注意：
	//
	// 若您已为应用启用 Adjust，Airbridge，Appsflyer，Branch，Kochava，或 Singular 移动合作伙伴 (MMP) 跟踪，且您的 MMP SDK 版本已更新至支持 SKAN 4.0 的版本，您可在事件管理平台将应用迁移至 SKAN 4.0 。若想了解为应用启用 MMP 跟踪的步骤，请查看如何在 TikTok 广告管理平台中设置应用归因。若想了解将应用迁移至 SKAN 4.0 的步骤，请查看关于 SKAN 4.0 和 TikTok 以及如何迁移到 SKAN 4.0。
	// 本字段设置后不允许修改。
	// 若operation_status设置为ENABLE或未传入，同时传入了本字段，将出现报错。
	// 若您传入了本字段，但指定的应用（app_id）未启用 SKAN 4.0 ，将出现报错。
	// 若campaign_type 设置为IOS14_CAMPAIGN，operation_status 设置为 DISABLE，且您已为应用启用了 SKAN 4.0，但您未传入本字段，则本字段值默认设置为POSTBACK_WINDOW_MODE1。
	// 回传归因窗口期的设置可能仅部分成功。若您传入了非专属推广系列的campaign_ids，或传入了已经设置了postback_window_mode的专属推广系列的campaign_ids，将忽略这些推广系列，仅为其余满足上文所列条件的推广系列设置postback_window_mode。
	// 若您已为应用启用了 SKAN 4.0，需定向 iOS 16.1 及以上版本的设备，从而确保能够接收 SKAN 4.0 的回传。若您想仅定向 iOS 16.1 及以上版本的设备，可在广告组层级将 min_ios_version 设置为 16.1 。
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
}

// Encode implements PostRequest interface
func (r *StatusUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// StatusUpdateResponse 更新推广系列操作状态 API Response
type StatusUpdateResponse struct {
	model.BaseResponse
	Data *StatusUpdateResult `json:"data,omitempty"`
}

type StatusUpdateResult struct {
	// CampaignIDs 修改的推广系列 ID列表
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// Status 已执行的操作，枚举值：DELETE，DISABLE，ENABLE。
	Status enum.OperationStatus `json:"status,omitempty"`
	// CampaignList 所指定的推广系列的相关信息
	CampaignList []UpdatedCampaign `json:"campaign_list,omitempty"`
}

// UpdatedCampaign 所指定的推广系列的相关信息
type UpdatedCampaign struct {
	// CampaignID 推广系列ID
	CampaignID string `json:"campaign_id,omitempty"`
	// Status 该推广系列（ campaign_id ）当前的状态。
	// 枚举值：DELETE （已删除）， DISABLE （已暂停）， ENABLE （投放中）。
	Status enum.OperationStatus `json:"status,omitempty"`
	// PostbackWindowMode 选定的模式决定您想要确保能接收到的 SKAN (SKAdNetwork) 4.0 回传。
	// 枚举值：
	// POSTBACK_WINDOW_MODE1 ：此模式确保能接收到第一次回传，对应的归因窗口期为 0-2 天。 回传数据最长需要 4 天时间接收，推广系列配额释放的时间最长也为 4 天。
	// POSTBACK_WINDOW_MODE2 ：此模式确保能接收到第一次及第二次回传，对应的归因窗口期为 3-7 天。回传数据最长需要 13 天时间接收，推广系列配额释放的时间最长也为 13 天。
	// POSTBACK_WINDOW_MODE3 ：此模式确保能接收到三次回传，对应的归因窗口期为 8-35 天。回传数据最长需要 41 天时间接收，推广系列配额释放的时间最长也为 41 天。
	PostbackWindowMode enum.PostbackWindowMode `json:"postback_window_mode,omitempty"`
}
