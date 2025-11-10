package enum

// MaterialAction 您想要对音乐进行的操作
type MaterialAction string

const (
	// MaterialAction_ADD_TO_LIKED：将音乐添加至收藏列表。
	MaterialAction_ADD_TO_LIKED MaterialAction = "ADD_TO_LIKED"
	// MaterialAction_ADD_TO_HISTORY：将音乐添加至历史列表。
	MaterialAction_ADD_TO_HISTORY MaterialAction = "ADD_TO_HISTORY"
	// MaterialAction_REMOVE_FROM_LIKED：将音乐从收藏列表中移除。
	MaterialAction_REMOVE_FROM_LIKED MaterialAction = "REMOVE_FROM_LIKED"
)
