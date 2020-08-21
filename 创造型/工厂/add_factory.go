package main



// 操作的抽象
type PlusOperatorFactory struct{}

// 操作类中包含操作数
type PlusOperator struct {
	*OperatorBase
}

// 实际运行
func (o PlusOperator) Result() int {
	return o.right + o.left
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{&OperatorBase{}}
}