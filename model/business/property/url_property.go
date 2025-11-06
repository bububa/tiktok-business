package property

// URLProperty URL 资源的信息
type URLProperty struct {
	// PropertyType URL 资源的类型。
	// 枚举值：
	// 1：网域。
	// 2：网址前缀。
	PropertyType int `json:"property_type,omitempty"`
	// URL 您想要添加的自有 URL。需根据您指定的 property_type 传入对应 URL。
	// 若 property_type 设置为 1，则需要指定基本域（例如，example.com）或子域（例如，subdomain.example.com）。
	// 若 property_type 设置为 2，则需要指定由以下内容组成的网址前缀：https:// + 主机名（必须是域名）+ 路径 + /，并且网址前缀不可重定向至其他 URL。例如，在网址前缀 https://www.example.com/folder/ 中，主机名是域名 www.example.com，路径是 /folder。
	URL string `json:"url,omitempty"`
	// PropertyStatus URL 资源的验证状态。
	// 枚举值：
	// 1：已验证。
	// 2：由于验证尝试失败而未验证。您需要重新验证 URL 资源。
	PropertyStatus int `json:"property_status,omitempty"`
	// Signature 在所有权验证过程中要使用的签名字符串
	Signature string `json:"signature,omitempty"`
	// FileName 在所有权验证过程中要使用的签名文件名
	FileName string `json:"file_name,omitempty"`
}
