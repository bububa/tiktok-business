package insight

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// CategoryGetRequest 获取热门商品类目列表 API Request
type CategoryGetRequest struct {
	// BcID 商务中心 ID。
	// 若想获取您有权访问的商务中心列表，可使用/bc/get/。
	BcID string `json:"bc_id,omitempty"`
	// CatalogID 商品库 ID。
	// 需指定一个包含至少 20 个商品的电商商品库。
	// 若想获取某个商务中心中的电商商品库列表，可使用 /catalog/get/，筛选对应的 catalog_type 为 ECOM 的 catalog_id 值。
	// 若想验证商品库中包含的商品至少为 20 个，可使用 /catalog/overview/，查看返回的字段 approved，rejected 和 processing 的值之和是否大于等于 20。
	CatalogID string `json:"catalog_id,omitempty"`
	// Filtering 筛选条件
	Filtering *CategoryGetFilter `json:"filtering,omitempty"`
}

type CategoryGetFilter struct {
	// CategoryIDs 若传入了 filtering，则本字段必填。
	// TikTok 商品类目 ID 列表，可用于筛选分析数据。
	// 最大数量：50。
	// 若想获取经过筛选的洞察数据，可先调用一次 /catalog/insight/filter/get/ 获取 TikTok 的前 50 个热门商品类目，随后在返回中筛选 category_id 字段值传入本字段。
	CategoryIDs []string `json:"category_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r *CategoryGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("bc_id", r.BcID)
	values.Set("catalog_id", r.CatalogID)
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CategoryGetResponse 获取热门商品类目列表 API Response
type CategoryGetResponse struct {
	model.BaseResponse
	Data struct {
		// CategoryInsights TikTok 的前 50 个热门商品类目，列表按热门程度倒序排列。
		// 例如，列表中第一个商品类目排名第 1，即最热门的商品类目，第二个商品类目排名第 2，即第二热门的商品类目。
		CategoryInsights []Category `json:"category_insights,omitempty"`
	} `json:"data"`
}

// Category 商品类目
type Category struct {
	// CategoryID 商品的 TikTok 商品类目 ID，由三个类目层级及 # 号分隔符构成，格式为"第一层级的 ID#第二层级的 ID#第三层级的 ID"。
	// 示例：601152#842248#601302。
	CategoryID string `json:"category_id,omitempty"`
	// LevelInfo 类目所属层级的详情
	LevelInfo *CategoryLevelInfo `json:"level_info,omitempty"`
	// TotalProductCount 商品库（catalog_id）中与该商品类目（category_id）匹配的商品数量
	TotalProductCount int `json:"total_product_count,omitempty"`
	// ProductAvailabilityRage 商品库中与该类目匹配的商品的可售库存百分比。
	// 若本字段值过低，您可使用 /catalog/product/update/ 将相关商品的 availability 更新为 IN_STOCK。
	ProductAvailabilityRate float64 `json:"product_availability_rate,omitempty"`
}

// CategoryLevelInfo 类目所属层级的详情
type CategoryLevelInfo struct {
	// LevelID1 TikTok 商品类目第一层级的 ID
	LevelID1 string `json:"level_id_1,omitempty"`
	// LevelName1 TikTok 商品类目第一层级的名称
	LevelName1 string `json:"level_name_1,omitempty"`
	// LevelID2 TikTok 商品类目第二层级的 ID
	LevelID2 string `json:"level_id_2,omitempty"`
	// LevelName2 TikTok 商品类目第二层级的名称
	LevelName2 string `json:"level_name_2,omitempty"`
	// LevelID3 TikTok 商品类目第三层级的 ID
	LevelID3 string `json:"level_id_3,omitempty"`
	// LevelName3 TikTok 商品类目第三层级的名称
	LevelName3 string `json:"level_name_3,omitempty"`
}
