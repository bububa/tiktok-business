package enum

// TTItemStatus TikTok帖子状态
type TTItemStatus string

const (
	// TTItemStatus_USER_DELETE 已由发帖用户删除。
	TTItemStatus_USER_DELETE TTItemStatus = "USER_DELETE"
	// TTItemStatus_AUDIT_DELETE 未通过审核，已由审核人员删除。
	TTItemStatus_AUDIT_DELETE TTItemStatus = "AUDIT_DELETE"
	// TTItemStatus_HESITATE_RECOMMEND 公开可见。
	TTItemStatus_HESITATE_RECOMMEND TTItemStatus = "HESITATE_RECOMMEND"
	// TTItemStatus_ONLY_AUTHOR_SEE 仅发帖用户可见。
	TTItemStatus_ONLY_AUTHOR_SEE TTItemStatus = "ONLY_AUTHOR_SEE"
	// TTItemStatus_ONLY_USER_PASS 发布后由审核人员删除。
	TTItemStatus_ONLY_USER_PASS TTItemStatus = "ONLY_USER_PASS"
	// TTItemStatus_ONLY_FRIEND_SEE 仅好友可见。
	TTItemStatus_ONLY_FRIEND_SEE TTItemStatus = "ONLY_FRIEND_SEE"
	// TTItemStatus_REVIEW_PENDING 审核中。
	TTItemStatus_REVIEW_PENDING TTItemStatus = "REVIEW_PENDING"
)
