package main

import "fmt"

// 把业务交给第三方去处理
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

func main(){
	var sub Subject
	sub=&Proxy{money: -100}
	fmt.Println(sub.Do())
}