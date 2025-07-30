package event

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/util"
)

// TrackRequest 报告应用、网站、线下或 CRM 事件 API Request
type TrackRequest struct {
	// EventSource 本字段用于指定通过事件 API 上传的事件类型。
	// 枚举值：
	// web：在您的网站上发生的事件，通过 Pixel 代码追踪。
	// 若想获取 cURL、Python 和 Node.js 格式的代码示例，从而通过事件 API 2.0 报告网站事件，可使用网站事件的有效负载助手。
	// app：在您的应用程序上发生的事件，通过 TikTok 应用 ID 追踪。
	// offline：在实体店铺中发生的转化事件，通过线下事件组 ID 追踪。
	// crm：在 CRM 系统中发生的线索事件，通过 CRM 事件组 ID 追踪。
	// 注意：使用事件 API 2.0 报告应用事件目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	EventSource enum.EventSource `json:"event_source,omitempty"`
	// EventSourceID 用于追踪事件的事件源 ID。
	// 当event_source设置为web时，需通过本字段指定 Pixel 代码。
	// 若要获取 Pixel 代码，请查看常见问题-如何获取Pixel代码。
	// 当event_source设置为offline时，需通过本字段指定线下事件组 ID。
	// 若要获取线下事件组 ID，请查看线下事件设置指南。
	// 当event_source设置为app时，需通过本字段指定 TikTok 应用 ID。
	// 若要获取 TikTok 应用 ID，请查看应用事件设置指南。
	// 当event_source设置为crm时，需通过本字段指定 CRM 事件组 ID。
	// 若要获取 CRM 事件组 ID，请使用 /crm/list/。
	EventSourceID string `json:"event_source_id,omitempty"`
	// Data 您要报告的事件的对象数组。
	// 若只发送单个事件，请在本对象数组中传递一个对象。
	// 若要在单个有效负载中批量发送多个事件，则单个请求中的本对象数组最多可包含 1,000 个事件。
	// 若单个请求中本对象数组的事件数超过 1,000 个，整个请求将被拒绝。为了优化广告活动的性能，强烈建议您一旦在服务器上捕捉到事件，便尽快以实时方式（即非批量方式）发送事件。
	Data []Event `json:"data,omitempty"`
}

