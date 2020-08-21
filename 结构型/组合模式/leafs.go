package main

import "fmt"

type Leaf struct {
	component
}

func NewLeaf() *Leaf { // 开辟一个叶子
	return &Leaf{}
}

func (c *Leaf) Print(pre string) {
	fmt.Println(pre, c.Name())
}

type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
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
