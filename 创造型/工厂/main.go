package main

import (
	"fmt"
	"pattern/创造型/工厂/operator"
	"pattern/创造型/工厂/operator/impl"
)

/*
	这个模式遵循的原则:
		满足了单一职责，add_factory 这个文件，就是支付者加的操作
		满足开闭原则: 无需修改原先调用的代码，就可以拓展新的操作类型
*/

func CreateOperatorFactory(operate string) operator.Operator {
	switch operate {
	case "add":
		return impl.NewAddFactory()
	case "sub":
		return impl.NewSubFactory()
	default:
		return nil
	}
}

func main() {
	adder := CreateOperatorFactory("add")
	adder.SetLeft(20)
	adder.SetRight(10)
	fmt.Println(adder.Result())

	subscriber := CreateOperatorFactory("sub")
	subscriber.SetLeft(20)
	subscriber.SetRight(10)
	fmt.Println(subscriber.Result())

}
