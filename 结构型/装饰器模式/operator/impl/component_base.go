package impl

import "pattern/结构型/装饰器模式/operator"

type ConcreateComponent struct {
	component operator.Component
}

func NewWarpComponent(cp operator.Component) operator.Component {
	return &ConcreateComponent{
		component: cp,
	}
}

func (c *ConcreateComponent) Calc() int {
	// 执行装饰器里面的逻辑
	c.component.Calc()
	// TODO：这里面书写我们所需要的核心逻辑
	return 0
}
