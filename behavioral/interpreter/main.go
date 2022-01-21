package main

import "fmt"

func main(){
	p:=&Parser{
		exp:   nil,
		index: 0,
		prev:  nil,
	}
	p.Parse("1 + 2 - 2 + 10 - 2")
	fmt.Println(p.Result().Interpret())

}