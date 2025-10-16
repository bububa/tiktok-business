package feed

import (
	"github.com/bububa/tiktok-business/enum"
	"github.com/bububa/tiktok-business/model"
)

// Feed 更新源
type Feed struct {
	// FeedID 更新源 ID
	FeedID string `json:"feed_id,omitempty"`
	// FeedName 更新源名称
	FeedName string `json:"feed_name,omitempty"`
	// Status 更新源的定时更新的状态。
	// 枚举值：
	// ON: 更新源有开启（"ON"）状态的定时更新。
	// OFF: 更新源无开启状态的定时更新。此时可能更新源未设置定时更新，或更新源的定时更新处于关闭（"OFF"）状态。
	Status enum.OnOff `json:"status,omitempty"`
	// LastUpdateParam 最后一次更新信息
	LastUpdateParam *ScheduleParam `json:"last_update_param,omitempty"`
	// NextUpdateTime 下次更新的日期和时间（UTC + 0 时间）。示例："2021-05-23 16:33:23"。
	NextUpdateTime model.DateTime `json:"next_update_time,omitzero"`
	// NumberOfProducts 更新源中的商品数量
	NumberOfProducts int `json:"number_of_products,omitempty"`
}

// ScheduleParam 定时更新的设置
type ScheduleParam struct {
	// Source 更新源
	Source *Source `json:"source,omitempty"`
	// URI 更新源地址，即更新源文件所在的 URL
	URI string `json:"uri,omitempty"`
	// IntervalType 定时更新间隔单位。
	// 枚举值: HOURLY（小时），DAILY（日），MONTHLY（月）。
	IntervalType enum.IntervalType `json:"interval_type,omitempty"`
	// IntervalCount 两次相邻的更新之间的间隔。
	// 若 interval_type 为 HOURLY，取值范围为 1-23。
	// 若 interval_type 为 DAILY，取值范围为 1-30。
	// 若 interval_type 为 MONTHLY，取值范围是 1-12。
	IntervalCount int `json:"interval_count,omitempty"`
	// Timezone 更新计划使用的时区。
	// 示例：Africa/Accra。
	// 时区列表可见 附录-时区信息。
	Timezone string `json:"timezone,omitempty"`
	// DayOfMonth 对于按月执行的定时更新，执行更新的月份的第几日
	DayOfMonth int `json:"day_of_month,omitempty"`
	// Hour 获取更新的小时时间。
	// 若定时更新是按小时执行，本字段代表首次执行更新的小时时间。
	// 若定时更新是按日或月执行，本字段代表执行定时更新的小时时间。
	Hour int `json:"hour,omitempty"`
	// Minute 获取更新的分钟时间。
	// 若定时更新是按小时执行，本字段代表首次更新的分钟时间。
	// 若定时更新是按日或月执行，本字段代表执行定时更新的分钟时间。
	Minute int `json:"minute,omitempty"`
}

// Source 更新源
type Source struct {
	// URI 更新源地址，即更新源文件所在的 URL
	// 传入interval_type 或 timezone时必填。
	// 更新源地址。请添加更新源文件所在的 URL。应以 http、https、ftp 或 sftp 开头。文件最大为 8 GB，可采用 CSV、TSV 或 XML (RSS/ATOM) 格式。
	URI string `json:"uri,omitempty"`
	// Username 如果更新源有密码保护，您需要用本字段传入用户名
	Username string `json:"username,omitempty"`
	// Password 如果更新源有密码保护，您需要用本字段传入密码
	Password string `json:"password,omitempty"`
}
