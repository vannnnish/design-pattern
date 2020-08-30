package main

import "fmt"

type Component interface {
	Calc() int
}

type ConcreateComponent struct {
}

func (*ConcreateComponent) Calc() int {
	return 0
}

func main() {
	fmt.Println(WarpAddComponent(&ConcreateComponent{}, 2).Calc())
}
