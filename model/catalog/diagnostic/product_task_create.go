package diagnostic

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ProductTaskCreateRequest 创建商品库商品诊断信息的异步下载任务 API Request
type ProductTaskCreateRequest struct {
	// CatalogID 商品库ID。您需有访问该商品库的权限。
	CatalogID string `json:"catalog_id,omitempty"`
	// BcID 商务中心ID。该商务中心需为所指定商品库的所属商务中心，或所指定商品库已作为资产分享给该商务中心
	BcID string `json:"bc_id,omitempty"`
	// FeedID 更新源 ID。
	FeedID string `json:"feed_id,omitempty"`
	// Lang 想要为返回的 issue_title 和 reason_and_suggestion 字段设置的语言。
	// 默认值 en 。
	// 支持的枚举值请查看下文的lang枚举值小节。
	Lang string `json:"lang,omitempty"`
	// IssueID 问题ID，用于下载对应诊断信息。
	// 您需指定一个可下载的问题的ID。可在下文的支持下载的问题小节的列表中选择问题ID。
	// 若未传入本字段，则返回所有发现的问题的诊断信息。
	IssueID string `json:"issue_id,omitempty"`
}

// Encode implements PostRequest interface
func (r *ProductTaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ProductTaskCreateResponse 创建商品库商品诊断信息的异步下载任务 API Response
type ProductTaskCreateResponse struct {
	model.BaseResponse
	Data struct {
		// TaskID 商品库诊断信息下载任务的ID
		TaskID string `json:"task_id,omitempty"`
	} `json:"data"`
}
