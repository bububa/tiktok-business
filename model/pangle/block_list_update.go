package pangle

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// BlockListUpdateRequest 修改Pangle屏蔽列表 API Request
type BlockListUpdateRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// AddAppList 需要新增的 App 列表，iOS 应用为 iTunes ID，Android 应用为包名
	AddAppList []string `json:"add_app_list,omitempty"`
	// DeleteAppList 需要删除的 App 列表，iOS 应用为 iTunes ID，Android 应用为包名
	DeleteAppList []string `json:"delete_app_list,omitempty"`
	// ClearOldApp 是否清除原 App 列表。默认值: False。若为 True，将在原列表基础上删除及新增对应 App。若为 False，将清除原有列表，同时新增对应 App
	ClearOldApp bool `json:"clear_old_app,omitempty"`
}

// Encode implements PostRequest
func (r *BlockListUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BlockListUpdateResponse 修改Pangle屏蔽列表 API Response
type BlockListUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// SuccessCount 成功添加到屏蔽列表的App数量。从屏蔽列表中删除的App数量将不会被记入，无论删除成功与否
		SuccessCount int `json:"success_count,omitempty"`
	} `json:"data"`
}
