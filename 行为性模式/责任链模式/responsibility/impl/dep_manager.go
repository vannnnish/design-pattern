package impl

import (
	"fmt"
)

type DepManger struct {
}

func NewDepManager() *DepManger {
	return &DepManger{}
}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManger{},
	}
}

func (dm *DepManger) HaveRight(money int) bool {
	return money < 5000
}

func (dm *DepManger) HandleFeeRequest(name string, money int) bool {
	if name == "lixiang" {
		fmt.Printf("DepManager 授权 %s %d\n", name, money)
		return true
	}
	fmt.Printf("DepManager not 授权 %s %d\n", name, money)
	return false
}
