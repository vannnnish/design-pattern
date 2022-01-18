package main

import "fmt"

/*
把一个类的接口变成客户端所期待的另一种接口，
从而使原本因接口原因不匹配而无法工作的两个类能够一起工作

最简单的用法比如: A方法里面接受的是B对象， 这时候也想使用C对象的部分方法， 这时候，就可以定义一个C的适配器，实现B的方法。然后去

一般用在版本迭代中
*/
// 适配目标接口
type Target interface {
	Request() string
}

type adapter struct {
	Adaptee
}

func NewAdapter(adapee Adaptee)Target{
	return &adapter{
		adapee,
	}
}

// 接口包装
func (adap *adapter) Request() string {
	return adap.SpecificRequest()
}

// 被适配
type Adaptee interface {
	SpecificRequest() string
}

// 载体，被适配的目标类
type adapeeImp struct {
}

// 工厂函数
func NewAdaptee() Adaptee {
	return &adapeeImp{}
}

func (ada *adapeeImp) SpecificRequest() string {
	return "SpecificRequest"
}


func main(){
	adapee:=NewAdaptee()// adapter
	target:=NewAdapter(adapee)
	fmt.Println(target.Request())
}