package main

type Component interface {
	Parent() Component
	SetParent(component Component)
	Name() string
	SetName(string)
	AddChild(component Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

type component struct {
	parent Component
	name   string
}

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()

	}
	c.SetName(name)
	return c
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(component Component) {
	c.parent = component
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(component Component) {
}

func (c *component) Print(s string) {
}



func main(){
	root:=NewComponent(CompositeNode,"root")
	r1:=NewComponent(CompositeNode,"r1")
	r2:=NewComponent(CompositeNode,"r2")
	r3:=NewComponent(CompositeNode,"r3")

	l1:=NewComponent(LeafNode,"l1")
	l2:=NewComponent(LeafNode,"l2")
	l3:=NewComponent(LeafNode,"l3")
	root.AddChild(r1)
	root.AddChild(l1)
	root.AddChild(r3)
	root.AddChild(r2)

	r1.AddChild(l1)
	r3.AddChild(l2)
	r1.AddChild(l3)
	root.Print("xixi")

}