package photo

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PublishRequest 将照片帖子发布到自有账号 API Request
type PublishRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// PhotoImages 用于发布图片内容的可公开访问 HTTP(s) URL 列表，最多支持 35 张照片的 URL。
	// 图片限制：
	// 图片的最大分辨率为 1080 x 1920 像素或 1920 x 1080 像素。
	// 图片文件大小：每张照片不超过 20 MB。
	// 图片文件类型：JPG，JPEG 或 WebP。
	// 示例：https://tiktok.com/tiktok-images/image.jpg
	// 重要提示：
	// 您不能使用未经验证的图片 URL 发布照片帖子。为保证流畅的 API 体验，我们建议您尽快验证 URL 资源的所有权，并指定已手动或自动验证过所有权的 URL。欲手动验证 URL 资源的所有权，可查看管理 URL 资源列出的对应步骤。
	// 需注意若您在2023年4月7日前使用过/business/video/publish/ 发布视频，我们已将您通过video_url参数指定的视频 URL 对应的域名自动设置为已验证状态。因此，建议您使用 /business/property/list/ 接口，确认您用于存放图片的域名是否已自动验证过。
	PhotoImages []string `json:"photo_images,omitempty"`
	// PhotoCoverIndex 用作照片帖子封面的图片索引。
	// 默认值： 0 （使用第一张图片）。
	// 例如，若您想将第二张图片设置为帖子封面图，可将本字段设置为1
	PhotoCoverIndex int `json:"photo_cover_index,omitempty"`
	// PostInfo 照片帖子的有关信息
	PostInfo *PostInfo `json:"post_info,omitempty"`
}

// PostInfo 照片帖子的有关信息
type PostInfo struct {
	// PrivacyLevel 照片帖子的隐私级别。
	// 枚举值：
	// PUBLIC_TO_EVERYONE：所有人。
	// MUTUAL_FOLLOW_FRIENDS：好友（您回关的粉丝）。
	// FOLLOWER_OF_CREATOR：粉丝。
	// SELF_ONLY：仅自己。
	// 若想查询 TikTok 账号允许的隐私级别，可使用/business/video/settings/ 并查看返回的privacy_level_options 。
	PrivacyLevel enum.PrivacyLevelOption `json:"privacy_level,omitempty"`
	// Title 帖子标题。
	// 长度限制：90 个字符（UTF-16 编码）。
	Title string `json:"title,omitempty"`
	// Caption 视频文案/描述 - 包含自有 TikTok 账号的（互关） 好友的话题标签（#hashtags）和提及（@mentions）。
	// 长度限制：4,000 个字符（UTF-16 编码），其中最多包含30个提及。
	// 您可在帖子描述中使用换行符"\n"，从而在 TikTok 手机端应用折行显示多行视频描述。 为使换行功能生效，文案必须包含足够的字符数从而触发截断。这种截断基于当前手机屏显设置和操作系统，截断标志为省略号（...）和单词“more”，用户可以点击该单词从而展开并查看完整的文案。
	// 若文案太短而不能触发截断，则将忽略“\n”字符，文案将以单行显示，无任何换行。
	// 需注意 TikTok 网页端应用中将忽略"\n"。
	// 示例：Me and me casa 🏠 #cat #kitten #britishshorthair
	// 注意：帖子描述中包含的任何话题标签和提及将会在 TikTok 应用中显示为纯文本格式。
	Caption string `json:"caption,omitempty"`
	// AutoAddMusic 是否在帖子中自动添加推荐音乐。
	// 枚举值：true，false。 默认值：false。
	// 将本字段设置为 true 后，您可以在 TikTok 应用中自主更换使用的音乐。
	AutoAddMusic bool `json:"auto_add_music,omitempty"`
	// IsBrandOrganic 是否为该视频开启原生品牌内容开关。
	// 本字段若设置为 true ，则视频将被标记为原生品牌内容，表示您在推广自己或自己的业务。视频上将附加一个"Promotional content"（推广内容）标签。
	// 注意：
	// 若同时将 is_brand_organic 和 is_branded_content 设置为 true，则 is_brand_organic 将被忽略，视频将被标记为品牌内容。
	IsBrandOrganic bool `json:"is_brand_organic"`
	// IsBrandedContent 是否为该视频开启品牌内容开关。
	// 本字段若设置为true ，则视频将被标记为品牌内容，表示您与某个品牌是合作伙伴赞助商关系。视频上将附加一个"Paid partnership"（合作伙伴赞助）标签。
	// 注意：
	// 若同时将 is_brand_organic 和 is_branded_content 设置为 true，则 is_brand_organic 将被忽略，视频将被标记为品牌内容。
	IsBrandedContent bool `json:"is_branded_content,omitempty"`
	// IsDraft 是否将帖子上传为草稿。
	// 支持的值：true，false。
	// 默认值：false。
	// 若您将本字段设置为 true，将忽略 post_info 中除 title 和 caption 的其他字段。
	// 注意: 若您想使用本字段将帖子上传为草稿，需先遵循授权和获取访问令牌中的步骤获取了带有 video.upload 权限的 TikTok 账号访问令牌以及 TikTok 账号对应的应用唯一 ID。
	// 若要确认 TikTok 账号访问令牌是否带有 video.upload 权限，可使用 /tt_user/token_info/get/，查看返回的 scope。若访问令牌无 video.upload 权限，需根据授权步骤重新向 TikTok 账号所有人请求草稿上传权限。
	IsDraft bool `json:"is_draft,omitempty"`
	// DisableComment 是否禁止对所发布的照片帖子发表评论。
	// 支持的值：
	// false：不禁止。
	// true：禁止。
	// 若想确认某个 TikTok 账号的帖子可设置的值，可使用 /business/video/settings/ 并查看返回的 comment_disabled 。
	// 若 comment_disabled 为 true，本字段仅可设置为true。
	// 若 comment_disabled 为 false，本字段可设置为true 或false。
	// 示例：false
	// 默认值：false
	DisableComment bool `json:"disable_comment,omitempty"`
}

// Encode implements PostRequest
func (r *PublishRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PublishResponse 将照片帖子发布到自有账号 API Response
type PublishResponse struct {
	model.BaseResponse
	Data struct {
		// ShareID 照片帖子发布任务的唯一标识符。
		// 您可将此 ID 传入 /business/publish/status/ 中的publish_id，从而查询照片帖子的发布状态。
		// 示例：p_pub_url~v1.2345123456789123456
		// 注意：
		// 针对已分享图片内容的帖子发布状态 Webhook 将包含此开发者应用标识符，以监测帖子上传的状态。帖子发布完成后 ，成功运行的 Webhook 也将包含为该帖子新建的唯一标识符，该标识符可用于获取帖子洞察信息。
		// 若想为您的开发者应用配置回调 URL，可使用 Webhook API 接口 /business/webhook/update/。
		ShareID string `json:"share_id,omitempty"`
	} `json:"data"`
}
