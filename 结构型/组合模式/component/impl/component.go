package impl

import component2 "pattern/结构型/组合模式/component"

const (
	LeafNode = iota
	CompositeNode
)

type component struct {
	parent component2.Component
	name   string
}

func NewComponent(kind int, name string) component2.Component {
	var c component2.Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()

	}
	c.SetName(name)
	return c
}

func (c *component) Parent() component2.Component {
	return c.parent
}

func (c *component) SetParent(component component2.Component) {
	c.parent = component
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(component component2.Component) {
}

func (c *component) Print(s string) {
}
