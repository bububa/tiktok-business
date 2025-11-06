package enum

// CreatorBadge 该创作者目前拥有的徽章。
type CreatorBadge string

const (
	// CreatorBadge_RESPONSIVE ：响应迅速。创作者在48小时内回复品牌商。
	CreatorBadge_RESPONSIVE CreatorBadge = "RESPONSIVE"
	// CreatorBadge_EXPERIENCED ：经验丰富。创作者已完成第一个推广系列。
	CreatorBadge_EXPERIENCED CreatorBadge = "EXPERIENCED"
	// CreatorBadge_PROFILE_COMPLETE ：已完善个人主页。创作者已完整填写个人资料。
	CreatorBadge_PROFILE_COMPLETE CreatorBadge = "PROFILE_COMPLETE"
)