// Encode implements PostRequest interface
func (r *TrackRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// Event 您要报告的事件的对象数组。
type Event struct {
	// Event 网站事件、应用事件、线下事件和 CRM 事件均需指定本字段。
	// 转化事件名称。
	// 若为网站和应用事件，支持设置标准事件或自定义事件。
	// 若为线下事件，仅支持设置标准事件。
	// 若为 CRM 事件，仅支持设置自定义事件。
	// 支持的标准事件枚举值参见事件 API 2.0 - 支持的事件。
	// 若您要报告自定义事件，您需通过本字段传入一个自定义名称。
	// 请勿在自定义事件的名称中使用敏感词，以免为可能引发的后果承担相应的责任。
	Event enum.EventType `json:"event,omitempty"`
	// EventTime 网站事件、应用事件、线下事件和 CRM 事件均需指定本字段。
	// 事件发生的时间，以 Unix 时间戳表示，以秒为单位，对应时区为 UTC+0 。
	EventTime int64 `json:"event_time,omitempty"`
	// EventID 网站事件、线下事件和 CRM 事件均需指定本字段。
	// 在以下任意场景中本字段均为必填：
	// 同时从 TikTok 浏览器 Pixel 和事件 API 发送网站事件（event_source = web）时。
	// 同时通过事件 API 以及在事件管理器上上传 CSV 文件发送 CRM 事件（event_source = crm）时。
	// 用于标识唯一事件的 ID ，允许使用允许使用经散列处理的值或未经散列处理的值。
	// TikTok 使用event_source_id、event_id和event对从同一个来源或多个来源（例如浏览器 Pixel 和事件 API）多次发送的相同事件进行去重。 若想了解事件去重及如何完成对应设置的详情，请查看事件 API 2.0 - 事件去重。
	EventID string `json:"event_id,omitempty"`
	// User 网站事件、应用事件、线下事件和 CRM 事件均可指定本字段。
	// 有关客户的信息。
	User *User `json:"user,omitempty"`
	// Properties 网站事件、应用事件、线下事件和 CRM 事件均可指定本字段。
	// 有关产品和订单信息和其他附加信息，可用于优化广告表现。
	Properties *Properties `json:"properties,omitempty"`
	// Page 网站事件需指定本字段。
	// 有关网页的信息。
	Page *Page `json:"page,omitempty"`
	// App 应用事件需指定本字段。
	// 有关应用的信息。
	// 若想查看app对象中包含的字段，请查看下文的" app参数"小节。
	// 注意：使用事件 API 2.0 报告应用事件目前为白名单功能。如需使用此功能，请联系您的TikTok销售代表。
	App *App `json:"app,omitempty"`
	// Ad 应用事件可指定本字段。
	// 有关广告的信息。
	Ad *Ad `json:"ad,omitempty"`
	// LimitedDataUse 网站事件和应用事件可指定本字段。
	// 若想要了解有关限制数据使用功能的更多信息，请查看事件 API 2.0 - 限制数据使用
	LimitedDataUse bool `json:"limited_data_use,omitempty"`
	// Lead CRM 事件需指定本字段。
	// 有关线索的信息
	Lead *Lead `json:"lead,omitempty"`
}

// User 网站事件、应用事件、线下事件和 CRM 事件均可指定本字段。
// 事件 API 2.0 使用客户信息，例如 TikTok 点击 ID、高级匹配以及其他信号（如 IP 地址和用户代理），将网页访问者事件与 TikTok 用户进行匹配。一旦匹配成功，通过事件 API 分享的信息可用于构建定向受众，为优化算法提供数据，以及优化对推广系列效果的度量。
// 为了提高定向和优化模型的准确性，强烈建议包括多种类型的匹配数据。您可以使用高级匹配、TikTok 点击 ID（ttclid）和 Cookie 来对转化进行归因。
type User struct {
	// Ttclid 仅网站事件可指定本字段。
	// TikTok 点击 ID是一个跟踪参数，每当用户点击 TikTok 上的广告时，该参数就会添加到落地页 URL 的末尾。
	// ttclid的有效期限与您在归因管理器中设置的 CTA 窗口期相同。
	// 若想了解有关发送 TikTok 点击 ID 的详情，请查看事件 API 2.0 - 发送TikTok 点击 ID（ttclid）。
	Ttclid string `json:"ttclid,omitempty"`
	// Email 网站事件、应用事件、线下事件和 CRM 事件均可指定本字段。
	// 需传入经 SHA-256 散列处理的值。
	// 客户的单个或多个电子邮件地址。
	// 按照以下说明获取有效值：
	// 在散列处理之前去除每个电子邮件地址首尾空格。
	// 在散列处理之前将所有字符转换为小写。
	// 随后使用 SHA-256 算法对预处理后的电子邮件地址值进行散列处理。
	// 示例：
	// 原始值：" ALICE_abc@gmail.com "
	// 预处理后的值："alice_abc@gmail.com"（去除首尾空格且字符转换为小写）
	// SHA-256 散列处理后的值："848a771458438fc2ec420560d769fb9b9b86851ee338ec56517baabd79d3bb4f"
	Email []string `json:"email,omitempty"`
	// 网站事件、应用事件、线下事件和 CRM 事件均可指定本字段。
	// 需传入经 SHA-256 散列处理的值。
	// 客户的单个或多个电话号码。
	// 按照以下说明获取有效值：
	// 将电话号码预处理为 E.164 格式（例如，"+12133734253"）。建议使用 https://github.com/catamphetamine/libphonenumber-js 进行 E.164 解析。
	// 必须包含以 + 号为前缀的国家电话区号，无任何括号，且开头无0。例如，美国电话号码应以国家电话区号+1为前缀。
	// 除了国家代码中的+号外不包括任何空格、字母或符号。
	// 在预处理后对电话号码进行 SHA-256 散列处理。
	// 示例：
	// 原始值："(+1)2133734253"
	// 预处理为 E.164 格式："+12133734253"
	// SHA-256 散列处理后的值："9f7ec22d72092cd3c0b58726ed9c2d91b92e51a3f29837508fb2948bb22dd2fd"
	Phone []string `json:"phone,omitempty"`
	// ExternalID 仅网站事件和 CRM 事件可指定本字段。
	// 需传入经 SHA-256 散列处理的值。
	// 外部 ID 是广告主侧的唯一标识符，例如会员 ID、用户 ID 和外部 Cookie ID。若想了解 external_id 参数的详情，请查看事件 API 2.0 - 设置外部 ID。
	// 示例：
	// 原始值："user_123"
	// SHA-256 散列处理后的值："80fba0ae1c48e3978e43e4efc365e14e12ea0c830ba8ba5b9a2dafc7e3f2ab8b"
	ExternalID []string `json:"external_id,omitempty"`
	// Ttp 仅网站事件可指定本字段。
	// Cookie ID。
	// 若您还使用 Pixel SDK 并启用了 Cookie，Pixel SDK 会在 _ttp Cookie 中自动保存一个唯一标识符。_ttp的值用于将网站访客事件与 TikTok 广告进行匹配。您可以提取 _ttp的值并通过本字段传入。
	// 若想了解ttp参数的详情，请查看事件 API 2.0 - 发送 TikTok Cookie (_ttp)。
	Ttp string `json:"ttp,omitempty"`
	// IP 仅网站事件和应用事件可指定本字段。
	// 用户设备的未经散列处理的公共 IP 地址。
	// 为增加网站访客事件与 TikTok 广告的匹配率，建议同时传入 ip 和 user_agent。
	// 支持 IPv4 和 IPv6 地址。若为 IPv6 地址，支持传入完整格式或压缩格式。
	// 例如：
	// 完整格式：2001:0db8:1111:000a:00b0:0000:0000:0200
	// 压缩格式：2001:db8:1111:a:b0::200
	IP string `json:"ip,omitempty"`
	// UserAgent 仅网站事件和应用事件可指定本字段。
	// 用户设备的未经散列处理的用户代理。
	// 为增加网站访客事件与 TikTok 广告的匹配率，建议同时传入 ip 和 user_agent。
	UserAgent string `json:"user_agent,omitempty"`
	// IDFA iOS IDFA（广告主标识符），仅应用事件可传入本字段。
	// 对于 iOS 14 及以上版本，请确保 IDFA 数据收集遵循 Apple 的政策。
	// 允许传入未经 SHA-256 散列处理或经SHA-256 散列处理的 IDFA：
	// 对于未经 SHA-256 散列处理的 IDFA，所有字符必须为大写。
	// 对于经SHA-256 散列处理的 IDFA，所有字符必须为小写，且散列处理前需去掉首尾空格。
	IDFA string `json:"idfa,omitempty"`
	// IDFV iOS IDFV（应用开发商标识符），仅应用事件可传入本字段。
	IDFV string `json:"idfv,omitempty"`
	// GAID GAID（谷歌广告 ID），仅应用事件可传入本字段。
	// 允许传入未经 SHA-256 散列处理或经SHA-256 散列处理的 GAID。
	// 对于未经 SHA-256 散列处理或经SHA-256 散列处理的 GAID：
	// 所有字符必须为小写，且
	// 散列处理前需去掉首尾空格t。
	GAID string `json:"gaid,omitempty"`
	// Locale 仅网站事件和应用事件可指定本字段。
	// BCP 47 语言标识符。
	// 若想了解该标识符的详情，请查看IETF BCP 47 标准语言代码。
	Locale string `json:"locale,omitempty"`
	// AttStatus 仅应用事件可指定本字段。
	// 用户的应用追踪权限。即用户是否授予您的应用访问应用相关数据的权限，以追踪用户及其设备。这是 iOS 独有字段。
	// 枚举值：
	// AUTHORIZED：用户已授权访问可用于追踪用户及其设备的应用相关数据。
	// DENIED：用户尚未通过 Apple 的 AppTrackingTransparency 框架授权访问可用于追踪用户及其设备的应用相关数据。
	// 若设置为此值，TikTok 将不会将应用事件里的个人信息数据（例如用户电子邮件地址/手机号）用于用户追踪。
	// NOT_DETERMINED：用户尚未收到允许访问应用相关数据来追踪用户及其设备的授权请求。
	// RESTRICTED：追踪用户及其设备的应用相关数据的授权处于限制状态。
	// NOT_APPLICABLE：iOS 版本低于14 或正在使用 Android 设备。
	AttStatus string `json:"att_status,omitempty"`
}

// Properties 网站事件、应用事件、线下事件和 CRM 事件均可指定本字段。
// 有关产品和订单信息和其他附加信息，可用于优化广告表现。
// 与事件相关的订单和产品信息。提供这些信息可以帮助 TikTok 优化您的广告性能。
type Properties struct {
	// ContentIDs 商品或内容的唯一 ID 或多个 ID 的列表。
	// 示例：["987","654"]。
	// 推荐使用 sku id 或 item_group_id。若商品库中的商品设置了sku_id或item_group_id，则需保持一致。
	// 本参数在视频购物广告场景为必填。
	ContentIDs []string `json:"content_ids,omitempty"`
	// Contents 事件中对应的有商品信息的商品。
	Contents []Content `json:"contents,omitempty"`
	// ContentType 事件中内容的类型。
	// 枚举值：product（商品），product_group（商品组）。
	// 若contents参数中的content_id设置为sku_id，则本参数需设置为product。
	// 若contents参数中的content_id设置item_group_id，则本参数需设置为 product_group。
	ContentType string `json:"content_type,omitempty"`
	// Currency 推荐为营收相关事件传入本字段。
	// ISO 4217 币种代码。
	// 示例： EUR, USD, JPY。
	// 目前支持的币种清单： AED，ARS，AUD，BDT，BHD，BIF，BOB，BRL，CAD，CHF，CLP，CNY，COP，CRC，CZK，DKK，DZD，EGP，EUR，GBP，GTQ，HKD，HNL，HUF，IDR，ILS，INR，ISK，JPY，KES，KHR，KRW，KWD，KZT，MAD，MOP，MXN，MYR，NGN，NIO，NOK，NZD，OMR，PEN，PHP，PHP，PKR，PLN，PYG，QAR，RON，RUB，SAR，SEK，SGD，THB，TRY，TWD，UAH，USD，VES，VND，ZAR
	Currency string `json:"currency,omitempty"`
	// Value 推荐为营收相关事件传入本字段。
	// 订单或所售商品的价值。
	// 本字段值的格式需为整数或小数（例如10.00），无论地域、币种或其他因素。同时不允许包含币种符号，特殊字符，字母或逗号。
	// 注意：price（价格）是单件商品的价格，而value（价值）是指订单的总价。例如，若您有两件商品，每件商品的销售价格是 10 美元，那么 price 参数需要设置为 10.00，value参数需要设置为20.00。
	Value float64 `json:"value,omitempty"`
	// NumItems 商品数量。
	NumItems int `json:"num_items,omitempty"`
	// SearchString 用户输入的用于搜索的文本字符串。
	SearchString string `json:"search_string,omitempty"`
	// Description 商品或页面的描述。
	Description string `json:"description,omitempty"`
	// OrderID 订单 ID
	OrderID string `json:"order_id,omitempty"`
	// ShopID 商店 ID
	ShopID string `json:"shop_id,omitempty"`
}

// Content 事件中对应的有商品信息的商品。
type Content struct {
	// Price 商品价格。
	// 示例：10，10.5。
	// 注意：price（价格）是单件商品的价格，而value（价值）是指订单的总价。例如，若您有两件商品，每件商品的销售价格是 10 美元，那么 price 参数需要设置为 10.00，value参数需要设置为20.00。
	Price float64 `json:"price,omitempty"`
	// ContentID 商品或内容的唯一ID。
	// 推荐使用 sku_id 或item_group_id。
	// 若商品库中的商品设置了sku_id或item_group_id，则需保持一致。
	ContentID string `json:"content_id,omitempty"`
	// ContentCategory 页面或商品的类别。
	ContentCategory string `json:"content_category,omitempty"`
	// ContentName 页面或商品的名称t。
	ContentName string `json:"content_name,omitempty"`
	// Brand 商品的品牌名称t。
	Brand string `json:"brand,omitempty"`
}

// Page 有关网页的信息。
type Page struct {
	// URL 事件发生所在的浏览器 URL，例如客户端 JavaScript 中的 location.href的值。
	// 建议使用完整的 URL，包括所有的 URL 参数。
	URL string `json:"url,omitempty"`
	// Referrer 来源页面的 URL。 例如客户端 JavaScript 中的 document.referrer，或服务器端的Referer 请求头。
	// 建议使用完整的 URL，包括所有的 URL 参数。
	Referrer string `json:"referrer,omitempty"`
}

// App 应用事件需指定本字段。
// 有关应用的信息。
type App struct {
	// AppID 应用事件需指定本字段。
	// 移动应用 ID。
	// 对于 iOS 应用，请使用在 App Store URL 中的应用 ID。
	// 示例：Apple Store URL：https://apps.apple.com/us/app/angry-birds/id343200656
	// id = "343200656"
	// 对于在 Google Play 商店中的 Android 应用，请使用在 Google Play 商店 URL 中的应用 ID。
	// 示例：Google Play 商店 URL：https://play.google.com/store/apps/details?id=com.rovio.angrybird
	// id = "com.rovio.angrybirds"
	// 对于未在 Google Play 商店中的 Android 应用，请使用应用包名。
	AppID string `json:"app_id,omitempty"`
	// AppName 应用名称。
	AppName string `json:"app_name,omitempty"`
	// AppVersion 应用版本号。
	AppVersion string `json:"app_version,omitempty"`
}

// Ad 应用事件可指定本字段。
// 有关广告的信息。
// 使用ad参数与事件 API 2.0 分享有关广告主的应用的信息。ad参数应仅在报告应用事件（event_source = app）时使用。
type Ad struct {
	// Callback 回调信息，用于帮助事件归因。
	Caltlback string `json:"callbeack,omitempty"`
	// CampaignID 推广系列 IDt。
	CampaignID string `json:"campaign_id,omitempty"`
	// AdID 广告组 ID
	AdID string `json:"ad_id,omitempty"`
	// CreativeID 广告 ID
	CreativeID string `json:"creative_id,omitempty"`
	// IsRetargeting 用户是否是重定向（再营销）客户。
	IsRetargeting *bool `json:"is_retargeting,omitempty"`
	// Attributed 事件是否已归因
	Attributed *bool `json:"attributed,omitempty"`
	// AttributeType 归因类型
	AttributeType string `json:"attribute_type,omitempty"`
	// AttributeProvider 归因服务提供商
	AttributeProvider string `json:"attribute_provider,omitempty"`
}

// Lead CRM 事件需指定本字段。
// 有关线索的信息
type Lead struct {
	// LeadID 对于 CRM 事件必填。
	// TikTok 线索的 ID。
	// 从 TikTok 导出的线索均有对应的lead_id。
	// 若想获取线索 ID，可使用 /page/lead/task/ 和 /page/lead/task/download/ 接口下载线索，或从 TikTok 广告管理平台中的第一方 CRM 平台 Leads Center 导出线索。
	LeadID string `json:"lead_id,omitempty"`
	// LeadEventSource 线索事件来源。
	// 您可将本字段设置为您的 CRM 系统的名称，例如 HubSpot 或 Salesforce。
	LeadEventSource string `json:"lead_event_source,omitempty"`
}
