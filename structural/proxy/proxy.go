package main

import "fmt"

// 把业务交给第三方去处理
/*
	代理模式和策略模式区别就是，
	proxy，在特定的条件下，会由代理人直接去执行相应的逻辑

	而策略模式，从始至终都是传入的对象去执行的逻辑。并且，策略模式从对象创建出来的那一刻就已经确定好了，
	谁将会去执行这个逻辑。
*/
type Subject interface {
	Do() string
}

type RealSubject struct {
}

func (db *RealSubject) Do() string {
	return "以太坊执行智能合约"
}

type Proxy struct {
	real  RealSubject
	money int
}

func (p *Proxy) Do() string {
	if p.money > 0 {
		return p.real.Do()
	} else {
		return "请充值"
	}
}

/*
	 proxy 主要应用与核心业务逻辑，与此业务逻辑的分离。
	比如例子中: RealSubject 里面Do执行的是核心逻辑，
	那么像余额判断，输出参数的校验之类的就可以放在代理里面进行校验。
	这样就可以做到 单一职责， 开闭原则
*/
func main(){
	var sub Subject
	sub=&Proxy{money: -100}
	fmt.Println(sub.Do())
}