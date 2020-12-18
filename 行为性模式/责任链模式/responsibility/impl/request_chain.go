package impl

import "pattern/行为性模式/责任链模式/responsibility"

type RequestChain struct {
	responsibility.Manager
	successor *RequestChain // 处理成功
}

// 判断责任链，在哪里结束
func (rc *RequestChain) SetSuccessor(endrc *RequestChain) {
	rc.successor = endrc
}

// 权利做判断
func (rc *RequestChain) HaveRight(money int) bool {
	return true
}

// 向后传递， 有处理权利则直接处理, 没有权利则向后传递
func (rc *RequestChain) HandleFeeRequest(name string, money int) bool {
	if rc.Manager.HaveRight(money) {
		return rc.Manager.HandleFeeRequest(name, money)
	}
	if rc.successor != nil {
		return rc.successor.HandleFeeRequest(name, money)
	}
	return false

}
