package enum

// TikTokItemStatus TikTok帖子状态
type TikTokItemStatus string

const (
	// TikTokItemStatus_USER_DELETE 已由发帖用户删除。
	TikTokItemStatus_USER_DELETE TikTokItemStatus = "USER_DELETE"
	// TikTokItemStatus_AUDIT_DELETE 未通过审核，已由审核人员删除。
	TikTokItemStatus_AUDIT_DELETE TikTokItemStatus = "AUDIT_DELETE"
	// TikTokItemStatus_HESITATE_RECOMMEND 公开可见。
	TikTokItemStatus_HESITATE_RECOMMEND TikTokItemStatus = "HESITATE_RECOMMEND"
	// TikTokItemStatus_ONLY_AUTHOR_SEE 仅发帖用户可见。
	TikTokItemStatus_ONLY_AUTHOR_SEE TikTokItemStatus = "ONLY_AUTHOR_SEE"
	// TikTokItemStatus_ONLY_USER_PASS 发布后由审核人员删除。
	TikTokItemStatus_ONLY_USER_PASS TikTokItemStatus = "ONLY_USER_PASS"
	// TikTokItemStatus_ONLY_FRIEND_SEE 仅好友可见。
	TikTokItemStatus_ONLY_FRIEND_SEE TikTokItemStatus = "ONLY_FRIEND_SEE"
	// TikTokItemStatus_REVIEW_PENDING 审核中。
	TikTokItemStatus_REVIEW_PENDING TikTokItemStatus = "REVIEW_PENDING"
	// TikTokItemStatus_ITEM_STATUS_HESITATE_RECOMMEND：对所有人可见。
	TikTokItemStatus_ITEM_STATUS_HESITATE_RECOMMEND TikTokItemStatus = "ITEM_STATUS_HESITATE_RECOMMEND"
	// TikTokItemStatus_STATUS_ONLY_FRIEND_SEE：仅对好友可见。
	TikTokItemStatus_STATUS_ONLY_FRIEND_SEE TikTokItemStatus = "STATUS_ONLY_FRIEND_SEE"
	// TikTokItemStatus_ITEM_STATUS_ONLY_AUTHOR_SEE：仅自己可见。
	TikTokItemStatus_ITEM_STATUS_ONLY_AUTHOR_SEE TikTokItemStatus = "ITEM_STATUS_ONLY_AUTHOR_SEE"
)
