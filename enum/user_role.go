package enum

// UserRole 用户（成员）在商务中心中的角色。
type UserRole string

const (
	// UserRole_ADMIN：管理员。管理员可以访问商务中心内的所有功能。
	UserRole_ADMIN UserRole = "ADMIN"
	// UserRole_STANDARD：标准。标准成员只能对分配给他们的账户进行操作
	UserRole_STANDARD UserRole = "STANDARD"
)
