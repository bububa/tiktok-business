package enum

// IDSchema 传送受众时所用的 ID 类型。
type IDSchema string

const (
	// IDSchema_IDFA_MD5: 以 MD5 加密 iOS IDFA。
	IDSchema_IDFA_MD5 IDSchema = "IDFA_MD5"
	// IDSchema_AAID_MD5: 以 MD5 加密 Android AAID 或 GAID。
	IDSchema_AAID_MD5 IDSchema = "AAID_MD5"
	// IDSchema_IDFA_SHA256: 以 SHA256 加密 iOS IDFA。
	IDSchema_IDFA_SHA256 IDSchema = "IDFA_SHA256"
	// IDSchema_AAID_SHA256: 以 SHA256 加密 Android AAID 或 GAID。
	IDSchema_AAID_SHA256 IDSchema = "AAID_SHA256"
	// IDSchema_EMAIL_SHA256: 以 SHA256 加密电子邮件地址。请参阅流式 API 相关提示 - 电子邮件地址标准化部分。
	IDSchema_EMAIL_SHA256 IDSchema = "EMAIL_SHA256"
	// IDSchema_PHONE_SHA256: 以 SHA256 加密 E.164格式 的电话号码。举例：+1231234567。
	IDSchema_PHONE_SHA256 IDSchema = "PHONE_SHA256"
)
