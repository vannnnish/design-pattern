package impl

import "fmt"

type SubComponent struct {
	num int
}

func (sub *SubComponent) Calc() int {
	// 执行除法逻辑
	fmt.Println("执行减法逻辑")
	return sub.num
}
