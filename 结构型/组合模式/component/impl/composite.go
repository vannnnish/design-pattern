package impl

import (
	"fmt"
	component2 "pattern/结构型/组合模式/component"
)

type Composite struct {
	component
	childs []component2.Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]component2.Component, 0),
	}
}

func (c *Composite) AddChild(child component2.Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child) // 加入孩子节点
}

// 打印显示每一个节点
func (c *Composite) Print(pre string) {
	fmt.Println(pre, c.name)
	pre += "  "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}
