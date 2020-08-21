package main

import "fmt"

func main() {
	var fac OperatorFactory
	fac = PlusOperatorFactory{}
	add := fac.Create()
	add.SetLeft(20)
	add.SetRight(10)
	fmt.Println(add.Result())

	fac=SubOperatorFactory{}
	sub:=fac.Create()
	sub.SetLeft(20)
	sub.SetRight(10)
	fmt.Println(sub.Result())

}
