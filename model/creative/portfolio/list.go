package portfolio

import (
	"strconv"

	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// ListRequest 获取广告账户下的素材包列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Filtering 筛选条件。
	// 若未传入本字段，将获取所指定的广告账户下除行动引导文案素材包外所有类型的创意素材包。
	Filtering *ListFilter `json:"filtering,omitempty"`
	// Page 当前页数。
	// 默认值: 1。
	// 取值范围: ≥ 1。
	Page int `json:"page,omitempty"`
	// PageSize 分页大小。
	// 默认值: 20。
	// 取值范围: [1, 100]
	PageSize int `json:"page_size,omitempty"`
}

// ListFilter 筛选条件。
type ListFilter struct {
	// CreativePortfolioTypes 想要筛选的创意素材包类型。
	// 枚举值：
	// CARD：展示卡片素材包。
	// WEB_INFO_CARD：网站信息卡片素材包。
	// DOWNLOAD_CARD ：下载卡素材包。
	// PRODUCT_CARD：商品卡片素材包。
	// PRODUCT_TILE：商品磁贴素材包。
	// COUNTDOWN_STICKER：无提醒的倒计时贴纸素材包。
	// REMINDER_COUNTDOWN_STICKER： 含非直播活动提醒的倒计时贴纸素材包
	// LIVE_REMINDER_COUNTDOWN_STICKER：含直播活动提醒的倒计时贴纸素材包。
	// GIFTCODE_STICKER：礼品码贴纸素材包。
	// PREMIUM_BADGE：弹出展示素材包（福袋素材包）。
	// GESTURE：交互手势素材包。
	// SUPERLIKE：点赞彩蛋素材包。
	CreativePortfolioTypes []enum.CreativePortfolioType `json:"creative_portfolio_types,omitempty"`
	// CreativePortfolioIDs 想要筛选的创意素材包 ID 列表。
	// 最大数量：100。
	// 若想创建创意素材包并获取素材包 ID，可使用/creative/portfolio/create/。
	// 注意：不支持传入行动引导文案素材包的 ID 列表。
	CreativePortfolioIDs []string `json:"creative_portfolio_ids,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
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

// ListResponse 获取广告账户下的素材包列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// CreativePortfolios 广告账户下创意素材包的相关信息
	CreativePortfolios []Portfolio `json:"creative_portfolios,omitempty"`
}
