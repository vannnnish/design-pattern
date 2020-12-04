package imp

// 操作类中包含操作数
type PlusOperator struct {
	*OperatorBase
}

// 实际运行
func (o PlusOperator) Result() int {
	return o.right + o.left
}

func NewAddFactory() *PlusOperator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}
