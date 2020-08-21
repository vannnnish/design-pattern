package main

// 工厂模式，就是对固有的属性，进行固化，新增功能时，不需要修改到之前的代码
type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}



