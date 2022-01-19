package main

import (
	"pattern/行为性模式/责任链模式/responsibility"
	"pattern/行为性模式/责任链模式/responsibility/impl"
)

/*
	责任链模式， 一般用在 权限验证处理这一块， 减少if 之类语句的使用。

	责任链模式的实用:  接口定义好， 权限判断条件， 以及相应的处理。
	对应的责任执行对象去实现相应的  实例中: dep genera project里面就是责任执行对象

	责任链对象里面定义好，处理流程。 request_chain就是责任执行流程
*/
func main() {
	c1 := impl.NewProjectManagerChain()
	c2 := impl.NewDepManagerChain()
	c3 := impl.NewGeneraManagerChain()
	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)
	var c responsibility.Manager = c1
	c.HandleFeeRequest("hupeng", 500)
	c.HandleFeeRequest("hupeng", 1500)
	c.HandleFeeRequest("lixiang", 4500)
	c.HandleFeeRequest("lixiang", 11500)
}
