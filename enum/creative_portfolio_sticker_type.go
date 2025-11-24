package enum

// CreativePortfolioStickerType 贴纸类型
type CreativePortfolioStickerType string

const (
	// CreativePortfolioStickerType_COUNTDOWN：无提醒的倒计时贴纸。
	CreativePortfolioStickerType_COUNTDOWN CreativePortfolioStickerType = "COUNTDOWN"
	// CreativePortfolioStickerType_REMINDER_COUNTDOWN： 含非直播活动提醒的倒计时贴纸。
	CreativePortfolioStickerType_REMINDER_COUNTDOWN CreativePortfolioStickerType = "REMINDER_COUNTDOWN"
	// CreativePortfolioStickerType_LIVE_REMINDER_COUNTDOWN：含直播活动提醒的倒计时贴纸。
	CreativePortfolioStickerType_LIVE_REMINDER_COUNTDOWN CreativePortfolioStickerType = "LIVE_REMINDER_COUNTDOWN"
	// CreativePortfolioStickerType_GIFTCODE：礼品码贴纸，用于通过提供虚拟礼品或优惠券兑换码来吸引新客户和促进现有客户。
	CreativePortfolioStickerType_GIFTCODE CreativePortfolioStickerType = "GIFTCODE"
)
