package video

import (
	"github.com/bububa/tiktok-business/model"
	"github.com/bububa/tiktok-business/util"
)

// PublishRequest 将公开视频帖子发布到自有账号 API Request
type PublishRequest struct {
	// BusinessID TikTok 账户的应用唯一识别ID。
	// 您需将/tt_user/oauth2/token/接口返回的open_id字段值传给本字段。
	BusinessID string `json:"business_id,omitempty"`
	// VideoURL 用于发布视频内容的可公开访问 HTTP(s) URL。建议最低 TTL （存活时间）为 30 分钟。
	// 视频限制：
	// 视频内容必须采用 .mp4、.mov 或 .webm 格式。
	// 视频大小上限为 1 GB。
	// 视频时长不得短于 3 秒，不得长于 600 秒。
	// 若想查询 TikTok 账号允许的最大视频帖子时长，可使用 /business/video/settings/ 并查看返回的max_video_post_duration_sec。
	// 视频高度和宽度至少达到360 像素。
	// 视频帧率至少达到23 FPS，至多不超过60 FPS。
	// 示例： https://www.example.com/tiktok-videos/video.mp4
	// 重要提示：
	// 自2023年11月16日起，您将不能使用未经验证的视频 URL 发布视频。为保证流畅的API体验，我们建议您尽快验证 URL 资源的所有权，并指定已手动或自动验证过所有权的 URL。欲手动验证 URL 资源的所有权，可查看管理 URL 资源列出的对应步骤。
	// 2023年11月16日之后，若您想在/business/video/publish/ 中传入无需验证的测试视频 URL，从而发布测试视频，您可以使用测试 URL https://sf16-va.tiktokcdn.com/obj/eden-va2/uvpapzpbxjH-aulauvJ-WV[[/ljhwZthlaukjlkulzlp/3min.mp4。
	// 需注意若您在2023年4月7日前使用过/business/video/publish/ 发布视频，我们已将您通过video_url参数指定的视频 URL 对应的域名自动设置为已验证状态。因此，建议您使用 /business/property/list/ 接口，确认您用于存放视频的域名是否已自动验证过。
	VideoURL string `json:"video_url,omitempty"`
	// CustomThumbnailURL 自定义图片的可公开访问 HTTP(s) URL，用作视频封面。
	// 照片限制：
	// 最大分辨率：1080 x 1920 像素 或 1920 x 1080 像素。
	// 最小分辨率：360 x 360 像素。
	// 文件大小：不超过 20 MB。
	// 文件类型：JPG、JPEG、WebP 或 PNG。
	// 示例：https://www.example.com/tiktok-images/image.png。
	// 传入 custom_thumbnail_url 时，将忽略 thumbnail_offset。
	// 重要提示：指定已手动或自动验证过所有权的 URL。欲手动验证 URL 资源的所有权，可查看管理 URL 资源列出的对应步骤。
	CustomThumbnailURL string `json:"custom_thumbnail_url,omitempty"`
	// PostInfo 帖子相关信息。
	// 本字段必填，即使您不使用任何 post_info 中的额外参数，包括caption、is_brand_organic、is_branded_content、disable_comment、disable_duet、 disable_stitch、thumbnail_offset、 is_ai_generated和is_ads_only。在此情况下，需将本字段设置为空对象（{}）
	PostInfo *PostInfo `json:"post_info,omitempty"`
}

