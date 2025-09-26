package enum

// SubscribeEntity 订阅对象。
type SubscribeEntity string

const (
	// SubscribeEntity_AD_ACCOUNT_SUSPENSION：广告账户的暂时停用状态。
	SubscribeEntity_AD_ACCOUNT_SUSPENSION SubscribeEntity = "AD_ACCOUNT_SUSPENSION"
	// SubscribeEntity_LEAD：线索。
	SubscribeEntity_LEAD SubscribeEntity = "LEAD"
	// SubscribeEntity_AD_GROUP：广告组的审核状态。
	SubscribeEntity_AD_GROUP SubscribeEntity = "AD_GROUP"
	// SubscribeEntity_AD：广告的审核状态。
	SubscribeEntity_AD SubscribeEntity = "AD"
	// SubscribeEntity_TCM_SPARK_ADS：上传至某个 TCM 工作流程 2.0 订单的视频的 Spark Ads 授权状态。
	SubscribeEntity_TCM_SPARK_ADS SubscribeEntity = "TCM_SPARK_ADS"
	// SubscribeEntity_CREATIVE_FATIGUE：单个广告、广告组中所有广告或广告账户下所有广告的疲劳状态。
	SubscribeEntity_CREATIVE_FATIGUE SubscribeEntity = "CREATIVE_FATIGUE"
)
