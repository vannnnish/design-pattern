package main

import "fmt"

type GeneralManger struct {
}

func NewGeneraProjectManager() *GeneralManger {
	return &GeneralManger{}
}

func NewGeneraManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManger{},
	}
}

func (gm *GeneralManger) HaveRight(money int) bool {
	return true
}

func (gm *GeneralManger) HandleFeeRequest(name string, money int) bool {
	if name == "weishangyin" {
		fmt.Printf("GeneralManager 授权 %s %d\n", name, money)
		return true
	}
	fmt.Printf("GeneralManager not 授权 %s %d\n", name, money)
	return false
}
