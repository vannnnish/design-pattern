package impl


/*
	这一层可要，可不要，
	要的话，1可以省略SetLeft，SetRight这一些重复代码的编写，
*/
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
