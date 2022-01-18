package main

import (
	"fmt"
	"pattern/structural/decorator/operator/impl"
)

// 这部分就是装饰的部分， 如果要对原函数功能进行扩充，那么不需要修改原函数，而是，新增一个文件，在这里进行修改
/*
	装饰器的核心就是， 你的本体是估计不变的，示例里面， base里面的逻辑就是我们的核心逻辑, 如果我们需要一些前置逻辑，
	那么就创建一个新的文件夹，里面实现你需要的方法。

	
*/
func main() {
	fmt.Println(impl.NewWarpComponent(&impl.AddComponent{}).Calc())
	fmt.Println(impl.NewWarpComponent(&impl.SubComponent{}).Calc())
}
