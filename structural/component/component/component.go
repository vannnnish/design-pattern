package component

type Component interface {
	Parent() Component
	SetParent(component Component)
	Name() string
	SetName(string)
	AddChild(component Component)
	Print(string)
}

