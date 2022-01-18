package main

/*
	状态模式的使用： 现将你所有的状态 和 行为都定义出来。
*/
func main(){
	ctx:=NewDayContext()
	todayAndNext:= func() {
		ctx.Today()
		ctx.Next()
	}
	for i:=0;i<5;i++{
		todayAndNext()
	}
}