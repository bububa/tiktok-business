package enum

// PrimaryStatus 一级状态
type PrimaryStatus string

const (
	// STATUS_DELETE 已删除
	STATUS_DELETE PrimaryStatus = "STATUS_DELETE"
	// STATUS_DELIVERY_OK 投放中
	STATUS_DELIVERY_OK PrimaryStatus = "STATUS_DELIVERY_OK"
	// STATUS_DISABLE 暂停投放
	STATUS_DISABLE PrimaryStatus = "STATUS_DISABLE"
	// STATUS_NOT_DELIVERY 不在投放中
	STATUS_NOT_DELIVERY PrimaryStatus = "STATUS_NOT_DELIVERY"
	// STATUS_TIME_DONE 已完成（推广系列无此状态）
	STATUS_TIME_DONE PrimaryStatus = "STATUS_TIME_DONE"
	// STATUS_RF_CLOSED 覆盖和频次广告组已关闭（仅覆盖和频次广告组和广告有此状态）
	STATUS_RF_CLOSED PrimaryStatus = "STATUS_RF_CLOSED"
	// STATUS_FROZEN 覆盖和频次广告组冻结中（仅覆盖和频次广告有此状态）
	STATUS_FROZEN PrimaryStatus = "STATUS_FROZEN"
	// STATUS_ALL 所有状态（包含已删除）
	STATUS_ALL PrimaryStatus = "STATUS_ALL"
	// STATUS_NOT_DELETE 所有状态（不包含已删除）。该枚举值为基础、受众分析或DSA类型的同步报表中的默认筛选条件。
	STATUS_NOT_DELETE PrimaryStatus = "STATUS_NOT_DELETE"
)
