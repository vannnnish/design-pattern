package main

import "fmt"

// 命令的接收者,也就是实际要执行的逻辑，也就是被调用者。
type TV struct{}

func (tv *TV) Open() {
	fmt.Println("打开电视")
}
func (tv *TV) Close() {
	fmt.Println("关闭电视")
}
func (tv *TV) Change() {
	fmt.Println("换台")
}
