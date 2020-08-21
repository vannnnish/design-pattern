package main


type OperatorBase struct {
	left  int
	right int
}

// 赋值
func (op *OperatorBase) SetLeft(left int) {
	op.left = left
}

func (op *OperatorBase) SetRight(right int) {
	op.right = right
}
