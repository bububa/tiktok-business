package enum

// AppDataSource Data source for your app.
// The attribution signals you select will be used for reporting.
type AppDataSource string

const (
	// AppDataSource_MMP: MMP signals. Signals from your MMP will be used for reporting.
	AppDataSource_MMP AppDataSource = "MMP"
	// AppDataSource_EVENT_SDK: App Events SDK signals. Signals from TikTok App Events SDK will be used for reporting.
	AppDataSource_EVENT_SDK AppDataSource = "EVENT_SDK"
	// AppDataSource_EVENT_API: Events API signals. Signals from TikTok Events API will be used for reporting.
	AppDataSource_EVENT_API AppDataSource = "EVENT_API"
)