// PostInfo 帖子相关信息。
type PostInfo struct {
	// Caption 视频文案/描述 - 包含自有 TikTok 账号的（互关） 朋友的话题标签（#hashtags）和提及（@mentions）。
	// 您可在视频描述中使用换行符"\n"，从而在 TikTok 手机端应用折行显示多行视频描述。
	// 为使换行功能生效，文案必须包含足够的字符数从而触发截断。这种截断基于当前手机屏显设置和操作系统，截断标志为省略号（...）和单词“more”，用户可以点击该单词从而展开并查看完整的文案。
	// 若文案太短而不能触发截断，则将忽略“\n”字符，文案将以单行显示，无任何换行。
	// 需注意 TikTok 网页端应用中将忽略"\n"。
	// 字符数上限为 2,200 个字符（UTF-16 编码），其中最多包含30个提及。
	// 注意：视频描述中包含的任何话题标签和提及将会在 TikTok 应用中显示为纯文本格式。
	// 示例： Me and me casa 🏠 #cat #kitten #britishshorthair
	Caption string `json:"caption,omitempty"`
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
	// DisableComment 是否禁止对所发布的视频帖子发表评论。
	// 支持的值：
	// false：不禁止。
	// true：禁止。
	// 若想确认某个 TikTok 账号的帖子可设置的值，可使用 /business/video/settings/ 并查看返回的 comment_disabled 。
	// 若 comment_disabled 为 true，本字段仅可设置为true。
	// 若 comment_disabled 为 false，本字段可设置为true 或false。
	// 示例：false
	// 默认值：false
	DisableComment bool `json:"disable_comment,omitempty"`
	// DisableDuet 是否禁止对所发布的视频帖子进行合拍。
	// 支持的值：
	// false：不禁止。
	// true：禁止。
	// 若想确认某个 TikTok 账号的帖子可设置的值，可使用 /business/video/settings/ 并查看返回的 duet_disabled 。
	// 若 duet_disabled 为 true，本字段仅可设置为true。
	// 若 duet_disabled 为 false，本字段可设置为true 或false。
	// 示例：false
	// 默认值：false
	DisableDuet bool `json:"disable_duet,omitempty"`
	// DisableStitch 是否禁止对所发布的视频帖子进行拼接。
	// 支持的值：
	// false：不禁止。
	// true：禁止。
	// 若想确认某个 TikTok 账号的帖子可设置的值，可使用 /business/video/settings/ 并查看返回的 stitch_disabled 。
	// 若 stitch_disabled 为 true，本字段仅可设置为true。
	// 若 stitch_disabled 为 false，本字段可设置为true 或false。
	// 示例：false
	// 默认值：false
	DisableStitch bool `json:"disable_stitch,omitempty"`
	// ThumbnailOffset 用于选择所发布视频中的一帧作为视频封面的设置，格式为以毫秒为单位的时间戳。
	// 示例: 8000。
	// 默认值：0。
	// 注意:
	// 传入 custom_thumbnail_url 时，将忽略 thumbnail_offset。
	// 若您传入的时间戳超过视频时长，则会触发失败的Webhook事件，视频将无法发布。
	// 您传入的时间戳会向上取整获取最近的一帧。例如，若视频帧率为2帧每秒，时长为1秒，则传入时间戳499时会返回第1帧，传入时间戳500时会返回第2帧。
	ThumbnailOffset int64 `json:"thumbnail_offset,omitempty"`
	// IsAIGenerated 是否为视频帖子启用“ AI 生成的内容”开关。
	// 若您启用此开关，视频帖子发布后将带有“Creator labeled as AI-generated”（创作者标记为 AI 生成）标签且不允许修改。“创作者标记为 AI 生成”标签说明内容完全由 AI 生成或通过 AI 进行过大量编辑。
	// 支持的值： true，false。
	// 默认值： false。
	// 注意：除非违反《社区自律公约》，启用 AI 生成的内容设置将不会影响您的视频帖子的传播。
	IsAIGenerated bool `json:"is_ai_generated,omitempty"`
	// UploadToDraft 是否将帖子上传为草稿。
	// 支持的值：true，false。 默认值：false。
	// 若您将本字段设置为 true ，将忽略 post_info 中的其他字段。
	// 若您成功将帖子上传为草稿，可通过 TikTok 应用中的收件箱通知查看草稿。
	// 注意：若您想使用本字段将帖子上传为草稿，需先遵循授权和获取访问令牌中的步骤获取了带有video.upload权限的 TikTok 账号访问令牌以及 TikTok 账号对应的应用唯一 ID。
	// 若要确认 TikTok 账号访问令牌是否带有video.upload权限，可使用 /tt_user/token_info/get/，查看返回的scope。若访问令牌无video.upload权限，需根据授权步骤重新向 TikTok 账号所有人请求草稿上传权限。
	UploadToDraft bool `json:"upload_to_draft,omitempty"`
	// IsAdsOnly 是否将该视频发布为“仅限广告”的视频。
	// 支持的值：
	// true：将该视频发布为广告，该视频将不会在您的个人资料页上公开显示。
	// false：不将该视频发布为广告。
	// 默认值：false。
	// 注意：若“仅限广告”的视频已用作 Spark Ads，您可以通过/business/video/list/拉取这些视频的详情。
	IsAdsOnly bool `json:"is_ads_only,omitempty"`
}

// Encode implements PostRequest
func (r *PublishRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PublishResponse 将公开视频帖子发布到自有账号 API Response
type PublishResponse struct {
	model.BaseResponse
	Data struct {
		// ShareID 视频帖子发布任务的唯一标识符。
		// 您可将此 ID 传入 /business/publish/status/ 中的publish_id，从而查询视频帖子的发布状态。
		// 示例：v_pub_url~v1.2345123456789123456。
		// 注意：
		// 针对已分享视频内容的帖子发布状态 Webhook 将包含此开发者应用标识符，以监测视频上传的状态。视频发布完成后 ，成功运行的 Webhook 也将包含为该视频新建的唯一标识符，该标识符可用于获取视频洞察信息。
		// 若想为您的开发者应用配置回调 URL，可使用 Webhook API 接口 /business/webhook/update/。
		ShareID string `json:"share_id,omitempty"`
	} `json:"data"`
}
