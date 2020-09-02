package main

import "fmt"

type LeaderStrategy interface {
	Way(ctx *LeaderContext)
}

type LeaderContext struct {
	Name       string
	Age        int
	wayStategy LeaderStrategy
}

func NewLeaderContext(name string, age int, ls LeaderStrategy) *LeaderContext {
	return &LeaderContext{
		name, age, ls,
	}
}

func (l *LeaderContext) Way() {
	l.wayStategy.Way(l)
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
