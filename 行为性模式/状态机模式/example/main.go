package main

import "fmt"

func main() {
	// 给上下文传了一个 关闭的状态
	context := &StateContext{&CloseState{}}
	// 对关闭进行打开，然后对其设置状态
	context.ToOpen(context)
	fmt.Println()
	context.ToPlay(context)
	context.ToClose(context)

	// 获取当前公司状态, 审核接口
	ctxFactory := NewStateFactory("open")
	//
	ctxFactory.ToPlay(ctxFactory)
	// 操作是关闭
/*	if ctxFactory.ToClose() {
		// TODO： 关闭成功
	} else {
		// TODO：没有关闭的这个操作
	}*/

}
