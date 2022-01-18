package main

import (
	"fmt"
)

// 目的是实现功能的可拓展性。
type API interface {
	Say(name string) string
}

func NewAPI(str string) API {
	switch str {
	case "en":
		return &English{}
	case "cn":
		return &Chinese{}
	default:
		return nil
	}
}

type Chinese struct{}

func (*Chinese) Say(name string) string {
	return "你好:" + name
}

type English struct{}

func (*English) Say(name string) string {
	return "hello:" + name
}

func main() {
	api := NewAPI("en")
	fmt.Println(api.Say("李白"))
}
