package main

/*
	感觉不是一个很实用的命令模式，
	因为， 如果对receiver 进行拓展，那么，invoker也是要修改的，
	接着client也需要相应的修改。 感觉这样就涉及到很多改动了。

	并且，代码逻辑看起来很不清晰！！！
*/
func main() {
	// 创建接收者
	tvR := NewTVRemote()
	tvR.Open()
	tvR.Change()
	tvR.Close()
}
