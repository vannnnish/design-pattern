package main

import "pattern/结构型/组合模式/component/impl"

/*
	组合模式， 又可以叫做，部分-整体模式，  在代码逻辑里面的两个对象，	具有不一样的功能，
	但是又希望将他们一致处理。


	实现思路就是， 先定义两个对象相关操作的全部集合， 然后定义一个component对象，实现组合操作的
	的相关方法， 最后在分别实现，这两个对象各自需要实现的方法。

	比如，例子中， leafs自己要实现的方法就是print
		composite要实现的就是 AddChild 和 Print
*/
func main() {
	root := impl.NewComponent(impl.CompositeNode, "root")
	r1 := impl.NewComponent(impl.CompositeNode, "r1")
	r2 := impl.NewComponent(impl.CompositeNode, "r2")
	r3 := impl.NewComponent(impl.CompositeNode, "r3")

	l1 := impl.NewComponent(impl.LeafNode, "l1")
	l2 := impl.NewComponent(impl.LeafNode, "l2")
	l3 := impl.NewComponent(impl.LeafNode, "l3")
	root.AddChild(r1)
	root.AddChild(l1)
	root.AddChild(r3)
	root.AddChild(r2)

	r1.AddChild(l1)
	r3.AddChild(l2)
	r1.AddChild(l3)
	root.Print("xixi")

}
