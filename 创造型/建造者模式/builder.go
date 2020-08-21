package main

import "fmt"

/*
	如果启动的时候，有一个规定的顺序， 比如先要初始化conf
		然后初始化mysql 再初始化 redis 这个时候 可以采用建造者


	建造者跟工厂的区别就是， 工厂模式是用户·输入不变，通过修改执行方法，以获取不同的输出。
	建造者模式就是， 修改输入的对象， 执行方法不变， 以获取不同的输出
*/

type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) MakeData() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

func main() {
	builder:=&StringBuilder{}
	//builder := &IntBuilder{}
	dict := NewDirector(builder)
	dict.MakeData()
	fmt.Println(builder.GetResult())
}
