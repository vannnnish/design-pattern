package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

// 原型对象
type PrototypeManger struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManger {
	return &PrototypeManger{make(map[string]Cloneable)}
}

func (p *PrototypeManger) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManger) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

type Type1 struct {
	name string
}

// 深复制
func (t *Type1) Clone() Cloneable {
	tc := *t // 开辟内存 新建变量，存储指针指向的内容
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	return t
}

func main() {
	mgr := NewPrototypeManager()
	t1 := &Type1{
		"type1"}
	mgr.Set("t1", t1)
	t11 := mgr.Get("t1")
	t22 := t11.Clone()
	if t11 == t22 {
		fmt.Println("浅复制")
	} else {
		fmt.Println("深复制")
	}

	t2 := &Type2{
		"type2",
	}
	mgr.Set("t2", t2)
	t33 := mgr.Get("t2")
	t44 := t33.Clone()
	if t33 == t44 {
		fmt.Println("浅复制")
	} else {
		fmt.Println("深复制")
	}
}
