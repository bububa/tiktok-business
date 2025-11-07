package enum

// ROIProtectionCompensationStatus 该推广系列的 ROI（投资回报率）保障状态。
type ROIProtectionCompensationStatus string

const (
	// ROIProtectionCompensationStatus_IN_EFFECT：符合 ROI 保障资格。如果您的广告在 24 小时内促成 20 次以上转化，但 ROI 结果低于目标的 90%，此推广系列有资格获得广告费用赠款。为确保推广系列持续符合 ROI 保障资格，请勿修改 ROI、暂停广告投放或启用最大投放量优化模式。推广系列资格每 24 小时重置一次。
	ROIProtectionCompensationStatus_IN_EFFECT ROIProtectionCompensationStatus = "IN_EFFECT"
	// ROIProtectionCompensationStatus_NOT_ELIGIBLE：不符合 ROI 保障资格。如果因为修改了 ROI、暂停了推广系列、启用了最大投放量优化模式，或者您的店铺、广告账号、直播或商品存在问题，导致没有达成目标 ROI，那么此推广系列不符合广告费用赠款资格。推广系列资格每 24 小时重置一次。
	ROIProtectionCompensationStatus_NOT_ELIGIBLE ROIProtectionCompensationStatus = "NOT_ELIGIBLE"
)
