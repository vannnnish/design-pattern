package main

import "fmt"

type ProjectManger struct {
}

func NewProjectManager() *ProjectManger {
	return &ProjectManger{}
}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManger{},
	}
}

func (pm *ProjectManger) HaveRight(money int) bool {
	return money < 1000
}

func (pm *ProjectManger) HandleFeeRequest(name string, money int) bool {
	if name == "hupeng" {
		fmt.Printf("ProjectManager 授权 %s %d\n", name, money)
		return true
	}
	fmt.Printf("ProjectManager not 授权 %s %d\n", name, money)
	return false
}
