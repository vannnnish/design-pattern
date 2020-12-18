package responsibility

type Manager interface {
	HaveRight(money int) bool                     // 判断权限
	HandleFeeRequest(name string, money int) bool // 向后传递
}

