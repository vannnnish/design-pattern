package main

// 这部分就是装饰的部分， 如果要对原函数功能进行扩充，那么不需要修改原函数，而是，新增一个文件，在这里进行修改
type AddComponent struct {
	Component
	num int
}

func WarpAddComponent(c Component,num int)Component{
	return &AddComponent{c,num}
}


func (mul *AddComponent)Calc()int{
	return mul.Component.Calc()+mul.num
}