package impl

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
