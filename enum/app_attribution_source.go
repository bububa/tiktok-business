package enum

type AppAttributionSource string

const (
	// AppAttributionSource_MMP: MMP (mobile measurement partner) attribution. Signals from your selected MMP will be used for attribution.
	AppAttributionSource_MMP AppAttributionSource = "MMP"
	// AppAttributionSource_SAN: SAN (self-attribution network) attribution. Signals from TikTok SAN will be used for attribution. The exact value you can specify depends on your app (app_id) setup. To retrieve the available attribution and data source settings for your app, use /tool/available/attribution_source/.
	AppAttributionSource_SAN AppAttributionSource = "SAN"
)
