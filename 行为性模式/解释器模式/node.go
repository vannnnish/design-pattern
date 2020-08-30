package main

// interpreter

// 节点返回一个数据
type Node interface {
	Interpret() int
}

type ValNode struct {
	val int
}

func (val *ValNode) Interpret() int {
	return val.val
}
