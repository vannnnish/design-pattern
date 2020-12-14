package impl

import "fmt"

type AddComponent struct {
	num int
}

func (mul *AddComponent) Calc() int {
	// 执行乘法逻辑
	fmt.Println("执行乘法逻辑")
	return mul.num
}
