package diagnostic

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/catalog/product"
	"github.com/bububa/tiktok-business/util"
)

// CatalogRequest 同步获取商品库商品诊断信息 API Reqyest
type CatalogRequest struct {
	// CatalogID 商品库ID。您需有访问该商品库的权限。
	CatalogID string `json:"catalog_id,omitempty"`
	// BcID 商务中心ID。该商务中心需为所指定商品库的所属商务中心，或所指定商品库已作为资产分享给该商务中心
	BcID string `json:"bc_id,omitempty"`
	// FeedID 更新源 ID。
	// 若未传入本字段，则使用商品库的默认更新源的ID。
	// 若想获取商品库所有更新源的诊断信息，需将本字段设置为 ALL 。
	// 若想获取某个商品库的所有更新源的ID，可使用 /catalog/feed/get/。
	FeedID string `json:"feed_id,omitempty"`
	// Filtering 筛选条件
	Filtering *CatalogFilter `json:"filtering,omitempty"`
	// Lang 想要为返回的 issue_title 和 reason_and_suggestion 字段设置的语言。
	// 默认值 en 。
	// 支持的枚举值请查看下文的lang枚举值小节。
	Lang string `json:"lang,omitempty"`
	// Page 当前页数。
	// 默认值：1。
	// 取值范围：≥1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值：10。
	// 取值范围：[1, 20]。
	PageSize int `json:"page_size,omitempty"`
}

// CatalogFilter 筛选条件
type CatalogFilter struct {
	// IssueLevel 用于筛选结果的问题级别。
	// 枚举值：
	// CRITICAL ：筛选需要立刻关注的严重问题。
	// WARNING ：筛选警告型问题。这些问题有对应的建议，可用于优化商品设置。
	IssueLevel enum.CatalogDiagnosticIssueLevel `json:"issue_level,omitempty"`
	// IssueCategory 用于筛选结果的问题类型。
	// 枚举值：
	// PRODUCT_ATTRIBUTES ：商品属性问题。
	// PRODUCT_REVIEW ：商品审核问题。
	// CATALOG ：商品库问题。
	// PIXEL_OR_EVENT ：Pixel 或事件问题。
	// FILE_UPLOAD_OR_FEED ：文件上传或更新源问题。
	IssueCategory enum.CatalogDiagnosticIssueCategory `json:"issue_category,omitempty"`
}

// Encode implements GetRequest interface
func (r *CatalogRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if r.FeedID != "" {
		values.Set("feed_id", r.FeedID)
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Lang != "" {
		values.Set("lang", r.Lang)
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

// CatalogResponse 同步获取商品库商品诊断信息 API Reqyest
type CatalogResponse struct {
	model.BaseResponse
	Data *CatalogResult `json:"data,omitempty"`
}

type CatalogResult struct {
	// DiagnosticDate 诊断信息生成的日期（UTC+0 时间），格式为 "YYYY-MM-DD" 。
	// 示例： "2023-06-15"
	DiagnosticDate string `json:"diagnositic_date,omitempty"`
	// Issues 有关商品库中发现的问题的诊断信息
	Issues []CatalogIssue `json:"issues,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// CatalogIssue 有关商品库中发现的问题的诊断信息
type CatalogIssue struct {
	// IssueID 问题ID。
	// 部分问题ID可传入/diagnostic/catalog/product/task/create/接口，用于创建对应问题诊断信息的异步下载任务。
	IssueID string `json:"issue_id,omitempty"`
	// IssueTitle 问题标题，即对问题的简要概括
	IssueTitle string `json:"issue_title,omitempty"`
	// ReasonAndSuggestion 原因和建议。
	// 本字段由问题的描述和如何解决该问题的建议组成。
	ReasonAndSuggestion string `json:"reason_and_suggestion,omitempty"`
	// IssueLevel 问题级别。
	// 枚举值：
	// CRITICAL ：该问题为需要立刻关注的严重问题。
	// WARNING：该问题为警告型问题，有对应的可用于优化商品设置的建议。
	IssueLevel enum.CatalogDiagnosticIssueLevel `json:"issue_level,omitempty"`
	// IssueCategory 问题类型。
	// 枚举值：
	// PRODUCT_ATTRIBUTES ：商品属性问题。
	// PRODUCT_REVIEW ：商品审核问题。
	// CATALOG ：商品库问题。
	// PIXEL_OR_EVENT：Pixel 或事件问题。
	// FILE_UPLOAD_OR_FEED ：文件上传或更新源问题。
	IssueCategory enum.CatalogDiagnosticIssueCategory `json:"issue_category,omitempty"`
	// IssueProductField 发现问题的商品字段
	IssueProductField string `json:"issue_product_field,omitempty"`
	// AffectedProductCount 发现问题的商品库商品数量
	AffectedProductCount int `json:"affected_product_count,omitempty"`
	// AffectedProductPercentage 发现问题的商品库商品所占百分比。
	// 取值范围：[0,100]。
	AffectedProductPercentage float64 `json:"affected_product_percentage,omitempty"`
	// ExampleAffectedProducts 仅当 affected_product_count 的值大于0时返回。
	// 最多包含10个有问题的商品示例的对象数组。您可以使用本对象数组中的字段详细分析发现的问题。
	// 本对象数组的结构与/catalog/product/get/返回中的list参数相同。需注意每个商品返回的具体字段根据商品类型变化。例如，仅酒店商品库的商品会返回 hotel_id 。
	ExampleAffectedProducts []product.Product `json:"example_affected_products,omitempty"`
}
