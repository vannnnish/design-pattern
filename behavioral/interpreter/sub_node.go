package main

type SubNode struct {
	left, right Node
}

func (s *SubNode) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}
