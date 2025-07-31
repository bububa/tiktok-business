package enum

// AudienceType 应用重定向受众类型
type AudienceType string

const (
	// AudienceType_CUSTOM_AUDIENCE 自定义受众。不支持TRAFFIC推广目标
	AudienceType_CUSTOM_AUDIENCE AudienceType = "CUSTOM_AUDIENCE"
	// AudienceType_EXCLUDE_INTERACT_USERS 排除曾与您的应用互动的用户。
	AudienceType_EXCLUDE_INTERACT_USERS AudienceType = "EXCLUDE_INTERACT_USERS"
	// AudienceType_INCLUDE_INTERACT_USERS 定向曾与您的应用互动的用户
	AudienceType_INCLUDE_INTERACT_USERS AudienceType = "INCLUDE_INTERACT_USERS"
	// AudienceType_INCLUDE_SPECIFIC_EVENTS 定向满足特定规则的受众
	AudienceType_INCLUDE_SPECIFIC_EVENTS AudienceType = "INCLUDE_SPECIFIC_EVENTS"
	// AudienceType_EXCLUDE_SPECIFIC_EVENTS 排除满足特定规则的受众。
	AudienceType_EXCLUDE_SPECIFIC_EVENTS AudienceType = "EXCLUDE_SPECIFIC_EVENTS"
	// AudienceType_NEW_CUSTOM_AUDIENCE 新的自定义受众类型。新建广告组目前唯一支持的应用重定向受众类型
	AudienceType_NEW_CUSTOM_AUDIENCE AudienceType = "NEW_CUSTOM_AUDIENCE"
)
