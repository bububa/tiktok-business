package gmvmax

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/model/file/video"
	"github.com/bububa/tiktok-business/model/identity"
	"github.com/bububa/tiktok-business/util"
)

// CustomAnchorVideoListGetRequest 获取自定义作品中视频的详情 API Request
type CustomAnchorVideoListGetRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// CampaignCustomAnchorVideoID 商品 GMV Max 推广系列中创建的自定义的作品的合集 ID。
	// 若想获取与商品 GMV Max 推广系列绑定的自定义的作品的合集，可使用 /campaign/gmv_max/info/ 并查看返回的 campaign_custom_anchor_video_id
	CampaignCustomAnchorVideoID string `json:"campaign_custom_anchor_video_id,omitempty"`
	// CustomAnchorVideoList 要筛选的带有自定义商品链接的视频（即自定义的作品中的视频）。
	// 最大数量：20。
	// 若想获取与商品 GMV Max 推广系列绑定的自定义作品，可使用 /campaign/gmv_max/info/ 并查看返回的 custom_anchor_video_list
	CustomAnchorVideoList []AnchorVideo `json:"custom_anchor_video_list,omitempty"`
}

// Encode implements GetRequest
func (r *CustomAnchorVideoListGetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("advertiser_id", r.AdvertiserID)
	values.Set("campaign_custom_anchor_video_id", r.CampaignCustomAnchorVideoID)
	if len(r.CustomAnchorVideoList) > 0 {
		values.Set("custom_anchor_video_list", string(util.JSONMarshal(r.CustomAnchorVideoList)))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// CustomAnchorVideoListGetResponse 获取自定义作品中视频的详情 API Response
type CustomAnchorVideoListGetResponse struct {
	model.BaseResponse
	Data struct {
		// CustomAnchorVideoList 自定义的作品中的视频列表
		CustomAnchorVideoList []AnchorVideo `json:"custom_anchor_video_list,omitempty"`
	} `json:"data"`
}

// AnchorVideo 与该 GMV Max 推广系列绑定的自定义 TikTok 作品（即 TikTok 帖子）。
type AnchorVideo struct {
	// ItemID TikTok 帖子 ID
	ItemID string `json:"item_id,omitempty"`
	// Text TikTok 帖子描述。
	// 示例：summer
	Text string `json:"text,omitempty"`
	// SpuIDList 与该 TikTok 帖子绑定的商品 SPU ID 列表
	SpuIDList []string `json:"spu_id_list,omitempty"`
	// IdentityInfo 与该 TikTok 帖子绑定的认证身份的有关信息
	IdentityInfo *identity.Identity `json:"identity_info,omitempty"`
	// VideoInfo 帖子中视频的详情
	VideoInfo *video.Video `json:"video_info,omitempty"`
}
