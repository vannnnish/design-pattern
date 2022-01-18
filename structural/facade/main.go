package main

import "fmt"

/*
	从实现角度来看: 外观模式更像是建造者模式的具体体现。
*/
func main() {
	api := NewTAPI()
	fmt.Println(api.Test())
}
