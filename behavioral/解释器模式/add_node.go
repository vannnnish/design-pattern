package main

type AddNode struct {
	left, right Node
}

func (a *AddNode) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}
