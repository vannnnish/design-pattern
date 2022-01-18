package main

import "fmt"

type LeaderStrategy interface {
	Way(ctx *LeaderContext)
}

// 上下文封装角色,实际接受请求并将请求委托给实际的算法实现类处理
/*
	策略模式的优点，将所有的抽象放入一组策略接口中，方便统一管理和实现

	缺点:策略模式每种策略都是单独类，策略很多时策略实现类数量也很可观。

	客户端初始化 Context 的时候需要指定策略类，这样就要求客户端要熟悉各个策略，对调用方要求较高。

	感觉并不是很实用，不如代理模式好用。 就像是一个工厂 + builder
*/

/*
		代理模式和策略模式区别就是，
	proxy，在特定的条件下，会由代理人直接去执行相应的逻辑

	而策略模式，从始至终都是传入的对象去执行的逻辑。并且，策略模式从对象创建出来的那一刻就已经确定好了，
	谁将会去执行这个逻辑。
*/

/*
	使用场景:
	需要自由切换算法场景，
	需要屏蔽算法实现场景的情况
*/
type LeaderContext struct {
	Name       string
	Age        int
	wayStrategy LeaderStrategy
}

func NewLeaderContext(name string, age int, ls LeaderStrategy) *LeaderContext {
	return &LeaderContext{
		name, age, ls,
	}
}

func (l *LeaderContext) Way() {
	l.wayStrategy.Way(l)
}

type Manager struct {
}

func (*Manager) Way(ctx *LeaderContext) {
	fmt.Println("跟管理员说项目进度OK", ctx.Name)
}

type Women struct{}

func (*Women) Way(ctx *LeaderContext) {
	fmt.Println("遇到BOSS 说 早上好", ctx.Name)
}
